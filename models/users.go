package models

import (
	"fmt"

	"github.com/ShohruzNuraddinov/go-menu-bot/config"

	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	_ "github.com/lib/pq"
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

func (u *TelegramUser) GetByTelegramID(telegram_id int64) (*TelegramUser, error) {
	db := config.GetDB()
	query := `SELECT * FROM users WHERE telegram_id = $1`

	err := db.QueryRow(query, telegram_id).Scan(&u.ID, &u.FirstName, &u.LastName, &u.UserName)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return u, nil
}

func (u *TelegramUser) GetByID(id int64) (*TelegramUser, error) {
	db := config.GetDB()
	query := `SELECT * FROM users WHERE id = $1`

	err := db.QueryRow(query, id).Scan(&u.ID, &u.FirstName, &u.LastName, &u.UserName)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return u, nil
}

func (u *TelegramUser) Update(user *TelegramUser) error {
	db := config.GetDB()
	query := `UPDATE users SET first_name = $1, last_name = $2, username = $3 WHERE id = $4`

	err := db.QueryRow(query, user.FirstName, user.LastName, user.UserName, user.ID).Scan()
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (u *TelegramUser) Delete(id int64) error {
	db := config.GetDB()
	query := `DELETE FROM users WHERE id = $1`

	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
