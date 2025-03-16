package handlers

import (
	"fmt"
	"menubot/buttons"
	"menubot/states"
	"menubot/utils"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func start(b *gotgbot.Bot, ctx *ext.Context) error {
	user := utils.TelegramUser{}
	userData := user.GetUserData(ctx)
	fullName := userData.FullName
	if _, err := user.Create(&userData); err != nil {
		return fmt.Errorf("failed to insert/update user: %w", err)
	}

	markup := buttons.StartInline()
	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Ассалому алейкум Хуш Келибсиз, %s", fullName), &gotgbot.SendMessageOpts{
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
