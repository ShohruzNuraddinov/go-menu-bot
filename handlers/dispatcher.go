package handlers

import (
	"menubot/states"

	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/conversation"

	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
)

func Dispatcher() *ext.Dispatcher {
	dispatcher := ext.NewDispatcher(nil)

	dispatcher.AddHandler(handlers.NewConversation(
		[]ext.Handler{handlers.NewCommand("start", start)},
		map[string][]ext.Handler{
			states.CATEGORIES: {
				handlers.NewCallback(callbackquery.Equal("categories"), Categories),
				handlers.NewCallback(callbackquery.Equal("back"), Categories),
			},
			states.CATEGORY: {
				handlers.NewCallback(callbackquery.Prefix("category_"), Category),
				handlers.NewCallback(callbackquery.Equal("back"), Category),
			},
			states.PRODUCTS: {
				handlers.NewCallback(callbackquery.Prefix("products_"), Products),
				handlers.NewCallback(callbackquery.Equal("back"), Products),
			},
		},
		&handlers.ConversationOpts{
			Exits: []ext.Handler{
				handlers.NewCommand("cancel", cancel),
			},
			StateStorage: conversation.NewInMemoryStorage(
				conversation.KeyStrategySenderAndChat,
			),
			AllowReEntry: true,
		},
	))

	return dispatcher
}
