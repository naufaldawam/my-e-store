package data

import (
	"project3/group3/domain"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID        		int
	Name      		string
	Size     		string
	PriceProduct  	int
}

func (p *Product) ToModel() domain.Product {
	return domain.Product{
		ID:        		int(p.ID),
		Name:     		p.Name,
		Size:     		p.Size,
		PriceProduct:   p.PriceProduct,
		CreatedAt: 		p.CreatedAt,
		UpdatedAt: 		p.UpdatedAt,
		// DeletedAt:  	p.DeletedAt,
	}
}
// func ParseToArr(arr []Product) []domain.Product {
// 	var res []domain.Product
// 	for _, val := range arr {
// 		res = append(res, val.ToModel())
// 	}

// 	return res
// }

func FromModel(data domain.Product) Product {
	var res Product
	res.Name = data.Name
	res.Size = data.Size
	res.PriceProduct = data.PriceProduct
	return res
}
