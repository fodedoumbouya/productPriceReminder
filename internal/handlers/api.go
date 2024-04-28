package handlers

import (
	"github.com/gin-gonic/gin"
)

/// where to put all the api handler only [function]
// so the the function that will be call should be in others file but still in the same package

func Handler(r *gin.Engine) {
	// Routes
	r.GET("/ping", pingTest)

}
