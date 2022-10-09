package user_repository

import (
	"andalas-care/internal/entity/user_entity"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		DB: db,
	}
}

func (u *userRepository) CreateUser(data user_entity.User) (user_entity.User, error) {
	err := u.DB.Create(&data).Error

	return data, err
}

func (u *userRepository) GetUsers(where any) (result []user_entity.User, err error) {
	err = u.DB.Where(where).Find(&result).Error

	return
}

func (u *userRepository) Update(where, value any) error {
	err := u.DB.Where(where).Updates(value).Error

	return err
}
