package utils

import (
	"fmt"
	"github.com/ShohruzNuraddinov/go-menu-bot/config"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type TelegramUser struct {
	ID        int64
	FirstName string
	LastName  string
	UserName  string
	FullName  string
}

func (u *TelegramUser) GetFullName(ctx *ext.Context) string {
	user := ctx.EffectiveUser
	if user.LastName == "" {
		return user.FirstName
	}
	return fmt.Sprintf("%s %s", user.FirstName, user.LastName)
}

func (u *TelegramUser) GetUserData(ctx *ext.Context) TelegramUser {
	user := ctx.EffectiveUser
	u.ID = user.Id
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.UserName = user.Username
	u.FullName = u.GetFullName(ctx)
	return *u
}

func (u *TelegramUser) Create(user *TelegramUser) (int64, error) {
	db := config.GetDB()
	query := `INSERT INTO users (telegram_id, first_name, last_name, username) VALUES ($1, $2, $3, $4) RETURNING id`

	var id int64
	err := db.QueryRow(query, user.ID, user.FirstName, user.LastName, user.UserName).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to insert/update user: %w", err)
	}
	return id, nil
}

func DeleteLastMessage(b *gotgbot.Bot, message *gotgbot.Message) error {
	opts := &gotgbot.DeleteMessageOpts{}
	_, err := message.Delete(b, opts)
	if err != nil {
		return fmt.Errorf("failed to delete last message: %w", err)
	}
	return nil
}

func SendMessage(b *gotgbot.Bot, user *TelegramUser, text string, opts *gotgbot.SendMessageOpts) error {
	_, err := b.SendMessage(int64(user.ID), text, opts)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}
