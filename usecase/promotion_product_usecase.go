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

// func (uc *promotionProductUseCase) GetPromotionProduct(form *domain.PromotionProductForm) ([]domain.PromotionProductReply, error) {
// 	promotion := &domain.PromotionProduct{
// 		PromotionID: form.PromotionID,
// 		ProductID:   form.ProductID,
// 	}

// 	product, err := uc.promotionProductRepo.FindAllByID(promotion)
// 	if err != nil {
// 		return nil, err
// 	}

// 	dp := []domain.Product{}

// 	promotions := []domain.PromotionProductReply{}
// 	for _, pro := range product {
// 		products, err := uc.promotionProductRepo.GetOneByID(promotion)
// 		if err != nil {
// 			return nil, err
// 		}

// 		fmt.Println(products, products.ProductID)

// 		productId := &domain.Product{}
// 		productId.ID = products.ProductID

// 		productDetail, err := uc.productRepo.GetOneByID(productId)
// 		fmt.Println(productDetail, productId)
// 		if err != nil {
// 			return nil, err
// 		}

// 		promotions = append(promotions, domain.PromotionProductReply{
// 			PromotionID: pro.PromotionID,
// 			Product: append(dp, domain.Product{
// 				Name:   productDetail.Name,
// 				Detail: productDetail.Detail,
// 				Price:  productDetail.Price,
// 			}),
// 		})
// 	}

// 	return promotions, nil
// }
