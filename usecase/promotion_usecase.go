package usecase

import "Food-delivery/domain"

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

func (p *promotionUseCase) GetPromotionById(id uint) (*domain.PromotionReply, error) {
	var promotion domain.Promotion
	promotions, err := p.promotionRepo.GetPromotionById(&promotion, id)
	if err != nil {
		return nil, err
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
