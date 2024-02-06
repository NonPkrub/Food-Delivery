package repository

import (
	"Food-delivery/domain"
	"fmt"

	"gorm.io/gorm"
)

type basketRepository struct {
	DB *gorm.DB
}

func NewBasketRepository(DB *gorm.DB) domain.BasketRepository {
	return &basketRepository{DB: DB}
}

func (br *basketRepository) CreateOne(form *domain.Basket) error {
	tx := br.DB.Create(form)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (br *basketRepository) Create(form *domain.Basket) error {
	tx := br.DB.Model(&domain.Basket{}).Where("user_id=?", form.UserID).Updates(form)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (br *basketRepository) Delete(form *domain.Basket) error {
	tx := br.DB.Model(&domain.Basket{}).Where("user_id=?", form.UserID).Update("promotion_id", nil)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (br *basketRepository) GetOneByID(form *domain.Basket) (*domain.Basket, error) {
	var basket domain.Basket
	tx := br.DB.Where("user_id=?", form.UserID).Find(&basket)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	if form.PromotionID != nil {
		basket.PromotionID = form.PromotionID
	}

	return &basket, nil
}
