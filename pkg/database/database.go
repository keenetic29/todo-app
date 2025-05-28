package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"todo-app/internal/domain"
)

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.AutoMigrate(&domain.User{}, &domain.Category{}, &domain.Task{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate models: %w", err)
	}

	return db, nil
}