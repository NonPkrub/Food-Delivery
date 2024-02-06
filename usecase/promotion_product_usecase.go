package usecase

import (
	"Food-delivery/domain"
)

type promotionProductUseCase struct {
	promotionProductRepo domain.PromotionProductRepository
	productRepo          domain.ProductRepository
}

func NewPromotionProductUseCase(promotionProductRepo domain.PromotionProductRepository, productRepo domain.ProductRepository) domain.PromotionProductUseCase {
	return &promotionProductUseCase{promotionProductRepo: promotionProductRepo, productRepo: productRepo}
}

func (uc *promotionProductUseCase) AddPromotionProduct(form *domain.PromotionProductForm) error {
	promotion := &domain.PromotionProduct{
		PromotionID: form.PromotionID,
		ProductID:   form.ProductID,
	}

	err := uc.promotionProductRepo.Create(promotion)
	if err != nil {
		return err
	}

	return nil
}

func (uc *promotionProductUseCase) EditPromotionProduct(form *domain.PromotionProductForm) error {
	promotion := &domain.PromotionProduct{
		PromotionID: form.PromotionID,
		ProductID:   form.ProductID,
	}

	err := uc.promotionProductRepo.Edit(promotion)
	if err != nil {
		return err
	}

	return nil
}
