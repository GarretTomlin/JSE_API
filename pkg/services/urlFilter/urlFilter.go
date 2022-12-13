package urlFilter

import (
	c "JSE_API/pkg/configs"
	"log"
	"regexp"
)

func VisitIfNotFiltered(urlFilters []*regexp.Regexp) bool {
	// Get the URL from the database
	u, err := c.Storage.Get()
	if err != nil {
		// Handle the error by logging it or taking some other action
		log.Println(err)
		return false
	} else {
		// Flag to track whether the URL was skipped by a filter
		skipped := false

		// Check if the URL matches any of the filters
		for _, filter := range urlFilters {
			if filter.MatchString(u.String()) {
				// If the URL matches a filter, skip it and set the skipped flag to true
				skipped = true
				continue
			}
		}

		// Check if the URL was not skipped by any of the filters
		if !skipped {
			// Visit the URL by calling the String method on the u variable
			c.C.Visit(u.String())
			return true
		}
	}
	return false
}
