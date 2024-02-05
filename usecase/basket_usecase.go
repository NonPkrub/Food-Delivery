package usecase

import (
	"Food-delivery/domain"
	"fmt"
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
		PromotionID: form.PromotionID,
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
		PromotionID: 0,
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
	if userBasket.PromotionID != 0 {
		promotionID = userBasket.PromotionID
	}

	baskets := &domain.BasketProduct{
		BasketID: userBasket.ID,
	}

	products, totalPrice, err := uc.basketProductUseCase.GetProductInBasket(baskets)
	fmt.Println(products, totalPrice)
	if err != nil {
		return nil, err
	}

	basketReply := &domain.BasketReply{
		ID:            userBasket.ID,
		UserID:        userBasket.UserID,
		BasketProduct: products,
		TotalPrice:    totalPrice,
		PromotionID:   promotionID,
	}

	return basketReply, nil

}
