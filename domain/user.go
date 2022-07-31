package domain

import (
	user "project3/group3/feature/users"
	"time"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//logic
type UserUseCase interface {
	AddUser(newUser User) (row int, err error)
	Login(auth user.LoginModel) (token, name string, err error)
	GetProfile(id int) (User, error)
	DeleteCase(userID int) (row int, err error)
	UpdateCase(input User, idUser int) (row int, err error)
}

//query
type UserData interface {
	Insert(newUser User) (row int, err error)
	LoginData(authData user.LoginModel) (token, name string, err error)
	GetSpecific(userID int) (User, error)
	DeleteData(userID int) (row int, err error)
	UpdateData(data map[string]interface{}, idUser int) (row int, err error)
}
