package database

import (
	"database/sql"
	"errors"
	"log"
	"net/url"
	"time"
	_ "github.com/lib/pq"

)

type UrlStorage struct {
	Db    *sql.DB
	Url   []string
	Cache map[string]bool
}

func ConnectToDB(driver string, dataSource string) (*sql.DB, error) {
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// It satisfies the queue.Storage interface.
func (s *UrlStorage) Init() error {
	// Initialize the storage here.
	// This might involve creating tables in the database,
	// opening files, or any other necessary setup.
	return nil

}

func (s *UrlStorage) Put(url *url.URL) error {

	// Insert the URL into the database if it is not already visited
	_, err := s.Db.Exec(
		"INSERT INTO urls (url) SELECT ? WHERE NOT EXISTS (SELECT url FROM visited WHERE url = ?)",
		url.String(), url.String(),
	)
	return err
}



func (s *UrlStorage) Visited(url *url.URL) error {
	// Insert the URL into the visited table
	_, err := s.Db.Exec("INSERT INTO visited (url) VALUES (?)", url.String())
	// Use a ticker to push visited URLs back into the queue after 24 hours
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for range ticker.C {

			// Query the database for all visited URLs
			rows, err := s.Db.Query("SELECT url FROM visited")
			if err != nil {
				// Log the error and continue
				log.Println(err)
				continue
			}
			// Iterate over the rows and enqueue each URL
			for rows.Next() {
				var urlString string
				if err := rows.Scan(&urlString); err != nil {
					// Log the error and continue
					log.Println(err)
					continue
				}
				// Parse the URL and enqueue it
				u, err := url.Parse(urlString)
				if err != nil {
					// Log the error and continue
					log.Println(err)
					continue
				}
				if err := s.Put(u); err != nil {
					// Log the error and continue
					log.Println(err)
					continue
				}
			}
		}
	}()
	return err
}

func (s *UrlStorage) Get() (*url.URL, error) {
	var urlString string
	// Check if the URL has already been retrieved
	if visited, ok := s.Cache[urlString]; ok && visited {
		// If it has been retrieved, return an error
		return nil, errors.New("URL has already been retrieved")
	}
	// Retrieve the next URL from the database
	err := s.Db.QueryRow(
		"SELECT url FROM urls WHERE url NOT IN (SELECT url FROM visited) LIMIT 1",
	).Scan(&urlString)
	if err != nil {
		return nil, err
	}
	// Parse the URL and return it
	url, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	// Mark the URL as visited in the cache
	s.Cache[urlString] = true
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for range ticker.C {
			s.Cache = make(map[string]bool)
		}
	}()
	// Mark the URL as visited in the database
	err = s.Visited(url)
	if err != nil {
		return nil, err
	}

	return url, nil

}

func (s *UrlStorage) AddRequest(urlArr []byte) error {
	// Parse the URL and add it to the storage.
	u, err := url.Parse(string(urlArr))
	if err != nil {
		return err
	}
	return s.Put(u)
}

// This method gets the next request from the storage.
// It satisfies the queue.Storage interface.
func (s *UrlStorage) GetRequest() ([]byte, error) {
	// Get the next URL from the storage.
	urlArr, err := s.Get()
	if err != nil {
		return nil, err
	}
	// Return the URL as a byte slice.
	return []byte(urlArr.String()), nil
}

// This method gets the size of the queue.
// It satisfies the queue.Storage interface.
func (s *UrlStorage) QueueSize() (int, error) {
	// Get the number of URLs in the queue.
	var count int
	err := s.Db.QueryRow(
		"SELECT COUNT(*) FROM urls WHERE url NOT IN (SELECT url FROM visited)",
	).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

