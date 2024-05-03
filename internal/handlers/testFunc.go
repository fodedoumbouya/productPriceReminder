package handlers

import (
	"github.com/fodedoumbouya/productPriceReminder/internal/scrapers"
	"github.com/gin-gonic/gin"
)

func testFunc(c *gin.Context) {
	resp := scrapers.Scrape()

	c.JSON(200, gin.H{
		"message": resp,
	})
}
