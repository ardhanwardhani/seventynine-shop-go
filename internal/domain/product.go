package domain

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Stock int    `jeson:"stock"`
}

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetByID(id int) (*Product, error)
	Create(product *Product) error
	Update(product *Product) error
	Delete(id int) error
}
