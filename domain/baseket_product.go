package domain

type BasketProduct struct {
	BasketID  uint `json:"basket_id"`
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

type BasketProductReply struct {
	BasketID uint          `json:"basket_id"`
	Product  []ProductForm `json:"product"`
	Quantity uint          `json:"quantity"`
}

type BasketProductPrice struct {
	BasketID uint    `json:"basket_id"`
	Name     string  `json:"name"`
	Detail   string  `json:"detail"`
	Price    float64 `json:"price"`
	Quantity uint    `json:"quantity"`
}

type BasketPromotionReply struct {
	PromotionID uint    `json:"promotion_id"`
	Code        string  `json:"code"`
	Discount    float64 `json:"discount"`
	Name        string  `json:"name"`
	ProductID   uint    `json:"product_id"`
}

func (p *BasketProduct) TableName() string {
	return "basket_products"
}

type BasketProductUseCase interface {
	AddProductInBasket(b *BasketProduct) error
	EditProductInBasket(b *BasketProduct) error
	DeleteProductInBasket(b *BasketProduct) error
	GetProductInBasket(b *BasketProduct) ([]BasketProductReply, float64, error)
}

type BasketProductRepository interface {
	Create(b *BasketProduct) error
	Edit(b *BasketProduct) error
	Delete(b *BasketProduct) error
	FindAllByID(b *BasketProduct) ([]BasketProduct, error)
	GetOneById(b *BasketProduct) (*BasketProduct, error)
	FindOne(b *BasketProduct) (*BasketProduct, error)
	GetPromotionByBasketID(b *BasketProduct) (uint, error)
}
