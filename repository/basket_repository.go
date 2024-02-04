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
	var basket domain.Basket
	tx := br.DB.Where("user_id =?", form.UserID).Where("promotion_id=?", 0).Create(&basket)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (br *basketRepository) Create(form *domain.Basket) error {
	var basket domain.Basket
	tx := br.DB.Model(&domain.Basket{}).Where("user_id=?", form.UserID).Updates(&basket)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (br *basketRepository) Delete(form *domain.Basket) error {
	var basket domain.Basket
	_ = br.DB.Find(&basket, form.UserID)

	form.PromotionID = 0

	tx := br.DB.Save(&basket)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (br *basketRepository) GetOneByID(form *domain.Basket) (*domain.Basket, error) {
	var basket domain.Basket
	tx := br.DB.Find(&basket, form.UserID)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	if form.PromotionID != 0 {
		basket.PromotionID = form.PromotionID
	}

	return &basket, nil
}
