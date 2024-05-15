package handlers

import (
	"github.com/fodedoumbouya/productPriceReminder/internal/scrapers"
	"github.com/gin-gonic/gin"
)

func testFunc(g *gin.Context) {
	url := g.Query("url")
	if url == "" {
		g.JSON(400, gin.H{
			"response": scrapers.ScrapersResponse{
				Message: "Url is required",
				IsError: true,
			},
		})
		return
	}
	resp := scrapers.Scrape(url)
	var code int
	if resp.IsError {
		code = 400
	} else {
		code = 200
	}
	g.JSON(code, gin.H{
		"response": resp,
	})
}
