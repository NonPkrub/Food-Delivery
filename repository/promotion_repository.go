package repository

import (
	"Food-delivery/domain"
	"fmt"

	"gorm.io/gorm"
)

type promotionRepository struct {
	DB *gorm.DB
}

func NewPromotionRepository(DB *gorm.DB) domain.PromotionRepository {
	return &promotionRepository{DB: DB}
}

func (p *promotionRepository) CreatePromotion(req *domain.Promotion) error {
	tx := p.DB.Create(req)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (p *promotionRepository) EditPromotion(req *domain.Promotion) error {
	tx := p.DB.Save(req)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (p *promotionRepository) DeletePromotion(req *domain.Promotion, id uint) error {
	tx := p.DB.Where("id=?", id).Delete(req)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (p *promotionRepository) GetPromotionById(req *domain.Promotion, id uint) (*domain.PromotionReply, error) {

	tx := p.DB.First(req, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	promotion := &domain.PromotionReply{
		Code:     req.Code,
		Discount: req.Discount,
		Name:     req.Name,
		Detail:   req.Detail,
	}

	return promotion, nil
}

func (p *promotionRepository) GetAllPromotion() ([]domain.Promotion, error) {
	pro := []domain.Promotion{}
	tx := p.DB.Find(pro)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return pro, nil
}
