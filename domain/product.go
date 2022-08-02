package domain

import "github.com/labstack/echo/v4"

type Product struct {
	ID           int
	ProductName  string
	Stock        int
	Price        int
	ProductImage string
}

//logic
type ProductUseCase interface {
	InsertProduct(newProduct Product) (row int, err error)
	DeleteProduct(productID int) (row int, err error)
}

//query
type ProductData interface {
	InsertProductDB(newProduct Product) (row int, err error)
	DeleteProductDB(productID int) (row int, err error)
}

// handler
type ProductHandler interface {
	InsertProductHandler() echo.HandlerFunc
	DeleteProductHandler() echo.HandlerFunc
}
