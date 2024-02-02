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

func (pr *promotionRepository) Create(form *domain.Promotion) error {
	tx := pr.DB.Create(form)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (pr *promotionRepository) Edit(form *domain.Promotion) error {
	tx := pr.DB.Model(&domain.Promotion{}).Where("id =?", form.ID).Updates(form)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (pr *promotionRepository) Delete(form *domain.Promotion) error {
	tx := pr.DB.Where("id=?", form.ID).Delete(form)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (pr *promotionRepository) GetAllByID(form *domain.PromotionProduct) ([]domain.PromotionProduct, error) {
	var promotionProduct []domain.PromotionProduct

	tx := pr.DB.Preload(clause.Associations).Where("promotion_id =?", form.PromotionID).Find(&promotionProduct)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return promotionProduct, nil
}

func (pr *promotionRepository) GetAll() ([]domain.Promotion, error) {
	promotion := []domain.Promotion{}
	tx := pr.DB.Find(&promotion)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return promotion, nil
}

// func (p *promotionRepository) SearchPromotion(req *domain.Promotion) (*domain.Promotion, error) {
// 	tx := p.DB.Where("code =?", req.Code).Find(req)

// 	if tx.Error != nil {
// 		fmt.Println(tx.Error)
// 		return nil, tx.Error
// 	}

// 	promotion := &domain.Promotion{
// 		Model:    req.Model,
// 		Code:     req.Code,
// 		Discount: req.Discount,
// 		Name:     req.Name,
// 		Detail:   req.Detail,
// 	}

// 	return promotion, nil
// }

func (pr *promotionRepository) GetOneByID(form *domain.PromotionProduct) (*domain.PromotionProduct, error) {
	var promotion domain.Product
	tx := pr.DB.Find(&promotion, form.ProductID)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return form, nil

}

func (pr *promotionRepository) FindOneByID(form *domain.PromotionProduct) (*domain.PromotionProduct, error) {
	var promotion domain.PromotionProduct
	tx := pr.DB.Where("promotion_id=? AND product_id=?", form.PromotionID, form.ProductID).Find(&promotion)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return &promotion, nil

}

func (pr *promotionRepository) FindOne(form *domain.Promotion) (*domain.Promotion, error) {
	tx := pr.DB.Find(form)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return form, nil
}

func (pr *promotionRepository) GetByQuery(form *domain.Promotion) (*domain.Promotion, error) {
	tx := pr.DB.Where("code =? OR name =?", form.Code, form.Name).Find(form)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return form, nil
}
