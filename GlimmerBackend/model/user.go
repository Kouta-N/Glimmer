package model

import (
	"database/sql"
	"time"
)

type User struct{
    ID           int64 `json:"id" "gorm:"primaryKey"`
    Name         string `json:"name"`
    Email        string `json:"email" "gorm:"unique"`
    Password     string `json:"password"`
    ProfileImage sql.NullString `json:profileImage`
    Plan         string `json:"plan"`
    UpdatedAt    *time.Time `json:"updated_at"`
    CreatedAt    *time.Time `json:"created_at"`
}

type UserResponse struct{
    ID    int64 `json:"id" "gorm:"primaryKey"`
    Email string `json:"email" "gorm:"unique"`
}