package scrapers

import (
	"context"
	"strings"

	"github.com/chromedp/chromedp"
)

var (
	possibleCurrency []string = []string{"€", "$"}
)

func aliexpress(url string) ScrapersResponse {
	channel := make(chan ScrapersResponse)
	// Create a chrome context

	priceSelector := `pdp-info-right`
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	go scrapeTheWebsite(ctx, priceSelector, &channel, url)
	defer close(channel)
	return <-channel
}
func (r *rawDataHtml) parsePriceInfo() *ScrapersResponse {
	priceStr := r.PriceSelector
	priceStr =
		strings.ReplaceAll(priceStr, ",", ".")
	priceStr = strings.ReplaceAll(priceStr, "-", "")
	currencySymbol := ""
	for _, v := range possibleCurrency {
		if strings.Contains(priceStr, v) {
			currencySymbol = v
			break
		}
	}
	if currencySymbol == "" {
		return &ScrapersResponse{
			Message: "Currency symbol not found",
			IsError: true,
		}
	}

	// Split the string by the euro symbol and the dash
	parts := strings.Split(priceStr, currencySymbol)
	if len(parts) != 3 {
		return &ScrapersResponse{
			Message: "Invalid format",
			IsError: true,
		}
	}
	discount := strings.TrimSuffix(parts[2], "%")
	isDiscount := discount != "0" && discount != ""
	return &ScrapersResponse{
		Message: "Successfully parsed the price",
		IsError: false,
		DataScraper: DataScraper{
			DiscountPrice: parts[0],
			NormalPrice:   parts[1],
			Discount:      discount,
			Currency:      currencySymbol,
			IsDiscount:    isDiscount,
		},
	}
}
