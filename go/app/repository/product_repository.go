package repository

import (
	"fmt"

	"github.com/yakob-abada/go-api/go/app/entity"
)

type ProductRepository struct {
	dBConnection DatabaseConnection
}

func NewProductRepository(dbConnection DatabaseConnection) *ProductRepository {
	return &ProductRepository{
		dBConnection: dbConnection,
	}
}

func (pr *ProductRepository) FindById(sku string) (*entity.Product, error) {
	var product entity.Product

	db, err := pr.dBConnection.Connect()

	if err != nil {
		return nil, err
	}

	row := db.QueryRow("SELECT sku, name, price, product_type, size FROM product where sku = ?", sku)

	if err := row.Scan(&product.Sku, &product.Name, &product.Price, &product.ProductType, &product.Size); err != nil {
		return nil, fmt.Errorf("product with id %s: %v", sku, err)
	}

	return &product, nil
}

func (pr *ProductRepository) FindAll() (*[]entity.Product, error) {
	var products []entity.Product = []entity.Product{}

	db, err := pr.dBConnection.Connect()

	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT sku FROM product")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.Sku); err != nil {
			return nil, fmt.Errorf("product: %v", err)
		}

		products = append(products, product)
	}

	return &products, nil
}
