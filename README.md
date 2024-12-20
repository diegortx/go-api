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

### New Features

- **Build for Multiple OS**: The Makefile now includes commands to build the application for Windows, Linux, and macOS.
- **.gitignore**: Added a .gitignore file to exclude unnecessary files from the repository.
- **Docker Support**: Added Docker support to build and run the application in a container.

## Requirements

- Go 1.16+
- MySQL
- Docker

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

## Building the Application

Run the command below to build the application for your OS:

- For Windows:
  ```sh
  make build-windows
  ```
- For Linux:
  ```sh
  make build-linux
  ```
- For macOS:
  ```sh
  make build-macos
  ```

## Using Docker

1. Build the Docker image:
   ```sh
   make docker-build
   ```
2. Run the Docker container:
   ```sh
   docker compose up -d
   ```
