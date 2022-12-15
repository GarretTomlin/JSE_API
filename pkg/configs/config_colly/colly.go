package config_colly

import (
	handler "JSE_API/pkg/configs/config_database"
	proxies "JSE_API/pkg/services/proxySwitcher"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"time"
)

var C *colly.Collector

func Init() {
	C = colly.NewCollector(
		// Use multiple threads to improve scraping performance
		colly.Async(true),
		colly.AllowedDomains("jse.com", "jse-api.com"),

	)
	C.OnHTML("a.core", func(e *colly.HTMLElement) {
	})
	// Set a custom user agent to avoid getting banned by the website
	C.UserAgent = "JSE_API/1.0"

	// Enable response caching to avoid making duplicate requests
	C.CacheDir = "./cache"
	// Create a new Storage instance

	// Set a limit on the number of threads that can be used to make requests concurrently
	C.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	// Rotate proxies
	C.SetProxyFunc(proxies.RandomProxySwitcher)

	q, _ := queue.New(
		// Use 10 threads for scraping
		10,
		// Use the Storage instance as the storage backend for the URL queue
		handler.Storage,
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
