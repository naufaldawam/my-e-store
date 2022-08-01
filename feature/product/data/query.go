package data

import (
	// "errors"
	// "fmt"
	// "log"
	"project3/group3/domain"
	// "project3/group3/feature/common"
	// product "project3/group3/feature/product"

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

// func (ud *userData) Insert(newUser domain.User) (row int, err error) {
// 	user := FromModel(newUser)
// 	passwordHashed, errorHash := _bcrypt.GenerateFromPassword([]byte(user.Password), 10)
// 	if errorHash != nil {
// 		fmt.Println("Error hash", errorHash.Error())
// 	}
// 	user.Password = string(passwordHashed)
// 	resultCreate := ud.db.Create(&user)
// 	if resultCreate.Error != nil {
// 		return 0, resultCreate.Error
// 	}
// 	if resultCreate.RowsAffected != 1 {
// 		return 0, errors.New("failed to insert data, your email is already registered")
// 	}
// 	return int(resultCreate.RowsAffected), nil
// }

// func (ud *userData) LoginData(authData user.LoginModel) (token, name string, err error) {
// 	userData := User{}
// 	result := ud.db.Where("email = ?", authData.Email).First(&userData)

// 	if result.Error != nil {
// 		return "", "", result.Error
// 	}

// 	if result.RowsAffected != 1 {
// 		return "", "", errors.New("failed to login")
// 	}

// 	errCrypt := _bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(authData.Password))

// 	if errCrypt != nil {
// 		return "", "", errors.New("password incorrect")
// 	}
// 	token, _ = common.GenerateToken(int(userData.ID))
// 	return token, userData.Name, nil
// }
// func (ud *userData) GetSpecific(userID int) (domain.User, error) {
// 	var tmp User
// 	err := ud.db.Where("ID = ?", userID).First(&tmp).Error
// 	if err != nil {
// 		log.Println("There is a problem with data", err.Error())
// 		return domain.User{}, err
// 	}

// 	return tmp.ToModel(), nil
// }
// func (ud *userData) DeleteData(userID int) (row int, err error) {
// 	res := ud.db.Delete(&User{}, userID)
// 	if res.Error != nil {
// 		log.Println("cannot delete data", res.Error.Error())
// 		return 0, res.Error
// 	}
// 	if res.RowsAffected < 1 {
// 		log.Println("no data deleted", res.Error.Error())
// 		return 0, errors.New("failed to delete data ")
// 	}
// 	return int(res.RowsAffected), nil
// }
// func (ud *userData) UpdateData(data map[string]interface{}, idFromToken int) (row int, err error) {
// 	res := ud.db.Model(&User{}).Where("id = ?", idFromToken).Updates(data)
// 	if res.Error != nil {
// 		return 0, res.Error
// 	}
// 	if res.RowsAffected != 1 {
// 		return 0, errors.New("failed update data")
// 	}
// 	return int(res.RowsAffected), nil
// }
