package scrapers

import (
	"bytes"
	"fmt"
	"log"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type ScrapersResponse struct {
	Message     string `json:"message"`
	IsError     bool   `json:"isError"`
	DataScraper DataScraper
}

type DataScraper struct {
	Price           string `json:"price"`
	Currency        string `json:"currency"`
	FormattedAmount string `json:"formattedAmount"`
	IsDiscount      bool   `json:"isDiscount"`
}

func Scrape(url string) ScrapersResponse {
	// Logic to scrape data from the web
	c := colly.NewCollector()

	return visitWeb(c, url)
}

func visitWeb(c *colly.Collector, url string) ScrapersResponse {
	// create copy to be used inside the goroutine
	cCp := c.Clone()
	resp := aliexpress(cCp, url)
	return resp
}

func aliexpress(c *colly.Collector, url string) ScrapersResponse {
	channel := make(chan ScrapersResponse)
	// discount skuActivityAmount
	// no discount skuAmount
	fieldTocheckForPrice := map[string]bool{
		"skuActivityAmount": true,
		"skuAmount":         false,
	}

	scrapeTheWebsite(c, fieldTocheckForPrice, channel, url)

	return <-channel
}

func scrapeTheWebsite(c *colly.Collector, fieldTocheckForPrice map[string]bool, channel chan ScrapersResponse, url string) {
	go func() {
		defer close(channel)
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting: ", r.URL)
		})
		c.OnError(func(_ *colly.Response, err error) {
			log.Println("Something went wrong: ", err)
			channel <- ScrapersResponse{IsError: true,
				Message: fmt.Sprint("Something went wrong error:", err),
			}

		})

		c.OnScraped(func(r *colly.Response) {
			fmt.Println(r.Request.URL, "\nscraped!")
			body := bytes.NewReader(r.Body)
			// Load the HTML document
			doc, err := goquery.NewDocumentFromReader(body)
			if err != nil {
				channel <- ScrapersResponse{
					Message: "Error loading HTML document",
					IsError: true}
				return
			}
			data := doc.Find("script").Text()
			var matches []string
			isDiscount := false
			for v, isProductDiscount := range fieldTocheckForPrice {
				re := regexp.MustCompile(fmt.Sprintf(`"%s"\s*:\s*{\s*"currency"\s*:\s*"([^"]+)"\s*,\s*"formatedAmount"\s*:\s*"([^"]+)"\s*,\s*"value"\s*:\s*([0-9.]+)`, v))
				matches = re.FindStringSubmatch(data)
				// fmt.Println(matches)
				if len(matches) > 3 {
					isDiscount = isProductDiscount
					break
				}
			}
			var response ScrapersResponse

			if len(matches) > 3 {
				fmt.Printf("Currency: %s\n", matches[1])
				fmt.Printf("Formatted Amount: %s\n", matches[2])
				fmt.Printf("Value: %s\n", matches[3])
				response = ScrapersResponse{
					DataScraper: DataScraper{
						Price:           matches[3],
						Currency:        matches[1],
						FormattedAmount: matches[2],
						IsDiscount:      isDiscount,
					},
					Message: "success",
				}
			} else {
				re2 := regexp.MustCompile(`"PRICE":{"skuSecondPriceInfoMap":{},"productId":1005006246199032,"discountExt":"[^"]+","targetSkuPriceInfo":{"sellerByLot":false,"salePriceLocal":"([^"]+)","salePriceString":"([^"]+)","priceFontColor":"#000000"},"selectedSkuId":12000036519122316,"skuPriceInfoMap":{"12000036519122317":{"sellerByLot":false,"salePriceLocal":"([^"]+)","salePriceString":"([^"]+)","priceFontColor":"#000000"},"12000036519122316":{"$ref":"$.PRICE.targetSkuPriceInfo"},"12000036519122313":{"sellerByLot":false,"salePriceLocal":"([^"]+)","salePriceString":"([^"]+)","priceFontColor":"#000000"},"12000036519122315":{"sellerByLot":false,"salePriceLocal":"([^"]+)","salePriceString":"([^"]+)","priceFontColor":"#000000"}},"isLot":false,"skuIdStrPriceInfoMap":{"12000036519122317":{"sellerByLot":false,"salePriceLocal":"([^"]+)","salePriceString":"([^"]+)","priceFontColor":"#000000"},"12000036519122316":{"sellerByLot":false,"salePriceLocal":"([^"]+)","salePriceString":"([^"]+)","priceFontColor":"#000000"},"12000036519122313":{"sellerByLot":false,"salePriceLocal":"([^"]+)","salePriceString":"([^"]+)","priceFontColor":"#000000"},"12000036519122315":{"sellerByLot":false,"salePriceLocal":"([^"]+)","salePriceString":"([^"]+)","priceFontColor":"#000000"}},"region":"FR","priceLocalConfig":"[^"]+"}`)
				matches2 := re2.FindStringSubmatch(data)
				response = ScrapersResponse{
					Message: doc.Text(),
					// "No match found",
					IsError: true}
				fmt.Println("response", matches2)
				// fmt.Println("doc", doc.Text())

			}

			channel <- response

		})
		c.Visit(url)
		//.pdp-info .pdp-info-right .pdp-comp-price-current
		// "https://fr.aliexpress.com/item/1005006330040361.html?spm=a2g0o.detail.pcDetailTopMoreOtherSeller.3.39c3hrgJhrgJm0&gps-id=pcDetailTopMoreOtherSeller&scm=1007.40050.354490.0&scm_id=1007.40050.354490.0&scm-url=1007.40050.354490.0&pvid=723e82eb-276b-40e9-b033-d7470a135c2a&_t=gps-id:pcDetailTopMoreOtherSeller,scm-url:1007.40050.354490.0,pvid:723e82eb-276b-40e9-b033-d7470a135c2a,tpp_buckets:668%232846%238115%232000&pdp_npi=4%40dis%21EUR%213.42%210.99%21%21%213.62%211.05%21%40210321dc17157644252677164ec0cb%2112000036783780582%21rec%21FR%21%21AB&utparam-url=scene%3ApcDetailTopMoreOtherSeller%7Cquery_from%3A")
		// "https://fr.aliexpress.com/item/1005005151305507.html?pdp_ext_f=%7B%22ship_from%22:%22CN%22,%22sku_id%22:%2212000031876709520%22%7D&scm=1007.44674.357725.0&scm_id=1007.44674.357725.0&scm-url=1007.44674.357725.0&pvid=d7a8734e-5967-4a36-bfad-057d8b8167dc&utparam=%257B%2522process_id%2522%253A%2522standard-portal-process-2%2522%252C%2522x_object_type%2522%253A%2522product%2522%252C%2522pvid%2522%253A%2522d7a8734e-5967-4a36-bfad-057d8b8167dc%2522%252C%2522belongs%2522%253A%255B%257B%2522floor_id%2522%253A%252242500065%2522%252C%2522id%2522%253A%252233245288%2522%252C%2522type%2522%253A%2522dataset%2522%257D%252C%257B%2522id_list%2522%253A%255B%255D%252C%2522type%2522%253A%2522gbrain%2522%257D%255D%252C%2522pageSize%2522%253A%252220%2522%252C%2522language%2522%253A%2522fr%2522%252C%2522scm%2522%253A%25221007.44674.357725.0%2522%252C%2522countryId%2522%253A%2522FR%2522%252C%2522scene%2522%253A%2522choiceTopNWaterfall%2522%252C%2522tpp_buckets%2522%253A%252221669%25230%2523265320%25239_21669%25234190%252319166%2523895_34674%25230%2523357725%25230%2522%252C%2522x_object_id%2522%253A%25221005005151305507%2522%257D&pdp_npi=4@dis!EUR!%E2%82%AC%204,23!%E2%82%AC%204,23!!!31.90!31.90!@2103251317149331851925342efdba!12000031876709520!gdf!FR!!&aecmd=true")
		// "https://fr.aliexpress.com/item/1005006457833026.html?pdp_ext_f=%7B%22ship_from%22:%22CN%22,%22sku_id%22:%2212000037268503195%22%7D&scm=1007.44674.357725.0&scm_id=1007.44674.357725.0&scm-url=1007.44674.357725.0&pvid=495773c5-0102-4ef9-abc7-eaccc82839b4&utparam=%257B%2522process_id%2522%253A%2522standard-portal-process-2%2522%252C%2522x_object_type%2522%253A%2522product%2522%252C%2522pvid%2522%253A%2522495773c5-0102-4ef9-abc7-eaccc82839b4%2522%252C%2522belongs%2522%253A%255B%257B%2522floor_id%2522%253A%252242500065%2522%252C%2522id%2522%253A%252233245288%2522%252C%2522type%2522%253A%2522dataset%2522%257D%252C%257B%2522id_list%2522%253A%255B%255D%252C%2522type%2522%253A%2522gbrain%2522%257D%255D%252C%2522pageSize%2522%253A%252220%2522%252C%2522language%2522%253A%2522fr%2522%252C%2522scm%2522%253A%25221007.44674.357725.0%2522%252C%2522countryId%2522%253A%2522FR%2522%252C%2522scene%2522%253A%2522choiceTopNWaterfall%2522%252C%2522tpp_buckets%2522%253A%252221669%25230%2523265320%252336_21669%25234190%252319158%252331_34674%25230%2523357725%25230%2522%252C%2522x_object_id%2522%253A%25221005006457833026%2522%257D&pdp_npi=4@dis!EUR!%E2%82%AC%2020,07!%E2%82%AC%200,99!!!151.25!7.47!@2103201917148367476367035eb816!12000037268503195!gdf!FR!!&aecmd=true")
		// ("https://www.amazon.com/dp/B0BNK5F2GN")
		// time.Sleep(5 * time.Second)

	}()

}

// / data structure for the response of the scrapers from aliexpress
type SkuPriceInfo struct {
	SellerByLot     bool   `json:"sellerByLot"`
	SalePriceLocal  string `json:"salePriceLocal"`
	SalePriceString string `json:"salePriceString"`
	PriceFontColor  string `json:"priceFontColor"`
}

type Price struct {
	SkuSecondPriceInfoMap map[string]interface{}  `json:"skuSecondPriceInfoMap"`
	ProductId             string                  `json:"productId"`
	DiscountExt           string                  `json:"discountExt"`
	TargetSkuPriceInfo    SkuPriceInfo            `json:"targetSkuPriceInfo"`
	SelectedSkuId         string                  `json:"selectedSkuId"`
	SkuPriceInfoMap       map[string]SkuPriceInfo `json:"skuPriceInfoMap"`
	IsLot                 bool                    `json:"isLot"`
	SkuIdStrPriceInfoMap  map[string]SkuPriceInfo `json:"skuIdStrPriceInfoMap"`
	Region                string                  `json:"region"`
	PriceLocalConfig      string                  `json:"priceLocalConfig"`
}
