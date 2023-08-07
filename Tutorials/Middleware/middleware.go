package main

import (
	"net/http"
	"github.com/gin-goinc/gin"
)

func main() {
	router := gin.Default()

	router.Get("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message":"hello!",
		})
	})

	router.Run()
}
