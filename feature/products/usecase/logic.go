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
