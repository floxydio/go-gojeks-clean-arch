package router

import (
	"gojeksrepo/config"
	"gojeksrepo/internal/admin"
	"gojeksrepo/internal/auth"
	clientorder "gojeksrepo/internal/client/client_order"
	"gojeksrepo/pkg"

	"github.com/labstack/echo/v4"
)

func InitRouter() *echo.Echo {
	db := config.GetDB()
	kafkaService := pkg.KafkaConnector()

	echoService := echo.New()

	authRouter := auth.AuthRoutes(db)
	authRouter.Register(echoService.Group("/auth"))

	adminRouter := admin.AdminRoutes(db)
	adminRouter.Register(echoService.Group("/admin"))

	clientorderRouter := clientorder.ClientOrderRoutes(db, kafkaService)
	clientorderRouter.Register(echoService.Group("/trx"))

	return echoService
}
