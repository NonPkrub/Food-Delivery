package usecase

import (
	"Food-delivery/domain"
)

type promotionProductUseCase struct {
	promotionProductRepo domain.PromotionProductRepository
}

func NewPromotionProductUseCase(promotionProductRepo domain.PromotionProductRepository) domain.PromotionProductUseCase {
	return &promotionProductUseCase{promotionProductRepo: promotionProductRepo}
}

func (p *promotionProductUseCase) AddPromotionProduct(req *domain.PromotionProductForm) error {

	promotion := &domain.PromotionProduct{
		PromotionID: req.PromotionID,
		ProductID:   req.ProductID,
	}

	err := p.promotionProductRepo.AddPromotionProduct(promotion)
	if err != nil {
		return err
	}

	return nil
}

func (p *promotionProductUseCase) EditPromotionProduct(req *domain.PromotionProductForm) error {
	promotion := &domain.PromotionProduct{
		PromotionID: req.PromotionID,
		ProductID:   req.ProductID,
	}

	err := p.promotionProductRepo.EditPromotionProduct(promotion)
	if err != nil {
		return err
	}

	return nil
}

func (p *promotionProductUseCase) GetPromotionProduct(req *domain.PromotionProductForm) ([]domain.PromotionProductReply, error) {
	promotion := &domain.PromotionProduct{
		PromotionID: req.PromotionID,
		ProductID:   req.ProductID,
	}

	product, err := p.promotionProductRepo.GetPromotionProduct(promotion)
	if err != nil {
		return nil, err
	}

	dp := []domain.Product{}

	promotions := []domain.PromotionProductReply{}
	for _, pro := range product {

		products, err := p.promotionProductRepo.GetProductById(promotion, pro.ProductID)
		if err != nil {
			return nil, err
		}

		promotions = append(promotions, domain.PromotionProductReply{
			PromotionID: pro.PromotionID,
			Product: append(dp, domain.Product{
				Name:   products.Name,
				Detail: products.Detail,
				Price:  products.Price,
			}),
		})
	}

	return promotions, nil
}
