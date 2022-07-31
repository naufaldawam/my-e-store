package delivery

import "project3/group3/domain"

type InsertFormat struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email" gorm:"unique" validate:"required,email"`
	Role     string `gorm:"default:user"`
	Password string `json:"password" form:"password"`
}

func (i *InsertFormat) ToModel() domain.User {
	return domain.User{
		Name:     i.Name,
		Email:    i.Email,
		Password: i.Password,
	}
}
