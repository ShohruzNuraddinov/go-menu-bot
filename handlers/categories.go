package handlers

import (
	"fmt"
	"github.com/ShohruzNuraddinov/go-menu-bot/buttons"
	"github.com/ShohruzNuraddinov/go-menu-bot/states"
	"github.com/ShohruzNuraddinov/go-menu-bot/utils"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func Categories(b *gotgbot.Bot, ctx *ext.Context) error {
	u := utils.TelegramUser{}
	user := u.GetUserData(ctx)

	del_err := utils.DeleteLastMessage(b, ctx.EffectiveMessage)
	if del_err != nil {
		return fmt.Errorf("failed to delete last message: %w", del_err)
	}

	markup := buttons.CategoriesInline()
	_, err := b.SendMessage(user.ID, "Kategoriyalar", &gotgbot.SendMessageOpts{
		ReplyMarkup: &markup,
	})
	if err != nil {
		return fmt.Errorf("failed to send categories message: %w", err)
	}
	return handlers.NextConversationState(states.CATEGORY)
}

func Category(b *gotgbot.Bot, ctx *ext.Context) error {
	u := utils.TelegramUser{}
	user := u.GetUserData(ctx)
	callbackData := ctx.CallbackQuery.Data

	del_err := utils.DeleteLastMessage(b, ctx.EffectiveMessage)
	if del_err != nil {
		return fmt.Errorf("failed to delete last message: %w", del_err)
	}

	if callbackData == "back" {
		markup := buttons.StartInline()
		_, err := b.SendMessage(user.ID, "Bosh sahifa", &gotgbot.SendMessageOpts{
			ReplyMarkup: &markup,
		})
		if err != nil {
			return fmt.Errorf("failed to send start message: %w", err)
		}
		return handlers.NextConversationState(states.CATEGORIES)
	}

	markup := buttons.ProductsInline()
	opts := &gotgbot.SendMessageOpts{
		ReplyMarkup: &markup,
	}
	_, err := b.SendMessage(user.ID, fmt.Sprintf("Kategoriya %v", ctx.CallbackQuery.Data), opts)
	if err != nil {
		return fmt.Errorf("failed to send category message: %w", err)
	}
	return handlers.NextConversationState(states.PRODUCTS)
}
