package main

import (
	"derco-backend/api"
	"derco-backend/config"

	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
	if os.Getenv("API_URL") == "" || os.Getenv("API_SECRET") == "" {
		log.Println("All env vars must be set!!!")
		return nil
	}

	config.Conf.JsonPlaceHolder.ApiKey = os.Getenv("API_SECRET")
	config.Conf.JsonPlaceHolder.Url = os.Getenv("API_URL")

	r := gin.Default()

	corsEnv := os.Getenv("CORS")
	if corsEnv != "" {
		config := cors.DefaultConfig()
		config.AllowOrigins = []string{corsEnv}
		r.Use(cors.New(config))
	}

	r.GET("/", api.GetPosts)

	return r
}
