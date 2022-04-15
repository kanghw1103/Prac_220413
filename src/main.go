package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong again",
			"name":    "MK",
		})
	})

	err := r.Run()
	if err != nil {
		return
	}
}
