package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world by gin",
		})
	})
	r.GET("/string", func(c *gin.Context) {
		c.String(200, "is a string")
	})
	r.GET("/stream", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		times := 10

		for {
			if times <= 0 {
				return
			}
			c.SSEvent("message", times)
			time.Sleep(1 * time.Second)
			times--
		}

	})
	r.Run(":8099") // listen and serve on

}
