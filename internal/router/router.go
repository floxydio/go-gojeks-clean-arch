package router

import (
	"github.com/labstack/echo/v4"
	"gojeksrepo/config"
	"gojeksrepo/internal/admin"
	"gojeksrepo/internal/auth"
)

func InitRouter() *echo.Echo {
	db := config.GetDB()
	echoService := echo.New()

	authRouter := auth.AuthRoutes(db)
	authRouter.Register(echoService.Group("/auth"))

	adminRouter := admin.AdminRoutes(db)
	adminRouter.Register(echoService.Group("/admin"))

	return echoService
}
