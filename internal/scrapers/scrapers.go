package scrapers

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func Scrape() {
	// Logic to scrape data from the web
	c := colly.NewCollector()
	visitWeb(c)
}

func visitWeb(c *colly.Collector) {
	fmt.Println("Visiting: https://en.wikipedia.org/wiki/Main_Page", *c)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		// printing all URLs associated with the a links in the page
		fmt.Println(e.Attr("href"))
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})
	c.Visit("https://en.wikipedia.org/wiki/Main_Page")

}
