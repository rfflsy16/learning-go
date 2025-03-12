package entity

import (
    "time"
    
    "github.com/go-playground/validator/v10"
)

type Product struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Title       string    `json:"title" binding:"max=255"`
    Price       float64   `json:"price" binding:"numeric"`
    Description string    `json:"description" binding:"max=255"`
    CategoryID  uint      `json:"category_id" gorm:"index"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

func (p *Product) Validate() error {
    validate := validator.New()
    return validate.Struct(p)
}
