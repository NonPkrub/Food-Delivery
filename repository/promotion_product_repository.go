package repository

import (
	"Food-delivery/domain"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	tx := p.DB.Model(&domain.PromotionProduct{}).Where("promotion_id=?", req.PromotionID).Updates(req)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (p *promotionProductRepository) GetPromotionProduct(req *domain.PromotionProduct) ([]domain.PromotionProduct, error) {
	var promotionProduct []domain.PromotionProduct

	tx := p.DB.Preload(clause.Associations).Where("promotion_id =?", req.PromotionID).Find(&promotionProduct)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return promotionProduct, nil
}
