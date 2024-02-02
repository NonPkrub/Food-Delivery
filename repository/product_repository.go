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

func (pr *productRepository) GetAll() ([]domain.Product, error) {
	product := []domain.Product{}
	tx := pr.DB.Find(product)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return product, nil

}

func (pr *productRepository) Create(form *domain.Product) (*domain.Product, error) {
	tx := pr.DB.Create(form)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return form, nil
}

func (pr *productRepository) Edit(form *domain.Product) (*domain.Product, error) {
	tx := pr.DB.Model(&domain.Product{}).Where("id=?", form.ID).Updates(form)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return form, nil
}

func (pr *productRepository) Delete(form *domain.Product) error {
	var product domain.Product
	tx := pr.DB.Where("id=?", form.ID).Delete(&product)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (pr *productRepository) GetOneByID(form *domain.Product) (*domain.Product, error) {
	var product domain.Product
	tx := pr.DB.First(&product, form.ID)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return form, nil
}
