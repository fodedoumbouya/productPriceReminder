package scrapers

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func Scrape() string {
	// Logic to scrape data from the web
	c := colly.NewCollector()
	return visitWeb(c)
}

func visitWeb(c *colly.Collector) string {
	fmt.Println("Visiting: https://en.wikipedia.org/wiki/Main_Page", *c)
	// make a channel
	channel := make(chan string)
	var mapping = []string{}
	fmt.Println(mapping, "Mapping")
	go func() {
		defer close(channel)
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting: ", r.URL)
		})

		c.OnError(func(_ *colly.Response, err error) {
			log.Println("Something went wrong: ", err)
			channel <- "error"

		})

		c.OnResponse(func(r *colly.Response) {
			body := bytes.NewReader(r.Body)

			// Load the HTML document
			doc, err := goquery.NewDocumentFromReader(body)
			if err != nil {
				log.Fatal(err)
			}

			// Find the review items
			doc.Find(".left-content article .post-title").Each(func(i int, s *goquery.Selection) {
				// For each item found, get the title
				title := s.Find("a").Text()
				fmt.Printf("Review %d: %s\n", i, title)
			})

			// channel <- doc.Text()

		})

		c.OnHTML("#root", func(e *colly.HTMLElement) {
			// printing all URLs associated with the a links in the page
			// fmt.Println(e.Attr("href"))
			// fmt.Println(e.DOM)

			// mapping = append(mapping, e.Attr("href"))

			body := bytes.NewReader(e.Response.Body)

			// Load the HTML document
			doc, err := goquery.NewDocumentFromReader(body)
			if err != nil {
				log.Fatal(err)
			}

			// Find the review items
			// doc.Find(".left-content article .post-title").Each(func(i int, s *goquery.Selection) {
			// 	// For each item found, get the title
			// 	title := s.Find("a").Text()
			// 	fmt.Printf("Review %d: %s\n", i, title)
			// })

			channel <- doc.Text()

			// body := string(e.Response.Body)
			// channel <- body

		})

		c.OnScraped(func(r *colly.Response) {
			fmt.Println(r.Request.URL, " scraped!")
			// body := string(r.Body)
			// fmt.Println(mapping)
			// channel <- strings.Join(mapping, ",\n\n\n")

			// body := bytes.NewReader(r.Body)

			// // Load the HTML document
			// doc, err := goquery.NewDocumentFromReader(body)
			// if err != nil {
			// 	log.Fatal(err)
			// }

			// // Find the review items
			// doc.Find(".left-content article .post-title").Each(func(i int, s *goquery.Selection) {
			// 	// For each item found, get the title
			// 	title := s.Find("a").Text()
			// 	fmt.Printf("Review %d: %s\n", i, title)
			// })

			channel <- "done"

		})
		c.Visit("https://fr.aliexpress.com/item/1005006457833026.html?pdp_ext_f=%7B%22ship_from%22:%22CN%22,%22sku_id%22:%2212000037268503195%22%7D&scm=1007.44674.357725.0&scm_id=1007.44674.357725.0&scm-url=1007.44674.357725.0&pvid=495773c5-0102-4ef9-abc7-eaccc82839b4&utparam=%257B%2522process_id%2522%253A%2522standard-portal-process-2%2522%252C%2522x_object_type%2522%253A%2522product%2522%252C%2522pvid%2522%253A%2522495773c5-0102-4ef9-abc7-eaccc82839b4%2522%252C%2522belongs%2522%253A%255B%257B%2522floor_id%2522%253A%252242500065%2522%252C%2522id%2522%253A%252233245288%2522%252C%2522type%2522%253A%2522dataset%2522%257D%252C%257B%2522id_list%2522%253A%255B%255D%252C%2522type%2522%253A%2522gbrain%2522%257D%255D%252C%2522pageSize%2522%253A%252220%2522%252C%2522language%2522%253A%2522fr%2522%252C%2522scm%2522%253A%25221007.44674.357725.0%2522%252C%2522countryId%2522%253A%2522FR%2522%252C%2522scene%2522%253A%2522choiceTopNWaterfall%2522%252C%2522tpp_buckets%2522%253A%252221669%25230%2523265320%252336_21669%25234190%252319158%252331_34674%25230%2523357725%25230%2522%252C%2522x_object_id%2522%253A%25221005006457833026%2522%257D&pdp_npi=4@dis!EUR!%E2%82%AC%2020,07!%E2%82%AC%200,99!!!151.25!7.47!@2103201917148367476367035eb816!12000037268503195!gdf!FR!!&aecmd=true")
		// ("https://www.amazon.com/dp/B0BNK5F2GN")
		time.Sleep(5 * time.Second)

		// "https://en.wikipedia.org/wiki/Main_Page")
	}()

	return <-channel
}
