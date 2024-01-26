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

func (b *basketProductRepository) GetProductInBasket(req *domain.BasketProduct) ([]domain.BasketProduct, error) {
	var basketProducts []domain.BasketProduct

	tx := b.DB.Preload(clause.Associations).Where("basket_id =?", req.BasketID).Find(&basketProducts)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return basketProducts, nil
}

func (b *basketProductRepository) GetProductById(req *domain.BasketProduct, id uint) (*domain.BasketProductPrice, error) {
	var pro domain.Product
	tx := b.DB.Find(&pro, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	product := &domain.BasketProductPrice{
		BasketID: id,
		// Product:  []domain.Product{pro},
		Name:     pro.Name,
		Detail:   pro.Detail,
		Price:    pro.Price,
		Quantity: req.Quantity,
	}

	return product, nil
}

func (b *basketProductRepository) GetPromotionBasket(bp *domain.BasketProduct, id uint) (*domain.BasketPromotionReply, error) {
	var pro domain.Promotion
	tx := b.DB.Find(&pro, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	promotion := &domain.BasketPromotionReply{
		PromotionID: id,
		Code:        pro.Code,
		Discount:    pro.Discount,
		Name:        pro.Name,
		ProductID:   bp.ProductID,
	}

	return promotion, nil
}

func (b *basketProductRepository) GetPromotionByBasketId(bp *domain.BasketProduct) (uint, error) {
	var pro domain.Promotion
	tx := b.DB.Find(&pro, bp.BasketID)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return 0, tx.Error
	}

	return pro.ID, nil
}
