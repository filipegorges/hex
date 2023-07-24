package application

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func NewProductService(persistence ProductPersistenceInterface) *ProductService {
	return &ProductService{Persistence: persistence}
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	return s.Persistence.Get(id)
}

func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price
	if _, err := product.IsValid(); err != nil {
		return nil, err
	}
	return s.Persistence.Save(product)
}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	if err := product.Enable(); err != nil {
		return nil, err
	}
	return s.Persistence.Save(product)
}

func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	if err := product.Disable(); err != nil {
		return nil, err
	}
	return s.Persistence.Save(product)
}
