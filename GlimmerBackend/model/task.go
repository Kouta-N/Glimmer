package model

import (
	"time"
)

type Task struct{
    ID        int64      `json:"id" "gorm:"primaryKey"`
    Title     string     `json:"title"`
    User      User       `json:"email" "gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE"`
    UserId    int64      `json:"user_id" "gorm:"not null"`
    UpdatedAt *time.Time `json:"updated_at"`
    CreatedAt *time.Time `json:"created_at"`
}

type TaskResponse struct{
    ID        int64      `json:"id" "gorm:"primaryKey"`
    Title     string     `json:"title" "gorm:"not null"`
    UpdatedAt *time.Time `json:"updated_at"`
    CreatedAt *time.Time `json:"created_at"`
}