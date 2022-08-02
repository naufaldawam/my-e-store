package delivery

import "project3/group3/domain"

type Product struct {
	ID           int    `json:"id"`
	ProductName  string `json:"product_name"`
	ProductImage string `json:"product_image"`
	Stock        int    `json:"stock"`
	Price        int    `json:"price"`
	Admin        User   `json:"user"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func FromModel(data domain.Product) Product {
	return Product{
		ID:           data.ID,
		ProductName:  data.ProductImage,
		ProductImage: data.ProductImage,
		Stock:        data.Stock,
		Price:        data.Price,
		Admin: User{
			ID:   data.User.ID,
			Name: data.User.Name,
			Role: data.User.Role,
		},
	}
}

func FromModelList(data []domain.Product) []Product {
	result := []Product{}
	for key := range data {
		result = append(result, FromModel(data[key]))
	}
	return result
}
