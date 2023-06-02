package model

import (
	"time"
)

type UserType int

const (
	Individual UserType = iota
	LegalEntity
	UnknownUserType
)

type User struct {
	ID               int64      `db:"id"`
	Firstname        *string    `db:"firstname"`
	Lastname         *string    `db:"lastname"`
	Birthday         *time.Time `db:"birthday"`
	Phone            string     `db:"phone"`
	Description      *string    `db:"description"`
	Email            *string    `db:"email"`
	Telegram         *string    `db:"telegram"`
	Whatsapp         *string    `db:"whatsapp"`
	Viber            *string    `db:"viber"`
	Instagram        *string    `db:"instagram"`
	PhotoURL         *string    `db:"photo_url"`
	Type             *UserType  `db:"type"`
	RegistrationDate *time.Time `db:"registration_date"`
	Rating           *float32   `db:"rating"`
	Login            *string    `db:"login"`
	Password         *string    `db:"password"`
}
