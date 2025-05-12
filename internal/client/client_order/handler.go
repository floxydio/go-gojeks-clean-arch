package clientorder

import (
	"gojeksrepo/config"
	"gojeksrepo/internal/client/client_order/dto"
	"gojeksrepo/pkg"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type ClientOrderHandler struct {
	clientOrderSvc *ClientOrderService
}

func NewClientOrderHandlerInit(clientOrderS *ClientOrderService) *ClientOrderHandler {
	return &ClientOrderHandler{
		clientOrderSvc: clientOrderS,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (svc *ClientOrderHandler) FindOrderNotSuccessfull(c echo.Context) error {
	data, err := svc.clientOrderSvc.FindLastTransactionNotSuccess(c.Request().Context())

	if err != nil {
		return c.JSON(500, config.GlobalResponseMsg{
			Status:  500,
			Error:   true,
			Message: err.Error(),
		})
	}

	return c.JSON(200, config.GlobalResponseData{
		Status:  200,
		Error:   false,
		Data:    data,
		Message: "Successfully Get Order not successfull",
	})
}

func (svc *ClientOrderHandler) CreateOrder(c echo.Context) error {
	var formData dto.TripRequest

	if errBind := c.Bind(&formData); errBind != nil {
		return c.JSON(500, config.GlobalResponseMsg{
			Status:  500,
			Error:   true,
			Message: errBind.Error(),
		})
	}

	dataCreated, err := svc.clientOrderSvc.CreateOrder(formData, c.Request().Context())

	if err != nil {
		return c.JSON(500, config.GlobalResponseMsg{
			Status:  500,
			Error:   true,
			Message: err.Error(),
		})
	}

	return c.JSON(200, config.GlobalResponseData{
		Status:  200,
		Error:   false,
		Data:    dataCreated,
		Message: "Successfully Create Order - Waiting for pick",
	})
}

func (svc *ClientOrderHandler) ClientNotifWebsocket(c echo.Context) error {
	userId := c.Param("user_id")
	wsckt, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		return c.JSON(500, config.GlobalResponseMsg{
			Status:  500,
			Error:   true,
			Message: err.Error(),
		})
	}

	pkg.SaveUserSocket(userId, wsckt)
	log.Println("User connected:", userId)

	defer func() {
		pkg.RemoveUserSocket(userId)
		wsckt.Close()
	}()

	for {
		_, _, err := wsckt.ReadMessage()
		if err != nil {
			log.Println("User disconnected:", userId)
			break
		}
	}

	return nil
}

func (svc *ClientOrderHandler) DriverNotifWebsocket(c echo.Context) error {
	driverId := c.Param("driver_id")

	wsck, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		return c.JSON(500, config.GlobalResponseMsg{
			Status:  500,
			Error:   true,
			Message: err.Error(),
		})
	}

	pkg.SaveDriverSocket(driverId, wsck)

	defer func() {
		pkg.RemoveDriverSocket(driverId)
		wsck.Close()
	}()

	for {
		_, _, err := wsck.ReadMessage()
		if err != nil {
			log.Println("Driver disconnected:", driverId)
			break
		}
	}
	return nil
}
