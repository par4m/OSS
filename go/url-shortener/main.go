package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/par4m/url-shortener/api/routes"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()
	setupRouters(router)

	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8080"
	}

	log.Fatal(router.Run("0.0.0.0:" + port))
}

func setupRouters(router *gin.Engine) {
	router.POST("/api/v1", routes.ShortenURL)
	router.GET("/api/v1/:shortID", routes.GetByShortID)
	router.PUT("/api/v1/:shortID", routes.EditURl)
	router.DELETE("/api/v1/:shortID", routes.DeleteUrl)
	router.POST("/api/v1/addTag", routes.AddTag)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

}
