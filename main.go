package main

import (
	"net/http"
	"test/api"
	"test/connect/db"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	d := &db.Database{}
	d.ConnectDB()

	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello"})
	})

	api := &api.Todoinfo{}
	route.GET("/all", api.GetAll)

	route.Run()
}
