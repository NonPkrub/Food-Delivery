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
	GetAll(string) ([]Product, error)
	DeleteProduct(id uint) error
	GetProductById(id uint) (*ProductReply, error)
	EditProduct(product *ProductForm, id uint) (*ProductReply, error)
	AddProduct(product *ProductForm) (*ProductReply, error)
}

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetOneByID(product *Product) (*Product, error)
	Delete(product *Product) error
	Edit(product *Product) (*Product, error)
	Create(product *Product) (*Product, error)
	GetByQuery(product *Product) (*Product, error)
}
