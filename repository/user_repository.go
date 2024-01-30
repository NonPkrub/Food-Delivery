package repository

import (
	"Food-delivery/domain"
	"fmt"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) domain.UserRepository {
	return &userRepository{DB: DB}
}

func (ur *userRepository) CreateUser(form *domain.User) (*domain.User, error) {
	tx := ur.DB.Begin()
	if err := tx.Create(form).Error; err != nil {
		tx.Rollback()
		fmt.Println(err)
		return nil, err
	}

	var req *domain.Basket
	if err := tx.Where("user_id = ?", form.ID).Where("promotion_id = ?", 0).Create(req).Error; err != nil {
		tx.Rollback()
		fmt.Println(err)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return form, nil
}

func (ur *userRepository) FindOne(form *domain.User) (*domain.User, error) {
	tx := ur.DB.Where("email =?", form.Email).Find(form)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return form, nil
}

func (ur *userRepository) GetOneByID(form *domain.User) (*domain.User, error) {
	tx := ur.DB.Find(form, form.ID)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return form, nil
}

func (ur *userRepository) GetOne(form *domain.User) (*domain.User, error) {
	tx := ur.DB.Find(form, form.ID)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return form, nil
}
