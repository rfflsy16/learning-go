package service

import (
    "rest-api-go/internal/module/user/entity"

    "gorm.io/gorm"
)

type UserService struct {
    db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
    return &UserService{db}
}

func (s *UserService) Create(user *entity.User) error {
    if err := user.Validate(); err != nil {
        return err
    }
    return s.db.Create(user).Error
}

func (s *UserService) GetByID(id uint) (*entity.User, error) {
    var user entity.User
    err := s.db.First(&user, id).Error
    return &user, err
}

func (s *UserService) GetAll() ([]entity.User, error) {
    var users []entity.User
    err := s.db.Find(&users).Error
    return users, err
}

func (s *UserService) Update(user *entity.User) error {
    if err := user.Validate(); err != nil {
        return err
    }

    // Cek apakah user ada
    var existingUser entity.User
    if err := s.db.First(&existingUser, user.ID).Error; err != nil {
        return err
    }

    return s.db.Save(user).Error
}

func (s *UserService) Delete(id uint) error {
    return s.db.Delete(&entity.User{}, id).Error
}
