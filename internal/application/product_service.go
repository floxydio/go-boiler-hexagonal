package application

import "github.com/floxydio/boiler-fiber/internal/domain"

type productService struct {
	repo domain.ProductRepository
}

func NewProductService(r domain.ProductRepository) domain.ProductService {
	return &productService{repo: r}
}

func (s *productService) Create(data *domain.ProductForm) error {
	return s.repo.Save(data)
}

func (s *productService) FindAll() ([]domain.Products, error) {
	return s.repo.FindAll()
}

func (s *productService) FindById(id uint64) (domain.Products, error) {
	return s.repo.FindById(id)
}
