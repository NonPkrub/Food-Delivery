package repository

import (
	"Food-delivery/domain"
	"fmt"

	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) domain.ProductRepository {
	return &productRepository{DB: DB}
}

func (p *productRepository) GetAll() ([]domain.Product, error) {
	pro := []domain.Product{}
	tx := p.DB.Find(&pro)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return pro, nil

}

func (p *productRepository) AddProduct(pro *domain.Product) (*domain.ProductReply, error) {

	tx := p.DB.Create(pro)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	product := &domain.ProductReply{
		Name:   pro.Name,
		Detail: pro.Detail,
		Price:  pro.Price,
	}

	return product, nil
}

func (p *productRepository) EditProduct(pro *domain.Product, id uint) (*domain.ProductReply, error) {

	tx := p.DB.Model(&domain.Product{}).Where("id=?", id).Updates(pro)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}
	product := &domain.ProductReply{
		Name:   pro.Name,
		Detail: pro.Detail,
		Price:  pro.Price,
	}

	return product, nil
}

func (p *productRepository) DeleteProduct(pro *domain.Product, id uint) error {
	tx := p.DB.Where("id=?", id).Delete(pro)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (p *productRepository) GetProductById(pro *domain.Product, id uint) (*domain.ProductReply, error) {

	tx := p.DB.First(pro, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	product := &domain.ProductReply{
		Name:   pro.Name,
		Detail: pro.Detail,
		Price:  pro.Price,
	}

	return product, nil
}
