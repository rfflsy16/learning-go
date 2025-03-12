package entity

import (
    "time"
    
    "github.com/go-playground/validator/v10"
)

type Category struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Name        string    `json:"name" binding:"max=255"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

func (p *Category) Validate() error {
    validate := validator.New()
    return validate.Struct(p)
}
