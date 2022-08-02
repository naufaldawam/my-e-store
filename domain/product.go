package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Product struct {
	ID           int
	ProductName  string
	Stock        int
	Price        int
	ProductImage string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UserID       int
	User         UserAdmin
}

type UserAdmin struct {
	ID   int
	Name string
	Role string
}

//logic
type ProductUseCase interface {
	InsertProduct(newProduct Product) (result Product, err error)
	DeleteProduct(productID int) (row int, err error)
	GetAllData(limit, offset int) (data []Product, err error)
	GetProductById(id int) (data Product, err error)
}

//query
type ProductData interface {
	GetUser(idProduct int) (result Product, err error) //buat ambil nama id
	InsertProductDB(newProduct Product) (result Product, err error)
	DeleteProductDB(productID int) (row int, err error)
	SelectData(limit, offset int) (data []Product, err error)
	SelectDataById(id int) (data Product, err error)
}

// handler
type ProductHandler interface {
	InsertProductHandler() echo.HandlerFunc
	DeleteProductHandler() echo.HandlerFunc
}
