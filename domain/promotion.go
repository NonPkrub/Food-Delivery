package domain

type Promotion struct {
	Model
	Code     string  `json:"code"`
	Discount float64 `json:"discount"`
	Name     string  `json:"name"`
	Detail   string  `json:"detail"`

	PromotionProduct []PromotionProduct `gorm:"foreignKey:PromotionID"`
}

func (p *Promotion) TableName() string {
	return "promotions"
}

type PromotionForm struct {
	Code     string  `json:"code"`
	Discount float64 `json:"discount"`
	Name     string  `json:"name"`
	Detail   string  `json:"detail"`
}

type PromotionReply struct {
	Code      string  `json:"code"`
	Discount  float64 `json:"discount"`
	Name      string  `json:"name"`
	Detail    string  `json:"detail"`
	ProductID uint    `json:"product_id"`
}

type SearchPromotionReply struct {
	Code     string    `json:"code"`
	Discount float64   `json:"discount"`
	Name     string    `json:"name"`
	Detail   string    `json:"detail"`
	Product  []Product `json:"product"`
}

type PromotionUseCase interface {
	CreatePromotion(p *PromotionForm) error
	EditPromotion(p *PromotionForm, id uint) error
	DeletePromotion(id uint) error
	GetPromotionById(id uint) ([]PromotionProductReply, error)
	GetAllPromotion() ([]Promotion, error)
	SearchPromotion(p *Promotion) ([]SearchPromotionReply, error)
}

type PromotionRepository interface {
	CreatePromotion(p *Promotion) error
	EditPromotion(p *Promotion, id uint) error
	DeletePromotion(p *Promotion, id uint) error
	GetPromotionProduct(*PromotionProduct) ([]PromotionProduct, error)
	SearchPromotion(p *Promotion) (*Promotion, error)
	GetAllPromotion() ([]Promotion, error)
	GetProductById(p *PromotionProduct, id uint) (*PromotionProductReplyId, error)
}
