package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	r.Run(":8080")
}