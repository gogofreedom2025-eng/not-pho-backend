package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type PingResponse struct {
	Pong       bool   `json:"pong"`
	ServerTime string `json:"serverTime"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

var messages = []string{
	"Your code compiles on the first try. Probably.",
	"Legends say you never need Stack Overflow.",
	"That idea? It's good. Ship it.",
	"You make merge conflicts nervous.",
	"Production deploys ask for your permission first.",
	"Git blame always points somewhere else.",
	"Your rubber duck takes notes from you.",
	"CI pipelines green-light themselves when you're online.",
	"Null pointers know better than to visit your code.",
	"Hotfixes cool down before they reach your branch.",
	"The linter has no notes for you today.",
	"Race conditions yield to you in traffic.",
	"Bugs file tickets against themselves.",
	"Your keyboard autocomplete finishes your thoughts.",
	"Semgrep runs out of things to say.",
	"Stack traces apologize before they print.",
	"Someone out there is glad you exist.",
	"Not pho, and that's good.",
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{http.MethodGet, http.MethodOptions},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	api := r.Group("/api")
	{
		api.GET("/ping", handlePing)
		api.GET("/message", handleMessage)
	}

	r.Run(":8081")
}

func handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, PingResponse{
		Pong:       true,
		ServerTime: time.Now().UTC().Format(time.RFC3339),
	})
}

func handleMessage(c *gin.Context) {
	c.JSON(http.StatusOK, MessageResponse{
		Message: messages[rand.Intn(len(messages))],
	})
}
