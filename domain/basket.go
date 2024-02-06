package domain

type Basket struct {
	Model
	UserID      uint      `json:"user_id"`
	User        User      `gorm:"foreignKey:UserID;references:ID"`
	PromotionID *uint     `json:"promotion_id" gorm:"nullable"`
	Promotion   Promotion `gorm:"foreignKey:PromotionID;reference:ID"`

	BasketProduct []BasketProduct `gorm:"foreignKey:BasketID"`
}

type BasketForm struct {
	UserID      uint `json:"user_id"`
	PromotionID uint `json:"promotion_id"`
}

type BasketReply struct {
	ID             uint                 `json:"id"`
	UserID         uint                 `json:"user_id"`
	BasketProducts []BasketProductReply `json:"basket_products"`
	PromotionID    uint                 `json:"promotion_id"`
	TotalPrice     float64              `json:"total_price"`
	SubTotalPrice  float64              `json:"subtotal_price"`
	Discount       float64              `json:"discount"`
}

type BasketPromotionForm struct {
	UserID      uint `json:"user_id"`
	PromotionID uint `json:"promotion_id"`
}

func (p *Basket) TableName() string {
	return "baskets"
}

type BasketUseCase interface {
	CreateBasket(basket *BasketForm) error
	AddPromotionBasket(basket *BasketPromotionForm) error
	DeletePromotionBasket(id uint) error
	GetBasketByUserId(id uint) (*BasketReply, error)
}

type BasketRepository interface {
	CreateOne(basket *Basket) error
	Create(basket *Basket) error
	Delete(basket *Basket) error
	GetOneByID(basket *Basket) (*Basket, error)
}
