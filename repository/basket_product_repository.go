package repository

import (
	"Food-delivery/domain"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type basketProductRepository struct {
	DB *gorm.DB
}

func NewBasketProductRepository(DB *gorm.DB) domain.BasketProductRepository {
	return &basketProductRepository{DB: DB}
}

func (br *basketProductRepository) Create(form *domain.BasketProduct) error {
	tx := br.DB.Create(form)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (br *basketProductRepository) Edit(form *domain.BasketProduct) error {
	tx := br.DB.Model(&domain.BasketProduct{}).Where("basket_id=?", form.BasketID).Updates(form)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (br *basketProductRepository) Delete(form *domain.BasketProduct) error {
	tx := br.DB.Where("basket_id=?", form.BasketID).Delete(form)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (br *basketProductRepository) FindAllByID(form *domain.BasketProduct) ([]domain.BasketProduct, error) {
	var basketProducts []domain.BasketProduct

	tx := br.DB.Preload(clause.Associations).Where("basket_id =?", form.BasketID).Find(&basketProducts)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return basketProducts, nil
}

func (br *basketProductRepository) GetOneById(form *domain.BasketProduct) (*domain.BasketProduct, error) {
	var product domain.Product
	tx := br.DB.Find(&product, form.ProductID)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return form, nil
}

func (br *basketProductRepository) FindOne(form *domain.BasketProduct) (*domain.BasketProduct, error) {
	tx := br.DB.Find(form, form.BasketID)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return form, nil
}

func (br *basketProductRepository) GetPromotionByBasketID(form *domain.BasketProduct) (uint, error) {
	var basket domain.Basket

	tx := br.DB.Find(&basket, form.BasketID)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return 0, tx.Error
	}

	return *basket.PromotionID, nil
}
