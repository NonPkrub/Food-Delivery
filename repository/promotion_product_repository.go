package repository

import (
	"Food-delivery/domain"
	"fmt"

	"gorm.io/gorm"
)

type promotionProductRepository struct {
	DB *gorm.DB
}

func NewPromotionProductRepository(DB *gorm.DB) domain.PromotionProductRepository {
	return &promotionProductRepository{DB: DB}
}

func (p *promotionProductRepository) AddPromotionProduct(req *domain.PromotionProduct) error {

	tx := p.DB.Create(req)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (p *promotionProductRepository) EditPromotionProduct(req *domain.PromotionProduct) error {

	tx := p.DB.Save(req)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (p *promotionProductRepository) GetPromotionProduct(req *domain.PromotionProduct) (*domain.PromotionProductReply, error) {
	tx := p.DB.Find(req, req.PromotionID)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	promotion := &domain.PromotionProductReply{
		PromotionID: req.PromotionID,
		ProductID:   req.ProductID,
	}
	return promotion, nil
}
