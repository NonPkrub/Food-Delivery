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

type PromotionDetail struct {
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
	GetAllPromotion(string, string) ([]PromotionDetail, error)
}

type PromotionRepository interface {
	Create(promotion *Promotion) error
	Edit(promotion *Promotion) error
	Delete(promotion *Promotion) error
	GetAllByID(*PromotionProduct) ([]PromotionProduct, error)
	GetAll() ([]Promotion, error)
	GetOneByID(promotion *PromotionProduct) (*PromotionProduct, error)
	FindOneByID(promotion *Promotion) (*PromotionProduct, error)
	FindOne(promotion *Promotion) (*Promotion, error)
	GetByQuery(promotion *Promotion) (*Promotion, error)
}
