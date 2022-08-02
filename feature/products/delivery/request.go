package delivery

import "project3/group3/domain"

type InsertProductFormat struct {
	ProductName  string `json:"product_name" form:"product_name"`
	Stock        int    `json:"stock" form:"stock"`
	Price        int    `json:"price" form:"price"`
	ProductImage string `form:"product_image"`
}

func (i *InsertProductFormat) ToModel() domain.Product {
	return domain.Product{
		ProductName:  i.ProductName,
		Stock:        i.Stock,
		Price:        i.Price,
		ProductImage: i.ProductImage,
	}
}
