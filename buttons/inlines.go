package buttons

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/ShohruzNuraddinov/go-menu-bot/models"
)

func StartInline() gotgbot.InlineKeyboardMarkup {
	btn := gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{
				{Text: "Kategoriyalar", CallbackData: "categories"},
			},
		},
	}
	return btn
}

func CategoriesInline(categories []models.Category) gotgbot.InlineKeyboardMarkup {
	var btns [][]gotgbot.InlineKeyboardButton
	for _, category := range categories {
		btns = append(btns, []gotgbot.InlineKeyboardButton{
			{Text: category.Name, CallbackData: fmt.Sprintf("category_%d", category.ID)},
		})
	}
	btns = append(btns, []gotgbot.InlineKeyboardButton{
		{Text: "Orqaga", CallbackData: "back"},
	})
	btn := gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: btns,
	}
	return btn
}

func ProductsInline(products []models.Product) gotgbot.InlineKeyboardMarkup {
	var btns [][]gotgbot.InlineKeyboardButton

	for _, product := range products {
		btns = append(btns, []gotgbot.InlineKeyboardButton{
			{Text: product.Name, CallbackData: fmt.Sprintf("products_%d", product.ID)},
		})
	}

	btns = append(btns, []gotgbot.InlineKeyboardButton{
		{Text: "Orqaga", CallbackData: "back"},
	})

	btn := gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: btns,
	}
	return btn
}

func BackInline() gotgbot.InlineKeyboardMarkup {
	btn := gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{
				{Text: "Orqaga", CallbackData: "back"},
			},
		},
	}
	return btn
}
