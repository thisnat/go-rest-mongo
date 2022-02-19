package main

import (
	"github.com/thisnat/go-rest-mongo/configs"
	"github.com/thisnat/go-rest-mongo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())

	configs.ConnectDB()

	r.GET("/heartbeat", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	routes.NoteRoute(r)

	r.Run(":3001")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
