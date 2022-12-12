package controllers

import (
	store "JSE_API/cmd"
	c "JSE_API/pkg/configs"
	"github.com/gocolly/colly"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"regexp"
)

type TradeSummary struct {
}

func (s *TradeSummary) GetStockAdvancing(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	type data struct {
		SYMBOL        string `json:"SYMBOL"`
		SECURITY      string `json:"SECURITY"`
		VOLUME        string `json:"VOLUME"`
		CLOSING_PRICE string `json:"CLOSING_PRICE"`
		PRICE_CHANGE  string `json:"PRICE_CHANGE"`
		CHANGE        string `json:"CHANGE"`
	}
	summary := make([]data, 0)

	// OnHTML registers a callback function which is called every time a matching HTML element is found in the page
	c.C.OnHTML("tr[class=tw-bg-white tw-text-gray-800]", func(e *colly.HTMLElement) {
		// Extract the data you want to scrape
		symbol := e.ChildText("td a")
		security := e.ChildText("td a")
		volume := e.ChildText("td[class=tw-px-3 tw-py-4 tw-text-sm tw-text-right tw-whitespace-nowrap]")
		closingPrice := e.ChildText("td[class=tw-px-3 tw-py-4 tw-text-sm tw-text-right tw-whitespace-nowrap]")
		priceChange := e.ChildText("td[class=tw-px-3 tw-py-4 tw-text-sm tw-text-right tw-whitespace-nowrap]")
		changePercent := e.ChildText("td[class=tw-px-3 tw-py-4 tw-text-sm tw-text-right tw-whitespace-nowrap]")

		// Create a new data struct and append it to the summary slice
		allData := data{
			SYMBOL:        symbol,
			SECURITY:      security,
			VOLUME:        volume,
			CLOSING_PRICE: closingPrice,
			PRICE_CHANGE:  priceChange,
			CHANGE:        changePercent,
		}

		summary = append(summary, allData)
	})

	// Create a list of URL filters as regular expressions
	var urlFilters = []*regexp.Regexp{
		regexp.MustCompile(`example\.com`),
		regexp.MustCompile(`example\.org`),
	}

	// Get the URL from the store
	u, err := store.Storage.Get()
	if err != nil {
		// Handle the error by logging it or taking some other action
		log.Println(err)
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
		}
	}

	// Send the JSON response
	SendJson(w, http.StatusOK, summary)
}
