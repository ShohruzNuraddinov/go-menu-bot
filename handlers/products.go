package handlers

import (
	"fmt"
	"strings"

	"github.com/ShohruzNuraddinov/go-menu-bot/buttons"
	"github.com/ShohruzNuraddinov/go-menu-bot/models"
	"github.com/ShohruzNuraddinov/go-menu-bot/states"
	"github.com/ShohruzNuraddinov/go-menu-bot/utils"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func Products(b *gotgbot.Bot, ctx *ext.Context) error {
	u := models.TelegramUser{}
	user := u.GetUserData(ctx)
	callbackData := ctx.CallbackQuery.Data

	if err := utils.DeleteLastMessage(b, ctx.EffectiveMessage); err != nil {
		return fmt.Errorf("failed to delete last message: %w", err)
	}

	if callbackData == "back" {
		c := models.Category{}
		categories, _ := c.GetCategories()
		markup := buttons.CategoriesInline(categories)
		_, err := b.SendMessage(user.ID, "Kategoriyalar", &gotgbot.SendMessageOpts{
			ReplyMarkup: &markup,
		})
		if err != nil {
			return fmt.Errorf("failed to send categories message: %w", err)
		}
		return handlers.NextConversationState(states.CATEGORY)
	}

	productId := strings.Split(callbackData, "_")[1]
	markup := buttons.BackInline()
	opts := &gotgbot.SendMessageOpts{
		ReplyMarkup: &markup,
	}
	p := models.Product{}
	products, _ := p.GetProductByID(productId)

	content := fmt.Sprintf("Nomi: %v\nNarx: %v\n\nHaqida: %v", products.Name, products.Price, products.Description)

	_, err := b.SendMessage(user.ID, content, opts)
	if err != nil {
		return fmt.Errorf("failed to send products message: %w", err)
	}
	return handlers.NextConversationState(states.PRODUCTS)
}
