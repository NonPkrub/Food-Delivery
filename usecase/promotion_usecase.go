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
	for i := range product {
		promotion := &domain.PromotionProduct{
			PromotionID: id,
			ProductID:   product[i].ProductID,
		}

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

		dp = append(dp, domain.Product{
			Name:   productDetail.Name,
			Detail: productDetail.Detail,
			Price:  productDetail.Price,
		})
	}

	promotionID := &domain.Promotion{}
	promotionID.ID = promotion.PromotionID
	promotionDetail, err := uc.promotionRepo.FindOne(promotionID)
	if err != nil {
		return nil, err
	}

	promotions = append(promotions, domain.PromotionProductReply{
		Code:     promotionDetail.Code,
		Discount: promotionDetail.Discount,
		Name:     promotionDetail.Name,
		Detail:   promotionDetail.Detail,
		Product:  dp,
	})

	return promotions, nil
}

func (uc *promotionUseCase) GetAllPromotion(queryCode, queryName string) ([]domain.PromotionDetail, error) {
	if queryCode == "" && queryName == "" {
		promotion, err := uc.promotionRepo.GetAll()
		if err != nil {
			return nil, err
		}

		dp := []domain.Product{}
		promotions := []domain.PromotionDetail{}

		for i := range promotion {
			promotionID := &domain.PromotionProduct{
				PromotionID: promotion[i].ID,
			}

			product, err := uc.promotionRepo.GetAllByID(promotionID)
			if err != nil {
				return nil, err
			}

			for range product {
				products, err := uc.promotionRepo.FindOneByID(&promotion[i])
				if err != nil {
					return nil, err
				}

				productId := &domain.Product{}
				productId.ID = products.ProductID
				productDetail, err := uc.productRepo.GetOneByID(productId)
				if err != nil {
					return nil, err
				}

				dp = append(dp, domain.Product{
					Name:   productDetail.Name,
					Detail: productDetail.Detail,
					Price:  productDetail.Price,
				})
			}

			promotions = append(promotions, domain.PromotionDetail{
				Code:     promotion[i].Code,
				Discount: promotion[i].Discount,
				Name:     promotion[i].Name,
				Detail:   promotion[i].Detail,
				Product:  dp,
			})
		}
		return promotions, nil
	}

	form := &domain.Promotion{Code: queryCode, Name: queryName}
	promotion, err := uc.promotionRepo.GetByQuery(form)
	if err != nil {
		return nil, err
	}

	dp := []domain.Product{}
	promotions := []domain.PromotionDetail{}

	promotionID := &domain.PromotionProduct{
		PromotionID: promotion.ID,
	}

	product, err := uc.promotionRepo.GetAllByID(promotionID)
	if err != nil {
		return nil, err
	}

	for range product {
		products, err := uc.promotionRepo.FindOneByID(promotion)
		if err != nil {
			return nil, err
		}

		productId := &domain.Product{}
		productId.ID = products.ProductID
		productDetail, err := uc.productRepo.GetOneByID(productId)
		if err != nil {
			return nil, err
		}

		dp = append(dp, domain.Product{
			Name:   productDetail.Name,
			Detail: productDetail.Detail,
			Price:  productDetail.Price,
		})
	}

	promotions = append(promotions, domain.PromotionDetail{
		Code:     promotion.Code,
		Discount: promotion.Discount,
		Name:     promotion.Name,
		Detail:   promotion.Detail,
		Product:  dp,
	})

	return promotions, nil
}
