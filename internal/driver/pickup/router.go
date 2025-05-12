package pickup

import (
	"gojeksrepo/ent"

	"github.com/labstack/echo/v4"
	"github.com/segmentio/kafka-go"
)

type PickupRouter struct {
	handler *PickupHandler
}

func PickupInitRouter(db *ent.Client, kafkaClient *kafka.Writer) *PickupRouter {
	hndlrPickup := PickupNewService(db, kafkaClient)
	svcHandler := InitPickupHandler(hndlrPickup)

	return &PickupRouter{
		handler: svcHandler,
	}
}

func (ctrl *PickupRouter) RegisterPickupRouter(g *echo.Group) {
	g.POST("/order", ctrl.handler.AcceptOrderByOrderId)
}
