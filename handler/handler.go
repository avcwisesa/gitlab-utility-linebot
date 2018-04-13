package handler

import (
	"fmt"
	"log"
	"regexp"

	c "github.com/avcwisesa/gitlab-reference-linebot/client"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

// Handler holds the structure for Handler
type Handler struct {
	client *c.Client
	bot    *linebot.Client
}

// New is a function for creating handler
func New(client *c.Client, bot *linebot.Client) *Handler {
	return &Handler{
		client: client,
		bot:    bot,
	}
}

// Ping is a function for handling healthcheck in top level routing
func (h *Handler) Ping(ctx *gin.Context) {
	select {
	case <-ctx.Request.Context().Done():
		ctx.JSON(408, nil)
		return
	default:
	}

	resp := "Ping!"
	ctx.JSON(200, resp)

	return
}

// MessageHandler is a function for handling LINE webhook
func (h *Handler) MessageHandler(ctx *gin.Context) {
	select {
	case <-ctx.Request.Context().Done():
		ctx.JSON(408, nil)
		return
	default:
	}

	events, err := h.bot.ParseRequest(ctx.Request)
	if err != nil {
		log.Println(err)
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			log.Println(event.Message)
			switch message := event.Message.(type) {
			case *linebot.TextMessage:

				issueRegex := regexp.MustCompile(`#\d+`)
				mergeRequestRegex := regexp.MustCompile(`!\d+`)

				issuesFound := issueRegex.FindAllString(message.Text, -1)
				mergeRequestFound := mergeRequestRegex.FindAllString(message.Text, -1)

				log.Printf("Issue(s): %v\n", issuesFound)
				log.Printf("Merge Request(s): %v\n", mergeRequestFound)

				if len(issuesFound) != 0 {
					reply := "Issue(s)\n"

					for _, issue := range issuesFound {
						gitlabIssue, err := h.client.GetIssue(issue[2:])
						if err != nil {
							log.Println(err)
						}
						reply += fmt.Sprintf("[%s]\nLink: %s\n", gitlabIssue.Title, gitlabIssue.WebURL)
					}

					if _, err = h.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
						log.Print(err)
					}
				}

				if len(mergeRequestFound) != 0 {
					reply := "Merge Request(s)\n"

					for _, mr := range mergeRequestFound {
						gitlabMR, err := h.client.GetMergeRequest(mr[2:])
						if err != nil {
							log.Println(err)
						}
						reply += fmt.Sprintf("[%s]\nLink: %s\n", gitlabMR.Title, gitlabMR.WebURL)
					}

					if _, err = h.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	}
}
