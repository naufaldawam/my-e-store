package delivery

import (
	"log"
	"net/http"
	"project3/group3/domain"
	_middleware "project3/group3/feature/common"
	product "project3/group3/feature/product"
	_helper "project3/group3/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

type orderHandler struct {
	ordersUseCase domain.OrdersUseCase
}

func New(ps domain.OrdersUseCase) domain.OrdersUseCase {
	return &ordersUseCase{
		ordersUseCase: os,
	}

}

// func (uh *userHandler) InsertUser() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var tmp InsertFormat
// 		err := c.Bind(&tmp)
// 		if err != nil {
// 			log.Println("cannot parse data", err)
// 			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
// 		}

// 		dataUser := tmp.ToModel()
// 		row, err := uh.userUsecase.AddUser(dataUser)
// 		if row == -1 {
// 			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("please make sure all fields are filled in correctly"))
// 		}
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("your email is already registered"))
// 		}
// 		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
// 	}
// }

// func (uh *userHandler) LoginAuth() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		authData := user.LoginModel{}
// 		c.Bind(&authData)
// 		token, name, e := uh.userUsecase.Login(authData)
// 		if e != nil {
// 			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("email or password incorrect"))
// 		}

// 		data := map[string]interface{}{
// 			"token": token,
// 			"name":  name,
// 		}
// 		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("login success", data))
// 	}
// }

// func (uh *userHandler) GetProfile() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		id, _ := _middleware.ExtractData(c)
// 		data, err := uh.userUsecase.GetProfile(id)

// 		if err != nil {
// 			if strings.Contains(err.Error(), "not found") {
// 				return c.JSON(http.StatusNotFound, err.Error())
// 			} else {
// 				return c.JSON(http.StatusInternalServerError, err.Error())
// 			}
// 		}
// 		return c.JSON(http.StatusOK, map[string]interface{}{
// 			"message": "data found",
// 			"data":    data,
// 		})
// 	}
// }

// func (uh *userHandler) DeleteById() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		idFromToken, _ := _middleware.ExtractData(c)
// 		if idFromToken == 0 {
// 			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("you dont have access"))
// 		}
// 		row, errDel := uh.userUsecase.DeleteCase(idFromToken)
// 		if errDel != nil {
// 			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to delete data user"))
// 		}
// 		if row != 1 {
// 			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to delete data user"))
// 		}
// 		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
// 	}
// }

// func (uh *userHandler) UpdateUser() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var tmp InsertFormat
// 		idFromToken, _ := _middleware.ExtractData(c)
// 		err := c.Bind(&tmp)
// 		if err != nil {

// 			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to bind data, check your input"))
// 		}
// 		row, err := uh.userUsecase.UpdateCase(tmp.ToModel(), idFromToken)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed update data users, cek your input email"))
// 		}
// 		if row == 0 {
// 			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed update data users, no data"))
// 		}

// 		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
// 	}
// }
