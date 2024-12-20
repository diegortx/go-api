package controller

import (
	"go-api/model"
	"go-api/requests"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Message: "Failed to get products",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Message: "Success",
		Data:    products,
	})
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var req requests.CreateProductRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "Failed to get payload on request",
			Data:    err.Error(),
		})
		return
	}

	product := model.Product{
		Name:  req.Name,
		Price: req.Price,
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "Failed to insert product",
			Data:    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, model.Response{
		Message: "Product created successfully",
		Data:    insertedProduct,
	})

}

func (p productController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("productId")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "id is required",
		})
		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "id must be a number",
		})
		return
	}

	products, err := p.productUseCase.GetProductById(productId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Message: "Failed to get product",
			Data:    err.Error(),
		})
		return
	}

	if products == nil {
		ctx.JSON(http.StatusNotFound, model.Response{
			Message: "Product not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Message: "Success",
		Data:    products,
	})
}

func (p *productController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("productId")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "id is required",
		})
		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "id must be a number",
		})
		return
	}

	var product model.Product

	err = ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "Failed to get payload",
			Data:    err.Error(),
		})
		return
	}

	updatedProduct, err := p.productUseCase.UpdateProduct(productId, product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "Failed to update product",
			Data:    err.Error(),
		})
		return
	}

	if updatedProduct == nil {
		ctx.JSON(http.StatusNotFound, model.Response{
			Message: "Product not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Message: "Product updated successfully",
		Data:    updatedProduct,
	})
}

func (p *productController) DeleteProductById(ctx *gin.Context) {
	id := ctx.Param("productId")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "id is required",
		})
		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "id must be a number",
		})
		return
	}

	product, err := p.productUseCase.DeleteProductById(productId)

	if err != nil {
		if len(err.Error()) > 0 {
			ctx.JSON(http.StatusNotFound, model.Response{
				Message: err.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, model.Response{
		Message: "Product deleted successfully",
		Data:    product,
	})
}
