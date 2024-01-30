package domain

type Product struct {
	Model
	Name   string  `json:"name"`
	Detail string  `json:"detail"`
	Price  float64 `json:"price"`

	BasketProduct    []BasketProduct    `gorm:"foreignKey:ProductID"`
	PromotionProduct []PromotionProduct `gorm:"foreignKey:ProductID"`
}

type ProductForm struct {
	Name   string  `json:"name"`
	Detail string  `json:"detail"`
	Price  float64 `json:"price"`
}

type ProductReply struct {
	Name   string  `json:"name"`
	Detail string  `json:"detail"`
	Price  float64 `json:"price"`
}

func (p *Product) TableName() string {
	return "products"
}

type ProductUseCase interface {
	GetAll() ([]Product, error)
	DeleteProduct(id uint) error
	GetProductById(id uint) (*ProductReply, error)
	EditProduct(p *ProductForm, id uint) (*ProductReply, error)
	AddProduct(p *ProductForm) (*ProductReply, error)
}

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetOneByID(p *Product) (*Product, error)
	Delete(p *Product) error
	Edit(p *Product) (*Product, error)
	Create(p *Product) (*Product, error)
}
