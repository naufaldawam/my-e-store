package data

import (
	"project3/group3/domain"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName  string
	ProductImage string
	Stock        int
	Price        int
}

func (p *Product) ToModel() domain.Product {
	return domain.Product{
		ID:           int(p.ID),
		ProductName:  p.ProductName,
		Stock:        p.Stock,
		Price:        p.Price,
		ProductImage: p.ProductImage,
	}
}

func ParseToArr(arr []Product) []domain.Product {
	var res []domain.Product
	for _, val := range arr {
		res = append(res, val.ToModel())
	}
	return res
}

func FromModel(data domain.Product) Product {
	var res Product
	res.ProductName = data.ProductName
	res.Stock = data.Stock
	res.Price = data.Price
	res.ProductImage = data.ProductImage
	return res
}
