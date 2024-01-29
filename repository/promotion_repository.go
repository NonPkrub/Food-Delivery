package repository

import (
	"Food-delivery/domain"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (p *promotionRepository) EditPromotion(req *domain.Promotion, id uint) error {
	tx := p.DB.Model(&domain.Promotion{}).Where("id =?", id).Updates(req)
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

func (p *promotionRepository) GetPromotionProduct(req *domain.PromotionProduct) ([]domain.PromotionProduct, error) {
	var promotionProduct []domain.PromotionProduct

	tx := p.DB.Preload(clause.Associations).Where("promotion_id =?", req.PromotionID).Find(&promotionProduct)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return promotionProduct, nil
}

func (p *promotionRepository) GetAllPromotion() ([]domain.Promotion, error) {
	pro := []domain.Promotion{}
	tx := p.DB.Find(&pro)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return pro, nil
}

func (p *promotionRepository) SearchPromotion(req *domain.Promotion) (*domain.Promotion, error) {
	tx := p.DB.Where("code =?", req.Code).Find(req)

	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	promotion := &domain.Promotion{
		Model:    req.Model,
		Code:     req.Code,
		Discount: req.Discount,
		Name:     req.Name,
		Detail:   req.Detail,
	}

	return promotion, nil
}

func (p *promotionRepository) GetProductById(req *domain.PromotionProduct, id uint) (*domain.PromotionProductReplyId, error) {
	var pro domain.Product

	req.ProductID = id
	tx := p.DB.Find(&pro, req.ProductID)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	promotion := &domain.PromotionProductReplyId{
		PromotionID: req.PromotionID,
		Name:        pro.Name,
		Detail:      pro.Detail,
		Price:       pro.Price,
	}

	return promotion, nil

}
