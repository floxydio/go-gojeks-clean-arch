package auth

import (
	"github.com/labstack/echo/v4"
	"gojeksrepo/ent"
)

type AuthRouter struct {
	userHandler *AuthHandler
}

func AuthRoutes(db *ent.Client) *AuthRouter {
	svc := NewAuthService(db)
	h := NewAuthHandler(svc)

	return &AuthRouter{
		userHandler: h,
	}
}

func (ctrl *AuthRouter) Register(g *echo.Group) {
	g.POST("/sign-up-user", ctrl.userHandler.SignUp)
	g.POST("/sign-up-driver", ctrl.userHandler.SignUpDriver)
	g.POST("/sign-in-user", ctrl.userHandler.SignInUser)
	g.POST("/sign-in-driver", ctrl.userHandler.SignInUserDriver)
}
