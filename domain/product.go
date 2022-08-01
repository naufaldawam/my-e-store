package domain

import (
	// product "project3/group3/feature/product"
	"time"

	// "github.com/labstack/echo/v4"
)

type Product struct {
	ID        		int
	Name      		string
	Size     		string
	PriceProduct  	int
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	// DeletedAt 		time.Time
}

type ProductData interface {
	// Insert(newUser User) (row int, err error)
	// LoginData(authData user.LoginModel) (token, name string, err error)
	// GetSpecific(userID int) (User, error)
	// DeleteData(userID int) (row int, err error)
	// UpdateData(data map[string]interface{}, idUser int) (row int, err error)
}


