package configs

import (
	"JSE_API/pkg/models"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"time"
)


var C *colly.Collector
var Storage *models.UrlStorage

func Init() {
	C = colly.NewCollector(
		// Use multiple threads to improve scraping performance
		colly.Async(true),
	)
	C.OnHTML("a.core", func(e *colly.HTMLElement) {
	})
	// Set a custom user agent to avoid getting banned by the website
	C.UserAgent = "JSE_API/1.0"

	// Enable response caching to avoid making duplicate requests
	C.CacheDir = "./cache"

	// Connect to the database
	db, err := models.ConnectToDB("sqlite3", "./urls.db")
	if err != nil {
		panic(err)
	}
	// Create a new Storage instance
	Storage := &models.UrlStorage{Db: db, Cache: make(map[string]bool)}

	q, _ := queue.New(
		// Use 10 threads for scraping
		10,
		// Use the Storage instance as the storage backend for the URL queue
		Storage,
	)
// Create a ticker that ticks every 24 hours
ticker := time.NewTicker(time.Hour * 24)

for {
	// Start scraping
	q.Run(C)

	// Wait for the ticker to tick
	<-ticker.C
}

}
