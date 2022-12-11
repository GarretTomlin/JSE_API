package controllers

import (
	c "JSE_API/pkg/configs"
	"github.com/gocolly/colly"
	"github.com/julienschmidt/httprouter"
	"net/http"
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

	// Send the JSON response
	SendJson(w, http.StatusOK, summary)
}
