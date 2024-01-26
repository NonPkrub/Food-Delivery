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

func (b *basketRepository) CreateBasket(req *domain.Basket) error {
	tx := b.DB.Where("user_id =?", req.UserID).Where("promotion_id=?", 0).Create(req)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (b *basketRepository) AddPromotionBasket(req *domain.Basket) error {
	tx := b.DB.Model(&domain.Basket{}).Where("user_id=?", req.UserID).Updates(req)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (b *basketRepository) DeletePromotionBasket(req *domain.Basket) error {
	tx := b.DB.Find(req, req.UserID)

	req.PromotionID = nil

	tx = b.DB.Save(req)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (b *basketRepository) GetBasketByUserId(req *domain.Basket) (*domain.BasketReply, error) {
	tx := b.DB.Find(req, req.UserID)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	basket := &domain.BasketReply{
		Model:  req.Model,
		UserID: req.UserID,
	}

	if req.PromotionID != nil {
		basket.PromotionID = *req.PromotionID
	}

	return basket, nil
}
