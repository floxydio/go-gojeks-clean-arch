package admin

import (
	"gojeksrepo/config"
	"gojeksrepo/internal/admin/dto"
	"gojeksrepo/pkg"

	"github.com/labstack/echo/v4"
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

func (s *AdminHandler) SignUpAdmin(c echo.Context) error {
	var body dto.SignUpAdmin

	if errBind := c.Bind(&body); errBind != nil {
		return c.JSON(400, config.GlobalResponseMsg{
			Status:  400,
			Error:   true,
			Message: "Something went wrong when bind",
		})
	}

	errService := s.adminService.SignupAdmin(body, c.Request().Context())

	if errService != nil {
		return c.JSON(500, config.GlobalResponseMsg{
			Status:  500,
			Error:   true,
			Message: "Something Went Wrong",
		})
	}

	return c.JSON(201, config.GlobalResponseMsg{
		Status:  201,
		Error:   false,
		Message: "Successfully Sign Up Admin",
	})
}

func (s *AdminHandler) SignInAdmin(c echo.Context) error {
	var body dto.SignInAdmin

	if err := c.Bind(&body); err != nil {
		return c.JSON(400, config.GlobalResponseMsg{
			Status:  400,
			Error:   true,
			Message: "Something went wrong when bind",
		})
	}

	findExisting, errFindExisting := s.adminService.FindExistUser(body, c.Request().Context())

	if errFindExisting != nil {
		return c.JSON(500, config.GlobalResponseMsg{
			Status:  500,
			Error:   true,
			Message: "Something went wrong when findExisting",
		})
	}

	if findExisting == nil {
		return c.JSON(400, config.GlobalResponseMsg{
			Status:  400,
			Error:   true,
			Message: "User not found",
		})
	}

	checkPassword := pkg.CheckPasswordHash(findExisting.Password, body.Password)

	if !checkPassword {
		return c.JSON(400, config.GlobalResponseMsg{
			Status:  400,
			Error:   true,
			Message: "Password salah",
		})
	}
	return c.JSON(200, config.GlobalResponseData{
		Status:  200,
		Error:   false,
		Message: "Successfully Login",
	})

}
