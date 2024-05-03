package scrapers

import (
	"fmt"
	"log"
	"strings"

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
	mapping := []string{}
	go func() {
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting: ", r.URL)
		})

		c.OnError(func(_ *colly.Response, err error) {
			log.Println("Something went wrong: ", err)
			channel <- "error"

		})

		c.OnResponse(func(r *colly.Response) {
			fmt.Println("Page visited: ", r.Request.URL)
			// channel <- r.Request.URL.String()

		})

		c.OnHTML("a", func(e *colly.HTMLElement) {
			// printing all URLs associated with the a links in the page
			// fmt.Println(e.Attr("href"))
			// fmt.Println(e.DOM)

			mapping = append(mapping, e.Attr("href"))

		})

		c.OnScraped(func(r *colly.Response) {
			fmt.Println(r.Request.URL, " scraped!")
			// fmt.Println(mapping)
			channel <- strings.Join(mapping, ",\n\n\n")

		})
		c.Visit("https://www.amazon.com/dp/B0BNK5F2GN")
		// "https://en.wikipedia.org/wiki/Main_Page")
	}()

	return <-channel
}
