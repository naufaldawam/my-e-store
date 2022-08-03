package delivery

import (
	"project3/group3/domain"
	_middleware "project3/group3/feature/common"

	"github.com/labstack/echo/v4"
)

func RouteProduct(e *echo.Echo, dp domain.ProductHandler) {
	e.POST("/products", dp.InsertProductHandler(), _middleware.JWTMiddleware())       //admin
	e.DELETE("/products/:id", dp.DeleteProductHandler(), _middleware.JWTMiddleware()) //admin
	e.GET("/products/:id", dp.GetById(), _middleware.JWTMiddleware())                 //user & admin
	e.GET("/products", dp.GetAll())                                                   // user admin
	e.PUT("/products/:id", dp.Update(), _middleware.JWTMiddleware())                  //admin
}
