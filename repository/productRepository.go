package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connecton *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return ProductRepository{
		connecton: db,
	}
}

func (pr *ProductRepository) GetProduct() ([]model.Product, error) {
	query := "SELECT id,product_name,price FROM products"
	rows, err := pr.connecton.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err

	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err := rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
		)
		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}
