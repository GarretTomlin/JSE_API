package main

import (
	"JSE_API/pkg/models"
	"github.com/gocolly/colly/queue"
	ds"JSE_API/pkg/configs"
)

var Storage *models.UrlStorage

func main(){
	  // Connect to the database
	  db, err := ds.ConnectToDB("sqlite3", "./urls.db")
	  if err != nil {
		 panic(err)
	}
	  // Create a new Storage instance
	  Storage := &models.UrlStorage{Db: db, Cache: make(map[string]bool)}

	  q, _ := queue.New(
		// Use 10 threads for scraping
		10,
		// Use the MyStorage instance as the storage backend for the URL queue
		Storage,
	  )



  // Start scraping
  q.Run(ds.C)

}
