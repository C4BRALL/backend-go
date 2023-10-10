package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/healthily", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	s := &http.Server{
		Addr:    ":8958",
		Handler: router,
	}
	s.ListenAndServe()
}
