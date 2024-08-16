package postgres

import (
	"database/sql"
	"seventynine-shop-go/internal/domain"
)

type PostgresProductRepository struct {
	db *sql.DB
}

func NewPostgresProductRepository(db *sql.DB) domain.ProductRepository {
	return &PostgresProductRepository{db: db}
}

func (r *PostgresProductRepository) GetAll() ([]domain.Product, error) {
	rows, err := r.db.Query("SELECT id, name, stock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Stock); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *PostgresProductRepository) GetByID(id int) (*domain.Product, error) {
	var product domain.Product
	err := r.db.QueryRow("SELECT id, name, stock FROM products WHERE id = $1", id).Scan(&product.ID, &product.Name, &product.Stock)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *PostgresProductRepository) Create(product *domain.Product) error {
	_, err := r.db.Exec("INSERT INTO products (name, stock) VALUES ($1, $2)", product.Name, product.Stock)
	return err
}

func (r *PostgresProductRepository) Update(product *domain.Product) error {
	_, err := r.db.Exec("UPDATE products SET name = $1, stock = $2 WHERE id = $3", product.Name, product.Stock, product.ID)
	return err
}

func (r *PostgresProductRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id = $1", id)
	return err
}
