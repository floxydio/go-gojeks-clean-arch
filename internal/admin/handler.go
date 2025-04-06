package admin

import (
	"github.com/labstack/echo/v4"
	"gojeksrepo/config"
	"gojeksrepo/internal/admin/dto"
)

type AdminHandler struct {
	adminService *Service
}

func NewAdminController(service *Service) *AdminHandler {
	return &AdminHandler{
		adminService: service,
	}
}

func (s *AdminHandler) ApproveByAdmin(c echo.Context) error {
	var approveForm dto.ApprovalAdminForm

	if err := c.Bind(&approveForm); err != nil {
		return c.JSON(400, config.GlobalResponseMsg{
			Status:  400,
			Error:   true,
			Message: "Something went wrong when bind",
		})
	}

	errService := s.adminService.ApproveAdmin(approveForm, c.Request().Context())

	if errService != nil {
		return c.JSON(500, config.GlobalResponseMsg{
			Status:  500,
			Error:   true,
			Message: "Something went wrong",
		})
	}
	return c.JSON(200, config.GlobalResponseMsg{
		Status:  200,
		Error:   false,
		Message: "Successfully Update Status Active / Inactive",
	})

}
