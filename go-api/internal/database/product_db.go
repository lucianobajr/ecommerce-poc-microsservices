package database

import (
	"database/sql"

	"github.com/lucianobajr/ecommerce-poc-microsservices/go-api/internal/entity"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (pd *ProductDB) GetProducts() ([]*entity.Product, error) {
	rows, err := pd.db.Query("SELECT id,name,description,price,category_id,image_url FROM products")

	if err != nil {
		return nil, err
	}

	// só vai ser executado quando os demais logicas forem executadas
	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var Product entity.Product

		if err := rows.Scan(&Product.ID, &Product.Name, &Product.Description, &Product.Price, &Product.CategoryID, &Product.ImageURL); err != nil {
			return nil, err
		}

		products = append(products, &Product)
	}

	return products, nil
}

func (pd *ProductDB) GetProduct(id string) (*entity.Product, error) {
	var Product entity.Product
	err := pd.db.QueryRow("SELECT id,name,description,price,category_id,image_url FROM products WHERE id = ?", id).Scan(&Product.ID, &Product.Name, &Product.Description, &Product.Price, &Product.CategoryID, &Product.ImageURL)

	if err != nil {
		return nil, err
	}

	return &Product, nil
}

func (pd *ProductDB) GetProductByCategoryID(categoryID string) ([]*entity.Product, error) {
	rows, err := pd.db.Query("SELECT id,name,description,price,category_id,image_url FROM products WHERE category_id = ?", categoryID)

	if err != nil {
		return nil, err
	}

	// só vai ser executado quando os demais logicas forem executadas
	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var Product entity.Product

		if err := rows.Scan(&Product.ID, &Product.Name, &Product.Description, &Product.Price, &Product.CategoryID, &Product.ImageURL); err != nil {
			return nil, err
		}

		products = append(products, &Product)
	}

	return products, nil
}

func (pd *ProductDB) CreateProduct(Product *entity.Product) (*entity.Product, error) {
	_, err := pd.db.Exec("INSERT INTO products (id, name, description, price, category_id, image_url) VALUES (?, ?, ?, ?, ?, ?)", Product.ID, Product.Name, Product.Description, Product.Price, Product.CategoryID, Product.ImageURL)

	if err != nil {
		return nil, err
	}

	return Product, nil
}
