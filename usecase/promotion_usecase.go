package usecase

import (
	"Food-delivery/domain"
)

type promotionUseCase struct {
	promotionRepo domain.PromotionRepository
}

func NewPromotionUseCase(promotionRepo domain.PromotionRepository) domain.PromotionUseCase {
	return &promotionUseCase{promotionRepo: promotionRepo}
}

func (p *promotionUseCase) CreatePromotion(req *domain.PromotionForm) error {
	promotion := &domain.Promotion{
		Code:     req.Code,
		Discount: req.Discount,
		Name:     req.Name,
		Detail:   req.Detail,
	}

	err := p.promotionRepo.CreatePromotion(promotion)
	if err != nil {
		return err
	}

	return nil
}

func (p *promotionUseCase) EditPromotion(req *domain.PromotionForm, id uint) error {
	promotion := &domain.Promotion{
		Code:     req.Code,
		Discount: req.Discount,
		Name:     req.Name,
		Detail:   req.Detail,
	}

	err := p.promotionRepo.EditPromotion(promotion, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *promotionUseCase) DeletePromotion(id uint) error {
	var promotion domain.Promotion
	err := p.promotionRepo.DeletePromotion(&promotion, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *promotionUseCase) GetPromotionById(id uint) ([]domain.PromotionProductReply, error) {
	// var promotion domain.Promotion
	// promotions, err := p.promotionRepo.GetPromotionById(&promotion, id)
	// if err != nil {
	// 	return nil, err
	// }

	// return promotions, nil
	promotion := &domain.PromotionProduct{
		PromotionID: id,
	}

	product, err := p.promotionRepo.GetPromotionProduct(promotion)
	if err != nil {
		return nil, err
	}

	dp := []domain.Product{}

	promotions := []domain.PromotionProductReply{}
	for _, pro := range product {

		products, err := p.promotionRepo.GetProductById(promotion, pro.ProductID)
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

func (p *promotionUseCase) GetAllPromotion() ([]domain.Promotion, error) {
	pro, err := p.promotionRepo.GetAllPromotion()
	if err != nil {
		return nil, err
	}

	promotions := []domain.Promotion{}
	for _, promotion := range pro {
		promotions = append(promotions, domain.Promotion{
			Model:    promotion.Model,
			Code:     promotion.Code,
			Discount: promotion.Discount,
			Name:     promotion.Name,
			Detail:   promotion.Detail,
		})
	}

	return promotions, nil
}

func (p *promotionUseCase) SearchPromotion(req *domain.Promotion) ([]domain.SearchPromotionReply, error) {
	promotion, err := p.promotionRepo.SearchPromotion(req)
	if err != nil {
		return nil, err
	}

	product := &domain.PromotionProduct{
		PromotionID: promotion.ID,
	}

	res, err := p.promotionRepo.GetPromotionProduct(product)
	if err != nil {
		return nil, err
	}

	dp := []domain.Product{}

	promotions := []domain.SearchPromotionReply{}
	promotions = append(promotions, domain.SearchPromotionReply{
		Code:     promotion.Code,
		Discount: promotion.Discount,
		Name:     promotion.Name,
		Detail:   promotion.Detail,
	})
	for _, pro := range res {

		products, err := p.promotionRepo.GetProductById(product, pro.ProductID)
		if err != nil {
			return nil, err
		}

		promotions = append(promotions, domain.SearchPromotionReply{
			Code:     promotion.Code,
			Discount: promotion.Discount,
			Name:     promotion.Name,
			Detail:   promotion.Detail,
			Product: append(dp, domain.Product{
				Name:   products.Name,
				Detail: products.Detail,
				Price:  products.Price,
			}),
		})
	}

	return promotions, nil

}
