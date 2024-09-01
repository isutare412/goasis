package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	Nickname    string  `gorm:"not null; size:64; check:nickname <> ''"`
	FamilyName  *string `gorm:"size:128"`
	GivenName   *string `gorm:"size:128"`
	DateOfBirth *time.Time
}
