package entity

import (
    "time"
    
    "github.com/go-playground/validator/v10"
)

type User struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Username        string    `json:"username" binding:"max=255"`
    Email        string    `json:"email" binding:"max=255"`
    Password        string    `json:"password" binding:"max=255"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

func (p *User) Validate() error {
    validate := validator.New()
    return validate.Struct(p)
}
