package admin

import (
	"github.com/labstack/echo/v4"
	"gojeksrepo/ent"
	"gojeksrepo/internal/middleware"
)

type AdminRouter struct {
	adminHandler *AdminHandler
}

func AdminRoutes(db *ent.Client) *AdminRouter {
	svc := NewServiceAdmin(db)
	handler := NewAdminController(svc)

	return &AdminRouter{
		adminHandler: handler,
	}
}

func (ctrl *AdminRouter) Register(g *echo.Group) {
	g.POST("/approve-admin", ctrl.adminHandler.ApproveByAdmin, middleware.AuthMiddleware)
}
