package router

import (
	"gojeksrepo/config"
	"gojeksrepo/internal/admin"
	"gojeksrepo/internal/auth"
	clientorder "gojeksrepo/internal/client/client_order"
	"gojeksrepo/internal/driver/pickup"
	"gojeksrepo/pkg"

	"github.com/labstack/echo/v4"
)

func InitRouter() *echo.Echo {
	db := config.GetDB()
	kafkaService := pkg.KafkaConnector("order-created")
	kafkaDriverAssignedService := pkg.KafkaConnector("order-assigned")

	echoService := echo.New()

	authRouter := auth.AuthRoutes(db)
	authRouter.Register(echoService.Group("/auth"))

	adminRouter := admin.AdminRoutes(db)
	adminRouter.Register(echoService.Group("/admin"))

	clientorderRouter := clientorder.ClientOrderRoutes(db, kafkaService)
	clientorderRouter.Register(echoService.Group("/trx"))

	pickupRouter := pickup.PickupInitRouter(db, kafkaDriverAssignedService)
	pickupRouter.RegisterPickupRouter(echoService.Group("/pick"))

	return echoService
}
