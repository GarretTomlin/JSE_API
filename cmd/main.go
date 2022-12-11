package main

import (
	"JSE_API/pkg/models"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	ds"JSE_API/pkg/configs"
)

func main(){
	  // Connect to the database
	  db, err := ds.ConnectToDB("sqlite3", "./urls.db")
	  if err != nil {
		 panic(err)
	}


	  // Create a new MyStorage instance
	  storage := &models.UrlStorage{Db: db, Cache: make(map[string]bool)}

	  // Create a new colly collector
	  c := colly.NewCollector(
		// Use multiple threads to improve scraping performance
		colly.Async(true),
	  )


	  q, _ := queue.New(
		// Use 10 threads for scraping
		10,
		// Use the MyStorage instance as the storage backend for the URL queue
		storage,
	  )



  // Start scraping
  q.Run(c)

}
