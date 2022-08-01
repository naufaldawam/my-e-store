package factory

import (
	// pd "project3/group3/feature/product/data"
	// productDelivery "project3/group3/feature/product/delivery"
	// ps "project3/group3/feature/product/usecase"

	ud "project3/group3/feature/users/data"
	userDelivery "project3/group3/feature/users/delivery"
	us "project3/group3/feature/users/usecase"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userData := ud.New(db)
	validator := validator.New()
	useCase := us.New(userData, validator)
	userHandler := userDelivery.New(useCase)
	userDelivery.RouteUser(e, userHandler)
}
