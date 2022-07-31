package factory

import (
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
	userDelivery.New(e, useCase)
}
