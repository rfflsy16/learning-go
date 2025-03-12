package service

import (
    "rest-api-go/internal/module/product/entity"

    "gorm.io/gorm"
)

type ProductService struct {
    db *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
    return &ProductService{db}
}

func (s *ProductService) Create(product *entity.Product) error {
    if err := product.Validate(); err != nil {
        return err
    }
    
    // Verify that the category exists
    var count int64
    if err := s.db.Model(&entity.Product{}).Where("id = ?", product.CategoryID).Count(&count).Error; err != nil {
        return err
    }
    
    return s.db.Create(product).Error
}

func (s *ProductService) GetByID(id uint) (*entity.Product, error) {
    var product entity.Product
    err := s.db.First(&product, id).Error
    return &product, err
}

func (s *ProductService) GetAll() ([]entity.Product, error) {
    var products []entity.Product
    err := s.db.Find(&products).Error
    return products, err
}

func (s *ProductService) Update(product *entity.Product) error {
    if err := product.Validate(); err != nil {
        return err
    }

    // Cek apakah product ada
    var existingProduct entity.Product
    if err := s.db.First(&existingProduct, product.ID).Error; err != nil {
        return err
    }

    return s.db.Save(product).Error
}

func (s *ProductService) Delete(id uint) error {
    return s.db.Delete(&entity.Product{}, id).Error
}

func (s *ProductService) GetByCategoryID(categoryID uint) ([]entity.Product, error) {
    var products []entity.Product
    err := s.db.Where("category_id = ?", categoryID).Find(&products).Error
    return products, err
}
