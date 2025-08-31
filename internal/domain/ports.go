package domain

type ProductRepository interface {
	FindAll() ([]Products, error)
	FindById(id uint64) (Products, error)
	Save(data *ProductForm) error
}

type ProductService interface {
	FindAll() ([]Products, error)
	FindById(id uint64) (Products, error)
	Create(data *ProductForm) error
}
