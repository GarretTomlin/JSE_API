package controllers

import (
	c "JSE_API/pkg/configs/config_colly"
	d "JSE_API/pkg/services/addLinkToDb"
	filter "JSE_API/pkg/services/urlFilter"
	"net/http"
	"net/url"
	"regexp"
	"log"
	"github.com/gocolly/colly"
	"github.com/julienschmidt/httprouter"
)

type TradeSummary struct {
}

// Create a list of URL filters as regular expressions
var urlFilters = []*regexp.Regexp{
	regexp.MustCompile(`https://www.jamstockex.com/trading/trade-summary/`),
}
func (s *TradeSummary) GetStockAdvancing(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	type data struct {
		SYMBOL            string `json:"SYMBOL"`
		SECURITY_NAME     string `json:"SECURITY"`
		VOLUME            string `json:"VOLUME"`
		CLOSING_PRICE     string `json:"CLOSING_PRICE"`
		PRICE_CHANGE      string `json:"PRICE_CHANGE"`
		CHANGE_PERCENTAGE string `json:"CHANGE"`
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
			SYMBOL:            symbol,
			SECURITY_NAME:     security,
			VOLUME:            volume,
			CLOSING_PRICE:     closingPrice,
			PRICE_CHANGE:      priceChange,
			CHANGE_PERCENTAGE: changePercent,
		}

		summary = append(summary, allData)
	})

	filter.VisitIfNotFiltered(urlFilters)
	// Send the JSON response
	SendJson(w, http.StatusOK, summary)
}
func (s *TradeSummary) GetStockDeclining(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	type data struct {
		SYMBOL        string `json:"SYMBOL"`
		SECURITY      string `json:"SECURITY"`
		VOLUME        string `json:"VOLUME"`
		CLOSING_PRICE string `json:"CLOSING_PRICE"`
		PRICE_CHANGE  string `json:"PRICE_CHANGE"`
		CHANGE        string `json:"CHANGE"`
	}
	summary := make([]data, 0)

	// Send the JSON response
	SendJson(w, http.StatusOK, summary)

}

func (s *TradeSummary) GetStockSummary() {

	var pattern = []*regexp.Regexp{
		regexp.MustCompile(`https://www\.jamstockex\.com/trading/instruments/\?instrument`),
		}

	// Register the callback function to be called for each matching HTML element
	c.C.OnHTML("a[href]", func(e *colly.HTMLElement){
		// Loop through the slice of regular expression objects
	for _, regex := range pattern {
		// Check if the `href` attribute matches the regex pattern
		if regex.MatchString(e.Attr("href")) {
			// Create a new slice containing the single regular expression object
			slice := []*regexp.Regexp{regex}
			// Parse the `e.Text` string value into a *url.URL value
			u, err := url.Parse(e.Text)
			if err != nil {
				// Handle the error by logging it or taking some other action
				log.Println(err)
				break
			}
			// Pass the slice and the *url.URL value as arguments to the AddLinkToDb function
			d.Ingest(u, slice)
			break
		}
	}
	})
	filter.VisitIfNotFiltered(urlFilters)
}
