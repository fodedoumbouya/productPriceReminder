package scrapers

import (
	"context"
	"fmt"
	"strings"

	"github.com/chromedp/chromedp"
)

type ScrapersResponse struct {
	Message     string `json:"message"`
	IsError     bool   `json:"isError"`
	DataScraper DataScraper
}

type DataScraper struct {
	DiscountPrice string `json:"discountPrice"`
	NormalPrice   string `json:"normalPrice"`
	Discount      string `json:"discount"`
	Currency      string `json:"currency"`
	IsDiscount    bool   `json:"isDiscount"`
}

func Scrape(url string) ScrapersResponse {
	// Logic to scrape data from the web
	// c := colly.NewCollector()

	return visitWeb(url)
}

func visitWeb(url string) ScrapersResponse {
	// create copy to be used inside the goroutine
	// cCp := c.Clone()
	resp := aliexpress(url)
	return resp
}

func scrapeTheWebsite(ctx context.Context, priceSelector string, channel *chan ScrapersResponse, url string) {
	var priceField string
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Text(priceSelector, &priceField),
	})
	if err != nil {
		fmt.Println(err)
		*channel <- ScrapersResponse{
			Message: "Error while scraping the website",
			IsError: true,
		}
	}
	priceField = strings.Split(priceField, "\n")[0]
	rawDataHtml := rawDataHtml{
		PriceSelector: priceField,
	}
	*channel <- *rawDataHtml.parsePriceInfo()
}

type parsePrice interface {
	parsePriceInfo() *ScrapersResponse
}

type rawDataHtml struct {
	PriceSelector string
}
