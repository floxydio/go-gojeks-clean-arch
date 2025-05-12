package clientorder

import (
	"gojeksrepo/ent"

	"github.com/labstack/echo/v4"
	"github.com/segmentio/kafka-go"
)

type ClientRouter struct {
	handler *ClientOrderHandler
}

func ClientOrderRoutes(db *ent.Client, kafka *kafka.Writer) *ClientRouter {
	svc := NewServiceClientOrderService(db, kafka)
	hdlr := NewClientOrderHandlerInit(svc)

	return &ClientRouter{
		handler: hdlr,
	}
}

func (ctrl *ClientRouter) Register(g *echo.Group) {
	g.GET("/history-fail", ctrl.handler.FindOrderNotSuccessfull)
	g.POST("/create-order", ctrl.handler.CreateOrder)
	g.GET("/ws/client-notif/:user_id", ctrl.handler.ClientNotifWebsocket)
	g.GET("/ws/driver-notif/:driver_id", ctrl.handler.DriverNotifWebsocket)

}
