package main

import "github.com/gin-gonic/gin"

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
	r.Run(":8099") // listen and serve on

}
