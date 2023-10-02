package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Hello test",
		})
	})

	router.Run()
}
