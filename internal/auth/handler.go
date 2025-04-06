package auth

import (
	"github.com/labstack/echo/v4"
	"gojeksrepo/config"
	"gojeksrepo/internal/auth/dto"
)

type AuthHandler struct {
	authService *DbService
}

func NewAuthHandler(authS *DbService) *AuthHandler {
	return &AuthHandler{authService: authS}
}

func (authS *AuthHandler) SignUp(c echo.Context) error {
	var body dto.SignUp

	if err := c.Bind(&body); err != nil {
		return c.JSON(400, config.GlobalResponseMsg{
			Status:  500,
			Error:   true,
			Message: "Something went wrong when bind",
		})
	}
	errSignUpDriver := authS.authService.SignUpUser(body, c.Request().Context())

	if errSignUpDriver != nil {
		return c.JSON(500, config.GlobalResponseMsg{
			Status:  500,
			Error:   true,
			Message: "Something went wrong" + errSignUpDriver.Error(),
		})
	}

	return c.JSON(200, config.GlobalResponseMsg{
		Status:  200,
		Error:   false,
		Message: "Successfully Sign Up User",
	})
}

func (authS *AuthHandler) SignUpDriver(c echo.Context) error {
	var body dto.SignUpDriver

	if err := c.Bind(&body); err != nil {
		return c.JSON(400, config.GlobalResponseMsg{
			Status:  500,
			Error:   true,
			Message: "Something went wrong when bind",
		})
	}
	errSignUpDriver := authS.authService.SignUpDriver(body, c.Request().Context())

	if errSignUpDriver != nil {
		return c.JSON(500, config.GlobalResponseMsg{
			Status:  500,
			Error:   true,
			Message: "Something went wrong" + errSignUpDriver.Error(),
		})
	}
	return c.JSON(200, config.GlobalResponseMsg{
		Status:  200,
		Error:   false,
		Message: "Successfully Sign Up Driver",
	})
}

func (authS *AuthHandler) SignInUser(c echo.Context) error {
	var body dto.SignInForm

	if err := c.Bind(&body); err != nil {
		return c.JSON(400, config.GlobalResponseMsg{
			Status:  400,
			Error:   true,
			Message: "Something went wrong when bind",
		})
	}

	data, errSignIn := authS.authService.SignInUser(body, c.Request().Context())

	if errSignIn != nil {
		return c.JSON(500, config.GlobalResponseMsg{
			Status:  500,
			Error:   true,
			Message: "Something went wrong",
		})
	}
	return c.JSON(data.Status, data)
}

func (authS *AuthHandler) SignInUserDriver(c echo.Context) error {
	var body dto.SignInFormDriver

	if err := c.Bind(&body); err != nil {
		return c.JSON(400, config.GlobalResponseMsg{
			Status:  400,
			Error:   true,
			Message: "Something went wrong when bind",
		})
	}

	data, errSignIn := authS.authService.SignInUserDriver(body, c.Request().Context())

	if errSignIn != nil {
		return c.JSON(500, config.GlobalResponseMsg{
			Status:  500,
			Error:   true,
			Message: "Something went wrong",
		})
	}
	return c.JSON(data.Status, data)
}
