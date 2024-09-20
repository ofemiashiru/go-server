package product

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	/*
		ABOVE
		importing with underscore means we only want the
		package for it's side effects not any of its exported identifiers
	*/
	"acme/model"
)

type PostgresProductRepository struct {
	DB *sqlx.DB
}

// New constructor
func NewPostgresProductRepository(db *sqlx.DB) *PostgresProductRepository {
	return &PostgresProductRepository{DB: db}
}

func (repo *PostgresProductRepository) GetProducts() ([]model.Product, error) {

	products := []model.Product{}

	err := sqlx.Select(repo.DB, &products, "SELECT * FROM products;")
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return []model.Product{}, errors.New("database could not be queried")
	}

	return products, nil
}

func (repo *PostgresProductRepository) AddProduct(product []model.Product) (id int, err error) {
	err = sqlx.Select(repo.DB, &product, "INSERT INTO products (name, price, stock_count) VALUES ($1, $2, $3) RETURNING id;", product[0].Name, product[0].Price, product[0].StockCount)
	if err != nil {
		fmt.Println("Error inserting product into the database:", err)
		return 0, errors.New("could not insert product")
	}

	return product[0].ID, nil
}

func (repo *PostgresProductRepository) Close() {

}
