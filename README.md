# Go API

Este projeto é uma API escrita em Go que interage com um banco de dados para gerenciar produtos.

## Estrutura do Projeto

- `model/`: Contém as definições das estruturas de dados.
- `repository/`: Contém a lógica de acesso ao banco de dados.

## Arquivo `productRepository.go`

Este arquivo contém a implementação do repositório de produtos. Ele inclui métodos para interagir com a tabela de produtos no banco de dados.

### Funções Principais

- `NewProductRepository(db *sql.DB) ProductRepository`: Cria uma nova instância do repositório de produtos.
- `GetProduct() ([]model.Product, error)`: Recupera todos os produtos do banco de dados.

### Exemplo de Uso

```go
package main

import (
	"database/sql"
	"fmt"
	"go-api/repository"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Conectar ao banco de dados
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Criar repositório de produtos
	productRepo := repository.NewProductRepository(db)

	// Obter produtos
	products, err := productRepo.GetProduct()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Exibir produtos
	for _, product := range products {
		fmt.Printf("ID: %d, Nome: %s, Preço: %.2f\n", product.ID, product.Name, product.Price)
	}
}
```

## Requisitos

- Go 1.16+
- MySQL

## Instalação

1. Clone o repositório:
   ```sh
   git clone https://github.com/seu-usuario/go-api.git
   ```
2. Instale as dependências:
   ```sh
   go mod tidy
   ```

## Configuração

Atualize a string de conexão com o banco de dados no exemplo de uso com suas credenciais e informações do banco de dados.

## Executando a Aplicação

Execute o comando abaixo para iniciar a aplicação:

```sh
go run main.go
```
