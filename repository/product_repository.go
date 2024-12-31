package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return ProductRepository{
		connection: db,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id,product_name,price FROM products"
	rows, err := pr.connection.Query(query)

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

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int
	query, err := pr.connection.Prepare("INSERT INTO products(product_name,price) VALUES($1,$2) returning id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	query.Close()
	return id, nil
}

func (pr *ProductRepository) GetProductById(id_product int) (*model.Product, error) {

	query, err := pr.connection.Prepare("SELECT id,product_name,price FROM products WHERE id=$1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var product model.Product

	err = query.QueryRow(id_product).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	query.Close()
	return &product, nil
}

func (pr *ProductRepository) UpdateProduct(id_product int, product model.Product) (*model.Product, error) {
	query, err := pr.connection.Prepare("UPDATE products SET product_name=$1, price=$2 WHERE id=$3")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_, err = query.Exec(product.Name, product.Price, id_product)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	query.Close()

	updatedProduct, err := pr.GetProductById(id_product)
	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func (pr *ProductRepository) DeleteProductById(id_product int) (*model.Product, error) {

	product, err := pr.GetProductById(id_product)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, fmt.Errorf("product not found")
	}

	query, err := pr.connection.Prepare("DELETE FROM products WHERE id=$1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_, err = query.Exec(id_product)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	query.Close()

	return product, nil
}
