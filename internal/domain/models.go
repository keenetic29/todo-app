package domain

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Task struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	Completed   bool      `gorm:"default:false" json:"completed"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}