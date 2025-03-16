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

func Products(b *gotgbot.Bot, ctx *ext.Context) error {
	u := utils.TelegramUser{}
	user := u.GetUserData(ctx)
	callbackData := ctx.CallbackQuery.Data

	del_err := utils.DeleteLastMessage(b, ctx.EffectiveMessage)
	if del_err != nil {
		return fmt.Errorf("failed to delete last message: %w", del_err)
	}

	if callbackData == "back" {
		markup := buttons.CategoriesInline()
		_, err := b.SendMessage(user.ID, "Kategoriyalar", &gotgbot.SendMessageOpts{
			ReplyMarkup: &markup,
		})
		if err != nil {
			return fmt.Errorf("failed to send categories message: %w", err)
		}
		return handlers.NextConversationState(states.CATEGORY)
	}

	markup := buttons.BackInline()
	opts := &gotgbot.SendMessageOpts{
		ReplyMarkup: &markup,
	}
	_, err := b.SendMessage(user.ID, fmt.Sprintf("Mahsulotlar %v", ctx.CallbackQuery.Data), opts)
	if err != nil {
		return fmt.Errorf("failed to send products message: %w", err)
	}
	return handlers.NextConversationState(states.PRODUCTS)
}
