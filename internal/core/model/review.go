package model

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	Score   int32   `gorm:"not null"` // 0 ~ 100
	Message *string `gorm:"size:2048"`

	CafeID int64 `gorm:"not null; index:,type:hash"`

	UserID int64 `gorm:"not null; index:,type:hash"`
	User   *User
}
