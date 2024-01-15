package usecase

import (
	"Food-delivery/domain"
)

type basketUseCase struct {
	basketRepo domain.BasketRepository
}

func NewBasketUseCase(basketRepo domain.BasketRepository) domain.BasketUseCase {
	return &basketUseCase{basketRepo: basketRepo}
}

func (b *basketUseCase) CreateBasket(req *domain.BasketForm) error {

	basket := &domain.Basket{
		UserID: req.UserID,
	}

	err := b.basketRepo.CreateBasket(basket)
	if err != nil {
		return err
	}

	return nil
}

func (b *basketUseCase) AddPromotionBasket(req *domain.BasketPromotionForm) error {

	basket := &domain.Basket{
		UserID:      req.UserID,
		PromotionID: req.PromotionID,
	}

	err := b.basketRepo.AddPromotionBasket(basket)
	if err != nil {
		return err
	}

	return nil
}

func (b *basketUseCase) DeletePromotionBasket(req *domain.BasketPromotionForm) error {

	basket := &domain.Basket{
		UserID:      req.UserID,
		PromotionID: req.PromotionID,
	}

	err := b.basketRepo.DeletePromotionBasket(basket)
	if err != nil {
		return err
	}

	return nil
}

func (b *basketUseCase) GetBasketByUserId(id uint) (*domain.BasketReply, error) {

	basket := &domain.Basket{
		UserID: id,
	}

	userBasket, err := b.basketRepo.GetBasketByUserId(basket)
	if err != nil {
		return nil, err
	}

	return userBasket, nil

}
