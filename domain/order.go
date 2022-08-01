package domain

import "time"

type Orders struct {
	ID        		int
	ID_Product    	int
	Qty     		int
	Total		  	int
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	// DeletedAt 		time.Time
}

type OrdersData interface {
	// Insert(newUser User) (row int, err error)
	// LoginData(authData user.LoginModel) (token, name string, err error)
	// GetSpecific(userID int) (User, error)
	// DeleteData(userID int) (row int, err error)
	// UpdateData(data map[string]interface{}, idUser int) (row int, err error)
}
