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
	PromotionID uint      `json:"promotion_id"`
	Product     []Product `json:"product"`
}

type PromotionProductReplyId struct {
	PromotionID uint    `json:"promotion_id"`
	Name        string  `json:"name"`
	Detail      string  `json:"detail"`
	Price       float64 `json:"price"`
}

func (p *PromotionProduct) TableName() string {
	return "promotions_products"
}

type PromotionProductUseCase interface {
	AddPromotionProduct(*PromotionProductForm) error
	EditPromotionProduct(*PromotionProductForm) error
	GetPromotionProduct(*PromotionProductForm) ([]PromotionProductReply, error)
}

type PromotionProductRepository interface {
	Create(*PromotionProduct) error
	Edit(*PromotionProduct) error
	FindAllByID(*PromotionProduct) ([]PromotionProduct, error)
	GetOneByID(p *PromotionProduct) (*PromotionProduct, error)
}
