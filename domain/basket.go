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
	Model
	UserID      uint `json:"user_id"`
	PromotionID uint `json:"promotion_id"`
}

type BasketPromotionForm struct {
	UserID      uint `json:"user_id"`
	PromotionID uint `json:"promotion_id"`
}

func (p *Basket) TableName() string {
	return "baskets"
}

type BasketUseCase interface {
	CreateBasket(b *BasketForm) error
	AddPromotionBasket(b *BasketPromotionForm) error
	DeletePromotionBasket(uint) error
	GetBasketByUserId(uint) (*BasketReply, error)
}

type BasketRepository interface {
	CreateBasket(b *Basket) error
	AddPromotionBasket(b *Basket) error
	DeletePromotionBasket(b *Basket) error
	GetBasketByUserId(b *Basket) (*BasketReply, error)
}
