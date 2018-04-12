package handler

import (
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
func New(client *c.Client) *Handler {
	return &Handler{
		client: client,
	}
}

// Ping if a function for handling healthcheck in top level routing
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
