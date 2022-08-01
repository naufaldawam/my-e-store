package delivery

import "project3/group3/domain"

type InsertFormat struct {
	Name      		string `json:"name" form:"name"`
	Size     		string `json:"size" form:"size"`
	PriceProduct  	int `json:"priceproduct" form:"priceproduct"`
}

func (i *InsertFormat) ToModel() domain.Product {
	return domain.Product{
		Name:	i.Name,
		Size:	i.Size,
		PriceProduct: i.PriceProduct
	}
}

