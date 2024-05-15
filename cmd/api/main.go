package main

import (
	"fmt"
	"os"

	"github.com/fodedoumbouya/productPriceReminder/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Load the .env file in the current directory
	godotenv.Load("../../.env")
	fmt.Print(`
	_______  ______     _______  _______  _______  ______            _______ _________  __________________
	(  ____ \(  __  \   (  ____ )(  ____ )(  ___  )(  __  \ |\     /|(  ____ \\__   __/  \__   __/\__   __/
	| (    \/| (  \  )  | (    )|| (    )|| (   ) || (  \  )| )   ( || (    \/   ) (        ) (      ) (   
	| (__    | |   ) |  | (____)|| (____)|| |   | || |   ) || |   | || |         | |        | |      | |   
	|  __)   | |   | |  |  _____)|     __)| |   | || |   | || |   | || |         | |        | |      | |   
	| (      | |   ) |  | (      | (\ (   | |   | || |   ) || |   | || |         | |        | |      | |   
	| )      | (__/  )  | )      | ) \ \__| (___) || (__/  )| (___) || (____/\   | |     ___) (___   | |   
	|/       (______/   |/       |/   \__/(_______)(______/ (_______)(_______/   )_(     \_______/   )_(   
													      
	
`)
}

func main() {
	// Creates a router without any middleware by default
	r := gin.New()
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	handlers.Handler(r)
	/// generate a cool print with "FD" in "-"
	port := os.Getenv("PORT")

	fmt.Println("Starting server on :", port)
	host := "0.0.0.0:" + port

	r.Run(
		host,
	) // listen and serve on

}

// func parsePriceInfoToMap(priceStr string) (map[string]string, error) {
// 	// Replace comma with dot for correct float parsing
// 	priceStr = strings.ReplaceAll(priceStr, ",", ".")

// 	// Split the string by the euro symbol and the dash
// 	parts := strings.Split(priceStr, "€")
// 	if len(parts) != 3 {
// 		return nil, fmt.Errorf("invalid format")
// 	}

// 	// Create the map
// 	priceInfo := map[string]string{
// 		"PriceDiscount": parts[0],
// 		"PriceNormal":   parts[1],
// 		"Discount":      strings.TrimSuffix(parts[2], "%"),
// 		"Currency":      "euro",
// 	}

// 	return priceInfo, nil
// }

// func main() {
// 	// Define the URL to scrape
// 	url := "https://fr.aliexpress.com/item/1005006127476840.html?gps-id=pcJustForYou&scm=1007.13562.333647.0&scm_id=1007.13562.333647.0&scm-url=1007.13562.333647.0&pvid=99a10a4d-1e76-4cf3-80eb-2fbc54f422f3&_t=gps-id:pcJustForYou,scm-url:1007.13562.333647.0,pvid:99a10a4d-1e76-4cf3-80eb-2fbc54f422f3,tpp_buckets:668%232846%238113%231998&pdp_npi=4@dis!EUR!44.65!7.79!!!341.59!59.58!@2101d00017157770823567919e9a24!12000035879136973!rec!FR!!AB&utparam-url=scene:pcJustForYou%7Cquery_from:"
// 	// "https://fr.aliexpress.com/item/1005005934689079.html?gps-id=pcJustForYou&scm=1007.13562.333647.0&scm_id=1007.13562.333647.0&scm-url=1007.13562.333647.0&pvid=d9f7e083-5d70-4bd2-916a-7047771d0ba0&_t=gps-id:pcJustForYou,scm-url:1007.13562.333647.0,pvid:d9f7e083-5d70-4bd2-916a-7047771d0ba0,tpp_buckets:668%232846%238113%231998&pdp_npi=4@dis!EUR!3.01!0.99!!!23.04!7.62!@2101eff117157661408466617e74ad!12000034942175840!rec!FR!!AB&utparam-url=scene:pcJustForYou%7Cquery_from:"
// 	// "https://fr.aliexpress.com/item/1005006246199032.html?pdp_ext_f=%7B%22sku_id%22:%2212000036519122316%22%7D&utparam-url=scene:search%7Cquery_from:category_navigate"

// 	// fmt.Println(parsePriceInfoToMap("7.79€44.65€-82%"))

// 	// Create a channel to receive the title
// 	titleCh := make(chan string)
// 	defer close(titleCh) // Close the channel after use

// 	// Create a chrome context
// 	ctx, cancel := chromedp.NewContext(context.Background())
// 	defer cancel()

// 	// Define the title element selector
// 	titleSelector := `pdp-info-right`
// 	// priceSelector := `.price--currentPriceText--V8_y_b5.pdp-comp-price-current.product-price-value`

// 	// Goroutine to scrape the title and send it to the channel
// 	go func() {
// 		var title string
// 		err := chromedp.Run(ctx, chromedp.Tasks{
// 			chromedp.Navigate(url),
// 			chromedp.Text(titleSelector, &title),
// 		})
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		titleCh <- title // Send the title to the channel
// 	}()

// 	// Wait for the title to be received on the channel
// 	title := <-titleCh
// 	title = strings.Split(title, "\n")[0]
// 	test, e := parsePriceInfo(title, []string{"€", "$"})
// 	if e != nil {
// 		fmt.Println(e)
// 	}
// 	fmt.Println(test)

// 	// Print the extracted title
// 	fmt.Println("Wikipedia Title:", title)
// }
// func parsePriceInfo(priceStr string, possibleCurrency []string) (*PriceInfo, error) {
// 	// var (
// 	// 	currency, discount, priceDiscount, priceNormal string
// 	// )

// 	// Replace comma with dot for correct float parsing
// 	priceStr =
// 		strings.ReplaceAll(priceStr, ",", ".")
// 	priceStr = strings.ReplaceAll(priceStr, "-", "")
// 	currencySymbol := ""
// 	for _, v := range possibleCurrency {
// 		if strings.Contains(priceStr, v) {
// 			currencySymbol = v
// 			break
// 		}
// 	}
// 	if currencySymbol == "" {
// 		return nil, fmt.Errorf("currency symbol not found")
// 	}

// 	// Split the string by the euro symbol and the dash
// 	parts := strings.Split(priceStr, currencySymbol)
// 	if len(parts) != 3 {
// 		return nil, fmt.Errorf("invalid format")
// 	}

// 	return &PriceInfo{
// 		PriceDiscount: parts[0],
// 		PriceNormal:   parts[1],
// 		Discount:      strings.TrimSuffix(parts[2], "%"),
// 		Currency:      currencySymbol,
// 	}, nil
// }

// type PriceInfo struct {
// 	PriceDiscount string
// 	PriceNormal   string
// 	Discount      string
// 	Currency      string
// }
