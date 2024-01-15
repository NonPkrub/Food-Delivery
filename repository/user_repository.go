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

func (u *userRepository) SignUp(d *domain.User) (*domain.UserReply, error) {

	tx := u.DB.Create(d)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	user := &domain.UserReply{
		FirstName:     d.FirstName,
		LastName:      d.LastName,
		Email:         d.Email,
		PhoneNumber:   d.Password,
		NationalID:    d.NationalID,
		Address:       d.Address,
		DetailAddress: d.DetailAddress,
	}

	return user, nil

}

func (u *userRepository) Login(d *domain.User) (*domain.UserLoginReply, error) {
	tx := u.DB.Where("email =?", d.Email).Find(d)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	user := &domain.UserLoginReply{
		Password: d.Password,
	}

	return user, nil
}

func (u *userRepository) GetUserById(d *domain.User, id uint) (*domain.UserReply, error) {

	tx := u.DB.First(d, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	user := &domain.UserReply{
		FirstName:     d.FirstName,
		LastName:      d.LastName,
		Email:         d.Email,
		PhoneNumber:   d.Password,
		NationalID:    d.NationalID,
		Address:       d.Address,
		DetailAddress: d.DetailAddress,
	}

	return user, nil
}
