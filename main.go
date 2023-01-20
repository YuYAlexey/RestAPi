package main

import (
	"net/http"

	"github.com/adYushinW/RestAPi/connect"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	connect.ConnectDB()

	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello"})
	})

	route.Run()
}
