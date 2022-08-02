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

func (pd *productData) GetUser(idProduct int) (result domain.Product, err error) {
	var tmp Product
	res := pd.db.Preload("User").Where("id = ?", idProduct).First(&tmp)
	if res.Error != nil {
		return domain.Product{}, res.Error
	}
	return tmp.ToDomain(), nil
}

func (pd *productData) InsertProductDB(newProduct domain.Product) (result domain.Product, err error) {
	product := FromDomain(newProduct)
	res := pd.db.Create(&product)
	if res.Error != nil {
		return domain.Product{}, res.Error
	}
	if res.RowsAffected != 1 {
		return domain.Product{}, errors.New("failed to insert product")
	}
	return product.ToDomain(), nil
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

func (pd *productData) SelectData(limit, offset int) (data []domain.Product, err error) {
	dataProduct := []Product{}
	res := pd.db.Preload("User").Limit(limit).Offset(offset).Find(&dataProduct)
	if res.Error != nil {
		return []domain.Product{}, nil
	}
	return ParseToArrProduct(dataProduct), nil
}
func (pd *productData) SelectDataById(id int) (data domain.Product, err error) {
	tmp := Product{}
	res := pd.db.Preload("User").Find(&tmp, id)
	if res.Error != nil {
		return domain.Product{}, res.Error
	}
	return tmp.ToDomain(), nil
}
