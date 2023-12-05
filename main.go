package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "go-url-shorter",
		})
	})

	err := r.Run(":9080")
	if err != nil {
		panic(fmt.Sprintf("failure to start web server, %s", err))
	}
}
