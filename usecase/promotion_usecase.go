package usecase

import (
	"Food-delivery/domain"
)

type promotionUseCase struct {
	promotionRepo domain.PromotionRepository
	productRepo   domain.ProductRepository
}

func NewPromotionUseCase(promotionRepo domain.PromotionRepository, productRepo domain.ProductRepository) domain.PromotionUseCase {
	return &promotionUseCase{promotionRepo: promotionRepo, productRepo: productRepo}
}

func (uc *promotionUseCase) CreatePromotion(form *domain.PromotionForm) error {
	promotion := &domain.Promotion{
		Code:     form.Code,
		Discount: form.Discount,
		Name:     form.Name,
		Detail:   form.Detail,
	}

	err := uc.promotionRepo.Create(promotion)
	if err != nil {
		return err
	}

	return nil
}

func (uc *promotionUseCase) EditPromotion(form *domain.PromotionForm, id uint) error {
	promotion := &domain.Promotion{
		Code:     form.Code,
		Discount: form.Discount,
		Name:     form.Name,
		Detail:   form.Detail,
	}
	promotion.ID = id
	err := uc.promotionRepo.Edit(promotion)
	if err != nil {
		return err
	}

	return nil
}

func (uc *promotionUseCase) DeletePromotion(id uint) error {
	var promotion domain.Promotion
	promotion.ID = id
	err := uc.promotionRepo.Delete(&promotion)
	if err != nil {
		return err
	}

	return nil
}

func (uc *promotionUseCase) GetPromotionById(id uint) ([]domain.PromotionProductReply, error) {
	promotion := &domain.PromotionProduct{
		PromotionID: id,
	}

	product, err := uc.promotionRepo.GetAllByID(promotion)
	if err != nil {
		return nil, err
	}

	dp := []domain.Product{}

	promotions := []domain.PromotionProductReply{}
	for _, pro := range product {
		products, err := uc.promotionRepo.GetOneByID(promotion)
		if err != nil {
			return nil, err
		}

		productId := &domain.Product{}
		productId.ID = products.ProductID

		productDetail, err := uc.productRepo.GetOneByID(productId)
		if err != nil {
			return nil, err
		}

		promotions = append(promotions, domain.PromotionProductReply{
			PromotionID: pro.PromotionID,
			Product: append(dp, domain.Product{
				Name:   productDetail.Name,
				Detail: productDetail.Detail,
				Price:  productDetail.Price,
			}),
		})
	}

	return promotions, nil
}

func (uc *promotionUseCase) GetAllPromotion(queryCode, queryName string) ([]domain.Promotion, error) {
	if queryCode == "" && queryName == "" {
		return uc.promotionRepo.GetAll()
	}

	form := &domain.Promotion{Code: queryCode, Name: queryName}
	promotion, err := uc.promotionRepo.GetByQuery(form)
	if err != nil {
		return nil, err
	}

	return []domain.Promotion{*promotion}, nil
}

// func (uc *promotionUseCase) SearchPromotion(req *domain.Promotion) ([]domain.SearchPromotionReply, error) {
// 	promotion, err := uc.promotionRepo.SearchPromotion(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	product := &domain.PromotionProduct{
// 		PromotionID: promotion.ID,
// 	}

// 	res, err := uc.promotionRepo.GetAllByID(product)
// 	if err != nil {
// 		return nil, err
// 	}

// 	dp := []domain.Product{}

// 	promotions := []domain.SearchPromotionReply{}
// 	promotions = append(promotions, domain.SearchPromotionReply{
// 		Code:     promotion.Code,
// 		Discount: promotion.Discount,
// 		Name:     promotion.Name,
// 		Detail:   promotion.Detail,
// 	})
// 	for _, pro := range res {

// 		products, err := uc.promotionRepo.GetOneByID(product)
// 		if err != nil {
// 			return nil, err
// 		}

// 		promotions = append(promotions, domain.SearchPromotionReply{
// 			Code:     promotion.Code,
// 			Discount: promotion.Discount,
// 			Name:     promotion.Name,
// 			Detail:   promotion.Detail,
// 			Product: append(dp, domain.Product{
// 				Name:   products.Name,
// 				Detail: products.Detail,
// 				Price:  products.Price,
// 			}),
// 		})
// 	}

// 	return promotions, nil

// }
