# Go API

This project is an API written in Go that interacts with a database to manage products.

## Project Structure

- `model/`: Contains data structure definitions.
- `repository/`: Contains database access logic.

## File `productRepository.go`

This file contains the implementation of the product repository. It includes methods to interact with the products table in the database.

### Main Functions

- `NewProductRepository(db *sql.DB) ProductRepository`: Creates a new instance of the product repository.
- `GetProduct() ([]model.Product, error)`: Retrieves all products from the database.

### Usage Example

```go
package main

import (
	"database/sql"
	"fmt"
	"go-api/repository"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to the database
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Create product repository
	productRepo := repository.NewProductRepository(db)

	// Get products
	products, err := productRepo.GetProduct()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display products
	for _, product := range products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", product.ID, product.Name, product.Price)
	}
}
```

## Requirements

- Go 1.16+
- MySQL

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/go-api.git
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```

## Configuration

Update the database connection string in the usage example with your credentials and database information.

## Running the Application

Run the command below to start the application:

```sh
go run main.go
```
