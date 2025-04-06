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

func (r *AuthRouter) Register(g *echo.Group) {
	g.POST("/sign-up-user", r.userHandler.SignUp)
	g.POST("/sign-up-driver", r.userHandler.SignUpDriver)
	g.POST("/sign-in-user", r.userHandler.SignInUser)
	g.POST("/sign-in-driver", r.userHandler.SignInUserDriver)
}
