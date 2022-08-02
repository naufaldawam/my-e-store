package data

import (
	"errors"
	"log"
	"project3/group3/domain"

	"gorm.io/gorm"
)

type productData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ProductData {
	return &productData{
		db: db,
	}
}

func (pd *productData) InsertProductDB(newProduct domain.Product) (row int, err error) {
	product := FromModel(newProduct)
	resultCreate := pd.db.Create(&product)

	if resultCreate.Error != nil {
		return 0, resultCreate.Error
	}
	if resultCreate.RowsAffected != 1 {
		return 0, errors.New("failed to insert product, your product is already registered")
	}

	return int(resultCreate.RowsAffected), nil
}

func (pd *productData) DeleteProductDB(productID int) (row int, err error) {
	res := pd.db.Delete(&Product{}, productID)
	if res.Error != nil {
		log.Println("cannot delete product", res.Error.Error())
		return 0, res.Error
	}
	if res.RowsAffected < 1 {
		log.Println("no product deleted", res.Error.Error())
		return 0, errors.New("failed to delete product")
	}
	return int(res.RowsAffected), nil
}
