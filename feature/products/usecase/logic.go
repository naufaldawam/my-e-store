package usecase

import (
	"errors"
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

func (pc *productUseCase) InsertProduct(newProduct domain.Product) (row int, err error) {
	if newProduct.ProductName == "" || newProduct.Stock == 0 || newProduct.Price == 0 {
		return -1, errors.New("please make sure all fields are filled in correctly")
	}
	row, err = pc.productData.InsertProductDB(newProduct)
	return row, err
}

func (pc *productUseCase) DeleteProduct(productID int) (row int, err error) {
	row, err = pc.productData.DeleteProductDB(productID)
	return row, err
}
