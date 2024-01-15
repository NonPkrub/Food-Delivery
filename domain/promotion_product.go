package domain

type PromotionProduct struct {
	PromotionID uint `json:"promotion_id"`
	ProductID   uint `json:"product_id"`
}

type PromotionProductForm struct {
	PromotionID uint `json:"promotion_id"`
	ProductID   uint `json:"product_id"`
}

type PromotionProductReply struct {
	PromotionID uint `json:"promotion_id"`
	ProductID   uint `json:"product_id"`
}

func (p *PromotionProduct) TableName() string {
	return "promotions_products"
}

type PromotionProductUseCase interface {
	AddPromotionProduct(*PromotionProductForm) error
	EditPromotionProduct(*PromotionProductForm) error
	GetPromotionProduct(*PromotionProductForm) (*PromotionProductReply, error)
}

type PromotionProductRepository interface {
	AddPromotionProduct(*PromotionProduct) error
	EditPromotionProduct(*PromotionProduct) error
	GetPromotionProduct(*PromotionProduct) (*PromotionProductReply, error)
}
