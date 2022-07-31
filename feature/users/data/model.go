package data

import (
	"project3/group3/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Phone    string
	Role     string `gorm:"default:user"`
	Email    string `gorm:"unique" validate:"required,email"`
	Password string
}

func (u *User) ToModel() domain.User {
	return domain.User{
		ID:        int(u.ID),
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
func ParseToArr(arr []User) []domain.User {
	var res []domain.User
	for _, val := range arr {
		res = append(res, val.ToModel())
	}

	return res
}

func FromModel(data domain.User) User {
	var res User
	res.Email = data.Email
	res.Name = data.Name
	res.Password = data.Password
	return res
}
