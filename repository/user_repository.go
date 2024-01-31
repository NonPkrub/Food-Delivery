package repository

import (
	"Food-delivery/domain"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) domain.UserRepository {
	return &userRepository{DB: DB}
}

func (ur *userRepository) CreateUser(form *domain.User) (*domain.User, error) {
	tx := ur.DB.Create(form)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
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

func (ur *userRepository) GetMe(ID string) (*domain.User, error) {
	var form domain.User
	uintID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		return nil, err
	}

	tx := ur.DB.Where("id =?", uintID).Find(&form)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}
	return &form, nil
}
