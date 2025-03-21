package handlers

import (
	"fmt"

	"github.com/ShohruzNuraddinov/go-menu-bot/buttons"
	"github.com/ShohruzNuraddinov/go-menu-bot/config"
	"github.com/ShohruzNuraddinov/go-menu-bot/models"
	"github.com/ShohruzNuraddinov/go-menu-bot/states"
	"gorm.io/gorm"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func start(b *gotgbot.Bot, ctx *ext.Context) error {
	var user models.TelegramUser
	userData := user.GetUserData(ctx)
	fullName := userData.FullName
	db := config.DB

	if obj := db.First(&user, "telegram_id = ?", userData.TelegramID); obj.Error == gorm.ErrRecordNotFound {
		db.Create(&user)
	} else {
		db.Model(&user).Updates(&userData)
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
