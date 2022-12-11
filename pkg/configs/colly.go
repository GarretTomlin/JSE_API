package configs

import (
	"github.com/gocolly/colly"

)


var C *colly.Collector

func init() {
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

}
