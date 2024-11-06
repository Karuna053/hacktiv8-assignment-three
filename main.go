package main

import (
	"autoreload-data/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// call the function with goroutine
	go controllers.AutoReload()

	// start api server
	StartAPIServer()
}

func StartAPIServer() {
	// start api server wtih gin
	r := gin.Default()

	// define routes
	r.POST("/auto-reload", controllers.AutoReloadHandler)

	// run the server
	r.Run(":8080")
}
