package main

import (
	"encoding/json"
	"log"
	"net/http"

	"menubot/config"
	"menubot/handlers"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/gin-gonic/gin"
)

func main() {
	token := "6738571788:AAEPdZ6Iu1C-79202rnv6vTKhFiduYqt9kU"
	webhookURL := "https://go-bot.jprq.site/webhook"
	port := "8081"
	if token == "" || webhookURL == "" || port == "" {
		panic("TOKEN, WEBHOOK_URL, or PORT environment variables are empty")
	}

	b, err := gotgbot.NewBot(token, nil)
	if err != nil {
		panic("failed to create bot: " + err.Error())
	}

	dispatcher := handlers.Dispatcher()

	config.InitDB()
	db := config.GetDB()

	r := gin.Default()

	r.POST("/webhook", func(c *gin.Context) {
		var update gotgbot.Update
		if err := json.NewDecoder(c.Request.Body).Decode(&update); err != nil {
			log.Println("failed to parse update:", err)
			c.Status(http.StatusBadRequest)
			return
		}

		// Process update in a goroutine
		go dispatcher.ProcessUpdate(b, &update, map[string]interface{}{})
		c.Status(http.StatusOK)
	})

	// Set webhook
	_, err = b.SetWebhook(webhookURL, nil)
	if err != nil {
		panic("failed to set webhook: " + err.Error())
	}

	log.Printf("Bot %s is running with webhook...", b.User.Username)
	r.Run(":" + port)
}
