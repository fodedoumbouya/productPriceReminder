package handlers

import (
	"github.com/fodedoumbouya/productPriceReminder/internal/scrapers"
	"github.com/gin-gonic/gin"
)

func testFunc(c *gin.Context) {
	scrapers.Scrape()

	c.JSON(200, gin.H{
		"message": "test",
	})
}
