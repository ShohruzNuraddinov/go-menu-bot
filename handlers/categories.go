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

func Categories(b *gotgbot.Bot, ctx *ext.Context) error {
	u := models.TelegramUser{}
	user := u.GetUserData(ctx)
	c := models.Category{}
	categories, _ := c.GetCategories()

	del_err := utils.DeleteLastMessage(b, ctx.EffectiveMessage)
	if del_err != nil {
		return fmt.Errorf("failed to delete last message: %w", del_err)
	}

	markup := buttons.CategoriesInline(categories)
	_, err := b.SendMessage(user.ID, "Kategoriyalar", &gotgbot.SendMessageOpts{
		ReplyMarkup: &markup,
	})
	if err != nil {
		return fmt.Errorf("failed to send categories message: %w", err)
	}
	return handlers.NextConversationState(states.CATEGORY)
}

func Category(b *gotgbot.Bot, ctx *ext.Context) error {
	u := models.TelegramUser{}
	user := u.GetUserData(ctx)
	callbackData := ctx.CallbackQuery.Data
	if err := utils.DeleteLastMessage(b, ctx.EffectiveMessage); err != nil {
		return fmt.Errorf("failed to delete last message: %w", err)
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
	categoryId := strings.Split(callbackData, "_")[1]
	ctx.Data["category"] = categoryId
	c := models.Category{}
	category, _ := c.GetCategoryByID(categoryId)

	p := models.Product{}
	products, _ := p.GetProductsByCategory(categoryId)
	markup := buttons.ProductsInline(products)
	opts := &gotgbot.SendMessageOpts{
		ReplyMarkup: &markup,
	}
	content := fmt.Sprintf("Kategoriya: %v", category.Name)
	_, err := b.SendMessage(user.ID, content, opts)
	if err != nil {
		return fmt.Errorf("failed to send category message: %w", err)
	}
	return handlers.NextConversationState(states.PRODUCTS)
}
