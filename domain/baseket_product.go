package domain

type BasketProduct struct {
	BasketID  uint `json:"basket_id"`
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

type BasketProductForm struct {
	BasketID  uint `json:"basket_id"`
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

type BasketProductReply struct {
	BasketID  uint `json:"basket_id"`
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

func (p *BasketProduct) TableName() string {
	return "basket_products"
}

type BasketProductUseCase interface {
	AddProductInBasket(b *BasketProductForm) error
	EditProductInBasket(b *BasketProductForm) error
	DeleteProductInBasket(b *BasketProductForm) error
	GetProductInBasket(b *BasketProductForm) (*BasketProductReply, error)
}

type BasketProductRepository interface {
	AddProductInBasket(b *BasketProduct) error
	EditProductInBasket(b *BasketProduct) error
	DeleteProductInBasket(b *BasketProduct) error
	GetProductInBasket(b *BasketProduct) (*BasketProductReply, error)
}
