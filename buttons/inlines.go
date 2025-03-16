package buttons

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
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

func CategoriesInline() gotgbot.InlineKeyboardMarkup {
	btn := gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{
				{Text: "Category 1", CallbackData: "category_1"},
			},
			{
				{Text: "Category 2", CallbackData: "category_2"},
			},
			{
				{Text: "Orqaga", CallbackData: "back"},
			},
		},
	}
	return btn
}

func ProductsInline() gotgbot.InlineKeyboardMarkup {
	btn := gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{
				{Text: "Product 1", CallbackData: "products_1"},
			},
			{
				{Text: "Product 2", CallbackData: "products_2"},
			},
			{
				{Text: "Orqaga", CallbackData: "back"},
			},
		},
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
