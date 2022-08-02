package data

import (
	"project3/group3/domain"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName  string `json:"product_name" form:"product_name"`
	ProductImage string `form:"product_image"`
	Stock        int    `json:"stock" form:"stock"`
	Price        int    `json:"price" form:"price"`
	UserID       int
	User         User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

type User struct {
	gorm.Model
	Name     string
	Phone    string
	Role     string
	Email    string `gorm:"unique" validate:"required,email"`
	Password string
	Product  []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (p *Product) ToDomain() domain.Product {
	return domain.Product{
		ID:           int(p.ID),
		ProductName:  p.ProductName,
		ProductImage: p.ProductImage,
		Stock:        p.Stock,
		Price:        p.Price,
		User: domain.UserAdmin{
			ID:   int(p.UserID),
			Name: p.User.Name,
			Role: p.User.Role,
		},
	}
}

func ParseToArrProduct(arr []Product) []domain.Product {
	var res []domain.Product
	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func FromDomain(data domain.Product) Product {
	var res Product
	res.ProductName = data.ProductName
	res.ProductImage = data.ProductImage
	res.Stock = data.Stock
	res.Price = data.Price
	res.UserID = data.User.ID
	return res
}
