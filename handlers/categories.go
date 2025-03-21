package handlers

import (
	"fmt"
	"strings"

	"github.com/ShohruzNuraddinov/go-menu-bot/buttons"
	"github.com/ShohruzNuraddinov/go-menu-bot/config"
	"github.com/ShohruzNuraddinov/go-menu-bot/models"
	"github.com/ShohruzNuraddinov/go-menu-bot/states"
	"github.com/ShohruzNuraddinov/go-menu-bot/utils"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func Categories(b *gotgbot.Bot, ctx *ext.Context) error {
	var u models.TelegramUser
	user := u.GetUserData(ctx)
	var categories []models.Category

	if err := config.DB.Find(&categories).Error; err != nil {
		return fmt.Errorf("failed to get categories: %w", err)
	}

	del_err := utils.DeleteLastMessage(b, ctx.EffectiveMessage)
	if del_err != nil {
		return fmt.Errorf("failed to delete last message: %w", del_err)
	}

	markup := buttons.CategoriesInline(categories)
	_, err := b.SendMessage(user.TelegramID, "Kategoriyalar", &gotgbot.SendMessageOpts{
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
		_, err := b.SendMessage(user.TelegramID, "Bosh sahifa", &gotgbot.SendMessageOpts{
			ReplyMarkup: &markup,
		})
		if err != nil {
			return fmt.Errorf("failed to send start message: %w", err)
		}
		return handlers.NextConversationState(states.CATEGORIES)
	}
	categoryId := strings.Split(callbackData, "_")[1]
	ctx.Data["category"] = categoryId
	var category models.Category

	if err := config.DB.Where("id = ?", categoryId).First(&category).Error; err != nil {
		return fmt.Errorf("failed to get category: %w", err)
	}

	p := models.Product{}
	var products []models.Product
	if err := config.DB.Model(&p).Where("category_id = ?", categoryId).Find(&products).Error; err != nil {
		return fmt.Errorf("failed to get products: %w", err)
	}
	markup := buttons.ProductsInline(products)
	opts := &gotgbot.SendMessageOpts{
		ReplyMarkup: &markup,
	}
	content := fmt.Sprintf("Kategoriya: %v", category.Name)
	_, err := b.SendMessage(user.TelegramID, content, opts)
	if err != nil {
		return fmt.Errorf("failed to send category message: %w", err)
	}
	return handlers.NextConversationState(states.PRODUCTS)
}
