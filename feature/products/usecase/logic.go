package usecase

import (
	"project3/group3/domain"
)

type productUseCase struct {
	productData domain.ProductData
}

func New(pd domain.ProductData) domain.ProductUseCase {
	return &productUseCase{
		productData: pd,
	}
}

func (pc *productUseCase) InsertProduct(newProduct domain.Product) (result domain.Product, err error) {
	result, err = pc.productData.InsertProductDB(newProduct)
	resultGet, _ := pc.productData.GetUser(result.ID)
	return resultGet, err
}

func (pc *productUseCase) DeleteProduct(productID int) (row int, err error) {
	row, err = pc.productData.DeleteProductDB(productID)
	return row, err
}
func (pu *productUseCase) GetAllData(limit, offset int) (data []domain.Product, err error) {
	res, err := pu.productData.SelectData(limit, offset)
	return res, err
}
func (pu *productUseCase) GetProductById(id int) (data domain.Product, err error) {
	data, err = pu.productData.SelectDataById(id)
	return data, err
}
func (pu *productUseCase) UpdateData(data domain.Product, idProduct, idFromToken int) (row int, err error) {
	reqData := map[string]interface{}{}
	if data.ProductName != "" {
		reqData["product_name"] = data.ProductName
	}
	if data.ProductImage != "" {
		reqData["product_image"] = data.ProductImage
	}
	if data.Stock != 0 {
		reqData["stock"] = data.Stock
	}
	if data.Price != 0 {
		reqData["price"] = data.Price
	}

	row, err = pu.productData.UpdateData(reqData, idProduct, idFromToken)
	return row, err
}
