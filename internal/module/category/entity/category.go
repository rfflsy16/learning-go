package entity

import (
    "rest-api-go/internal/module/product/entity"
    "time"
    
    "github.com/go-playground/validator/v10"
)

type Category struct {
    ID          uint                `json:"id" gorm:"primaryKey"`
    Name        string              `json:"name" binding:"max=255"`
    Products    []entity.Product    `json:"products,omitempty" gorm:"foreignKey:CategoryID"`
    CreatedAt   time.Time           `json:"created_at"`
    UpdatedAt   time.Time           `json:"updated_at"`
}

func (p *Category) Validate() error {
    validate := validator.New()
    return validate.Struct(p)
}
