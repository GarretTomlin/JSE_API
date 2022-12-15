package addLinkToDb

import(
	handler "JSE_API/pkg/configs/config_database"
	"log"
	"net/url"
	"regexp"

)

func Ingest(url *url.URL, StockInfoFilters []*regexp.Regexp) bool{
	u, err := url.Parse(url.String())
	if err != nil {
		// Handle the error by logging it or taking some other action
		log.Println(err)
		return false
	} else {
		// Flag to track whether the URL was skipped by a filter
		skipped := false

		// Check if the URL matches any of the filters
		for _, filter := range StockInfoFilters {
			if filter.MatchString(u.String()) {
				// If the URL matches a filter, skip it and set the skipped flag to true
				skipped = true
				continue
			}
		}

		// Check if the URL was not skipped by any of the filters
		if !skipped {
			// Visit the URL by calling the String method on the u variable
			handler.Storage.Put(u)
		return true
		}
	}
	return false
}
