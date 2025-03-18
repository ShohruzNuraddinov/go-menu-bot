package handlers

import (
	"fmt"

	"github.com/ShohruzNuraddinov/go-menu-bot/buttons"
	"github.com/ShohruzNuraddinov/go-menu-bot/models"
	"github.com/ShohruzNuraddinov/go-menu-bot/states"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func start(b *gotgbot.Bot, ctx *ext.Context) error {
	user := models.TelegramUser{}
	userData := user.GetUserData(ctx)
	fullName := userData.FullName
	if _, err := user.GetByTelegramID(userData.ID); err != nil {
		user.Create(&userData)
	}

	markup := buttons.StartInline()
	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Assalomu aleykum, Xush kelibsiz, %s", fullName), &gotgbot.SendMessageOpts{
		ReplyMarkup: &markup,
	})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return handlers.NextConversationState(states.CATEGORIES)
}

func cancel(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, "Oh, goodbye!", nil)
	if err != nil {
		return fmt.Errorf("failed to send cancel message: %w", err)
	}
	return handlers.EndConversation()
}
