package models

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type TelegramUser struct {
	gorm.Model
	TelegramID int64  `gorm:"unique"`
	FirstName  string `gorm:"size:255"`
	LastName   string `gorm:"size:255"`
	UserName   string `gorm:"size:255"`
	FullName   string `gorm:"size:255"`
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
	u.TelegramID = user.Id
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.UserName = user.Username
	u.FullName = u.GetFullName(ctx)
	return *u
}
