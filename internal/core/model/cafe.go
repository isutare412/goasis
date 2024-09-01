package model

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Cafe struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	Name       string `gorm:"not null; size:256"`
	Location   string `gorm:"not null; check:location <> ''"`
	ReporterID *int64
}

var _ schema.Tabler = &Cafe{}

func (*Cafe) TableName() string { return "cafes" }
