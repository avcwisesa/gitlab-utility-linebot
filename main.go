package main

import (
	"log"
	"os"

	"github.com/avcwisesa/gitlab-reference-linebot/client"
	"github.com/avcwisesa/gitlab-reference-linebot/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	channelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	channelAccessToken := os.Getenv("LINE_CHANNEL_TOKEN")

	privateToken := os.Getenv("GITLAB_PRIVATE_TOKEN")
	projectID := os.Getenv("GITLAB_PROJECT_ID")

	port := os.Getenv("WEB_PORT")

	log.Println(channelSecret)
	log.Println(channelAccessToken)
	log.Println(privateToken)
	log.Println(projectID)

	bot, err := linebot.New(channelSecret, channelAccessToken)
	if err != nil {
		log.Println(err)
		log.Panic("Error initiating Linebot client")
	}

	client := client.New(projectID, privateToken)
	handler := handler.New(client, bot)

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", handler.Ping)

	r.Run(port)
}
