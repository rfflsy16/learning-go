package service

import (
    "rest-api-go/internal/module/category/entity"

    "gorm.io/gorm"
)

type CategoryService struct {
    db *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryService {
    return &CategoryService{db}
}

func (s *CategoryService) Create(category *entity.Category) error {
    if err := category.Validate(); err != nil {
        return err
    }
    return s.db.Create(category).Error
}

func (s *CategoryService) GetByID(id uint) (*entity.Category, error) {
    var category entity.Category
    err := s.db.Preload("Products").First(&category, id).Error
    return &category, err
}

func (s *CategoryService) GetAll() ([]entity.Category, error) {
    var categories []entity.Category
    err := s.db.Preload("Products").Find(&categories).Error
    return categories, err
}

func (s *CategoryService) Update(category *entity.Category) error {
    if err := category.Validate(); err != nil {
        return err
    }

    // Cek apakah category ada
    var existingCategory entity.Category
    if err := s.db.First(&existingCategory, category.ID).Error; err != nil {
        return err
    }

    return s.db.Save(category).Error
}

func (s *CategoryService) Delete(id uint) error {
    return s.db.Delete(&entity.Category{}, id).Error
}
