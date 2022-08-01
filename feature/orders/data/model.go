package data

import (
	"project3/group3/domain"
	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	Qty     		int
	Total		  	int
}

func (o *Orders) ToModel() domain.Orders {
	return domain.Orders{
		ID:        		int(o.ID),
		Qty: 			o.Qty,
		Total:			o.Total,
		CreatedAt: 		o.CreatedAt,
		UpdatedAt: 		o.UpdatedAt,
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

func FromModel(data domain.Orders) Orders {
	var res Orders
	res.Qty = data.Qty
	res.Total = data.Total
	return res
}
