package main

import (
	"net/http"
	"os"

	controllers "main/Controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	gogpt "github.com/sashabaranov/go-openai"
)

func CORSMiddleware() gin.HandlerFunc {
	godotenv.Load()
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		//c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, gin.H{
				"message": "bruh",
			})
			return
		}

		c.Next()
	}
}
func main() {
	godotenv.Load()
	client := gogpt.NewClient(os.Getenv("OPENAPI_KEY"))

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/summarize/", func(c *gin.Context) {
		controllers.GetSummary(c, client)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
