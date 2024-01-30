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

func (ppr *promotionProductRepository) Create(form *domain.PromotionProduct) error {
	tx := ppr.DB.Create(form)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (ppr *promotionProductRepository) Edit(form *domain.PromotionProduct) error {
	tx := ppr.DB.Model(&domain.PromotionProduct{}).Where("promotion_id=?", form.PromotionID).Updates(form)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (ppr *promotionProductRepository) FindAllByID(form *domain.PromotionProduct) ([]domain.PromotionProduct, error) {
	var promotionProduct []domain.PromotionProduct
	tx := ppr.DB.Preload(clause.Associations).Where("promotion_id =?", form.PromotionID).Find(&promotionProduct)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return promotionProduct, nil
}

func (ppr *promotionProductRepository) GetOneByID(form *domain.PromotionProduct) (*domain.PromotionProduct, error) {
	var product domain.Product
	tx := ppr.DB.Find(&product, form.ProductID)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return form, nil
}
