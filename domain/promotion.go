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
	Code      string  `json:"code"`
	Discount  float64 `json:"discount"`
	Name      string  `json:"name"`
	Detail    string  `json:"detail"`
	ProductID uint    `json:"product_id"`
}

type PromotionReply struct {
	Code      string  `json:"code"`
	Discount  float64 `json:"discount"`
	Name      string  `json:"name"`
	Detail    string  `json:"detail"`
	ProductID uint    `json:"product_id"`
}

type PromotionUseCase interface {
	CreatePromotion(p *PromotionForm) error
	EditPromotion(p *PromotionForm) error
	DeletePromotion(id uint) error
	GetPromotionById(id uint) (*PromotionReply, error)
	GetAllPromotion() ([]Promotion, error)
}

type PromotionRepository interface {
	CreatePromotion(p *Promotion) error
	EditPromotion(p *Promotion) error
	DeletePromotion(p *Promotion, id uint) error
	GetPromotionById(p *Promotion, id uint) (*PromotionReply, error)
	GetAllPromotion() ([]Promotion, error)
}
