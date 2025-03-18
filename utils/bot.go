package utils

import (
	"fmt"
	"github.com/ShohruzNuraddinov/go-menu-bot/models"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func DeleteLastMessage(b *gotgbot.Bot, message *gotgbot.Message) error {
	opts := &gotgbot.DeleteMessageOpts{}
	_, err := message.Delete(b, opts)
	if err != nil {
		return fmt.Errorf("failed to delete last message: %w", err)
	}
	return nil
}

func SendMessage(b *gotgbot.Bot, user *models.TelegramUser, text string, opts *gotgbot.SendMessageOpts) error {
	_, err := b.SendMessage(int64(user.ID), text, opts)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}
