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

func (b *basketProductRepository) AddProductInBasket(req *domain.BasketProduct) error {
	tx := b.DB.Create(req)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (b *basketProductRepository) EditProductInBasket(req *domain.BasketProduct) error {
	tx := b.DB.Model(&domain.BasketProduct{}).Where("basket_id=?", req.BasketID).Updates(req)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (b *basketProductRepository) DeleteProductInBasket(req *domain.BasketProduct) error {

	tx := b.DB.Where("basket_id=?", req.BasketID).Delete(req)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (b *basketProductRepository) GetProductInBasket(req *domain.BasketProduct) (*domain.BasketProductReply, error) {

	tx := b.DB.Preload(clause.Associations).Where("basket_id =?", req.BasketID).First(req)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	basket := &domain.BasketProductReply{
		BasketID:  req.BasketID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	return basket, nil
}
