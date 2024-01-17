package domain

type BasketProduct struct {
	BasketID  uint `json:"basket_id"`
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

type BasketProductReply struct {
	BasketID uint      `json:"basket_id"`
	Product  []Product `json:"product"`
	Quantity uint      `json:"quantity"`
}

type BasketProductPrice struct {
	BasketID uint    `json:"basket_id"`
	Name     string  `json:"name"`
	Detail   string  `json:"detail"`
	Price    float64 `json:"price"`
	Quantity uint    `json:"quantity"`
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
	AddProductInBasket(b *BasketProduct) error
	EditProductInBasket(b *BasketProduct) error
	DeleteProductInBasket(b *BasketProduct) error
	GetProductInBasket(b *BasketProduct) ([]BasketProduct, error)
	GetProductById(b *BasketProduct, id uint) (*BasketProductPrice, error)
}
