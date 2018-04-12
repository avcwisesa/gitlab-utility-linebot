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

	privateToken := os.Getenv("GITLAB_PRIVATE_TOKEN")
	projectID := os.Getenv("GITLAB_PROJECT_ID")

	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	if err != nil {
		log.Panic("Error initiating Linebot client")
	}

	client := client.New(projectID, privateToken)
	handler := handler.New(client, bot)

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", handler.Ping)

	r.Run()
}
