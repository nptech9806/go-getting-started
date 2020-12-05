package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/timestreamquery"
	"log"
	"net/http"
	"os"
	 "time"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		//c.HTML(http.StatusOK, "health.tmpl.html", nil)
		c.HTML(http.StatusOK, "pingdom.tmpl.html", nil)
	})
	
	router.GET("/sleep2", func(c *gin.Context) {
		time.Sleep(2 * time.Second)
	})

	router.GET("/us-east-1", func(c *gin.Context) {
		sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
		}
	})

	router.GET("/pingdom", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pingdom.tmpl.html", nil)
	})
	
	router.Run(":" + port)
}
