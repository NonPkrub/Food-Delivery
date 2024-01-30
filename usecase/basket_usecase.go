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

// func (b *basketUseCase) CreateBasket(req *domain.BasketForm) error {

// 	basket := &domain.Basket{
// 		UserID:      req.UserID,
// 		PromotionID: nil,
// 	}

// 	err := b.basketRepo.CreateBasket(basket)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (uc *basketUseCase) AddPromotionBasket(req *domain.BasketPromotionForm) error {
	basket := &domain.Basket{
		UserID:      req.UserID,
		PromotionID: &req.PromotionID,
	}

	err := uc.basketRepo.Create(basket)
	if err != nil {
		return err
	}

	return nil
}

func (uc *basketUseCase) DeletePromotionBasket(id uint) error {
	basket := &domain.Basket{
		UserID: id,
	}

	err := uc.basketRepo.Delete(basket)
	if err != nil {
		return err
	}

	return nil
}

func (uc *basketUseCase) GetBasketByUserId(id uint) (*domain.BasketReply, error) {
	basket := &domain.Basket{
		UserID: id,
	}

	userBasket, err := uc.basketRepo.GetOneByID(basket)
	if err != nil {
		return nil, err
	}

	var promotionID uint
	if userBasket.PromotionID != nil {
		promotionID = *userBasket.PromotionID
	}

	basketReply := &domain.BasketReply{
		UserID:      userBasket.UserID,
		PromotionID: promotionID,
	}

	return basketReply, nil

}
