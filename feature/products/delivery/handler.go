package delivery

import (
	"log"
	"net/http"
	"project3/group3/domain"
	_helper "project3/group3/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	productUseCase domain.ProductUseCase
}

func New(ps domain.ProductUseCase) domain.ProductHandler {
	return &productHandler{
		productUseCase: ps,
	}
}

func (ph *productHandler) InsertProductHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertProductFormat
		err := c.Bind(&tmp)
		if err != nil {
			log.Println("cannot parse data", err)
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
		}

		dataProduct := tmp.ToModel()
		row, err := ph.productUseCase.InsertProduct(dataProduct)
		if row == -1 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("please make sure all fields are filled in correctly"))
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("your product is already registered"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
	}
}

func (ph *productHandler) DeleteProductHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idproduct := c.Param("id")
		id, _ := strconv.Atoi(idproduct)
		if id == 0 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("you dont have access"))
		}
		row, errDel := ph.productUseCase.DeleteProduct(id)

		if errDel != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to delete data product"))
		}

		if row != 1 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to delete data product"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
	}
}
