package postgres

import (
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	db *gorm.DB
}

func NewClient(cfg Config) (*Client, error) {
	db, err := gorm.Open(
		postgres.Open(buildDataSourceName(cfg)),
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			TranslateError:                           true,
		})
	if err != nil {
		return nil, fmt.Errorf("opening gorm db: %w", err)
	}

	return &Client{
		db: db,
	}, nil
}

func (c *Client) Initialize(ctx context.Context) error {
	if err := c.db.WithContext(ctx).AutoMigrate(); err != nil {
		return fmt.Errorf("migrating schemas: %w", err)
	}
	return nil
}

func buildDataSourceName(cfg Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)
}
