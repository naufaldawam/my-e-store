package delivery

import (
	"project3/group3/domain"
	_middleware "project3/group3/feature/common"

	"github.com/labstack/echo/v4"
)

func RouteProduct(e *echo.Echo, dp domain.ProductHandler) {
	e.POST("/products", dp.InsertProductHandler(), _middleware.JWTMiddleware())
	e.DELETE("/products/:id", dp.DeleteProductHandler(), _middleware.JWTMiddleware())
}
