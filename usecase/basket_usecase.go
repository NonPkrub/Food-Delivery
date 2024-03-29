package usecase

import (
	"Food-delivery/domain"
)

type basketUseCase struct {
	basketRepo           domain.BasketRepository
	basketProductUseCase domain.BasketProductUseCase
}

func NewBasketUseCase(basketRepo domain.BasketRepository, basketProductUseCase domain.BasketProductUseCase) domain.BasketUseCase {
	return &basketUseCase{basketRepo: basketRepo, basketProductUseCase: basketProductUseCase}
}

func (uc *basketUseCase) CreateBasket(form *domain.BasketForm) error {
	basket := &domain.Basket{
		UserID: form.UserID,
	}

	err := uc.basketRepo.CreateOne(basket)
	if err != nil {
		return err
	}

	return nil
}

func (uc *basketUseCase) AddPromotionBasket(form *domain.BasketPromotionForm) error {
	basket := &domain.Basket{
		UserID:      form.UserID,
		PromotionID: &form.PromotionID,
	}

	err := uc.basketRepo.Create(basket)
	if err != nil {
		return err
	}

	return nil
}

func (uc *basketUseCase) DeletePromotionBasket(id uint) error {
	basket := &domain.Basket{
		UserID:      id,
		PromotionID: nil,
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

	baskets := &domain.BasketProduct{
		BasketID: userBasket.ID,
	}

	products, totalPrice, subTotalPrice, discount, err := uc.basketProductUseCase.GetProductInBasket(baskets)
	if err != nil {
		return nil, err
	}

	basketReply := &domain.BasketReply{
		ID:             userBasket.ID,
		UserID:         userBasket.UserID,
		BasketProducts: products,
		TotalPrice:     totalPrice,
		PromotionID:    promotionID,
		SubTotalPrice:  subTotalPrice,
		Discount:       discount,
	}

	return basketReply, nil

}
