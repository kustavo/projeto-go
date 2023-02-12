package router

import (
	"github.com/kustavo/benchmark/go/application/interfaces"
	"github.com/kustavo/benchmark/go/infrastructure/http/controller"
)

func GetUserRouter(application *interfaces.Application) []route {
	userClr := controller.NewUserController(application)
	routes := make([]route, 0)

	routes = append(routes, route{method: "POST", url: "/user/", handler: userClr.Create, requireAuth: false})
	routes = append(routes, route{method: "GET", url: "/user/", handler: userClr.GetById, requireAuth: true})
	routes = append(routes, route{method: "PUT", url: "/user/", handler: userClr.Update, requireAuth: true})
	routes = append(routes, route{method: "DELETE", url: "/user/", handler: userClr.Delete, requireAuth: true})
	routes = append(routes, route{method: "GET", url: "/user/get-by-email", handler: userClr.GetByEmail, requireAuth: true})
	routes = append(routes, route{method: "POST", url: "/login", handler: userClr.Login, requireAuth: false})
	routes = append(routes, route{method: "POST", url: "/logout", handler: userClr.Logout, requireAuth: true})
	routes = append(routes, route{method: "POST", url: "/refresh-auth", handler: userClr.RefreshAuth, requireAuth: true})

	return routes
}
