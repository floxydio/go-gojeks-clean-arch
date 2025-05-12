package pickup

import (
	"gojeksrepo/config"
	"gojeksrepo/internal/driver/pickup/dto"

	"github.com/labstack/echo/v4"
)

type PickupHandler struct {
	svc *PickupServiceClient
}

func InitPickupHandler(service *PickupServiceClient) *PickupHandler {
	return &PickupHandler{
		svc: service,
	}
}

func (r *PickupHandler) AcceptOrderByOrderId(c echo.Context) error {
	var formData dto.OrderPickup

	if errBind := c.Bind(&formData); errBind != nil {
		return c.JSON(400, config.GlobalResponseMsg{
			Status:  400,
			Error:   true,
			Message: errBind.Error(),
		})
	}

	err := r.svc.AcceptOrder(formData.OrderId, formData.DriverId, c.Request().Context())

	if err != nil {
		return c.JSON(500, config.GlobalResponseMsg{
			Status:  500,
			Error:   true,
			Message: err.Error(),
		})
	}

	return c.JSON(200, config.GlobalResponseMsg{
		Status:  200,
		Error:   false,
		Message: "Successfully Take Order",
	})

}
