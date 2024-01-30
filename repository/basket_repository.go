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

// func (b *basketRepository) CreateBasket(req *domain.Basket) error {
// 	tx := b.DB.Where("user_id =?", req.UserID).Where("promotion_id=?", 0).Create(req)
// 	if tx.Error != nil {
// 		fmt.Println(tx.Error)
// 		return tx.Error
// 	}

// 	return nil
// }

func (br *basketRepository) Create(form *domain.Basket) error {
	tx := br.DB.Model(&domain.Basket{}).Where("user_id=?", form.UserID).Updates(form)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (br *basketRepository) Delete(form *domain.Basket) error {
	tx := br.DB.Find(form, form.UserID)

	form.PromotionID = nil

	tx = br.DB.Save(form)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (br *basketRepository) GetOneByID(form *domain.Basket) (*domain.Basket, error) {
	tx := br.DB.Find(form, form.UserID)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	if form.PromotionID != nil {
		form.PromotionID = form.PromotionID
	}

	return form, nil
}
