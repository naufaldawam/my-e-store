// package delivery

// import (
// 	"project3/group3/domain"
// 	_middleware "project3/group3/feature/common"

// 	"github.com/labstack/echo/v4"
// )

// func RouteUser(e *echo.Echo, du domain.UserHandler) {
// 	e.POST("/login", du.LoginAuth())
// 	e.POST("/register", du.InsertUser())
// 	e.GET("/users", du.GetProfile(), _middleware.JWTMiddleware())
// 	e.DELETE("/users", du.DeleteById(), _middleware.JWTMiddleware())
// 	e.PUT("/users", du.UpdateUser(), _middleware.JWTMiddleware())
// }
