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

func Products(b *gotgbot.Bot, ctx *ext.Context) error {
	var u models.TelegramUser
	user := u.GetUserData(ctx)
	callbackData := ctx.CallbackQuery.Data

	if err := utils.DeleteLastMessage(b, ctx.EffectiveMessage); err != nil {
		return fmt.Errorf("failed to delete last message: %w", err)
	}

	if callbackData == "back" {
		var categories []models.Category

		if err := config.DB.Model(&models.Category{}).Preload("Products").Find(&categories).Error; err != nil {
			return fmt.Errorf("failed to get categories: %w", err)
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

	productId := strings.Split(callbackData, "_")[1]
	markup := buttons.BackInline()
	opts := &gotgbot.SendMessageOpts{
		ReplyMarkup: &markup,
	}

	var product models.Product
	if err := config.DB.Where("id = ?", productId).First(&product).Error; err != nil {
		return fmt.Errorf("failed to get product: %w", err)
	}

	content := fmt.Sprintf("Nomi: %v\nNarx: %v\n\nHaqida: %v", product.Name, product.Price, product.Description)

	_, err := b.SendMessage(user.TelegramID, content, opts)
	if err != nil {
		return fmt.Errorf("failed to send products message: %w", err)
	}
	return handlers.NextConversationState(states.PRODUCTS)
}
