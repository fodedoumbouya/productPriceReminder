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
