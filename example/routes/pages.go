package routes

import (
	"net/http"
	"rhodium"
	"rhodium/example/controllers"
)

func NewPageRoutes(controller *controllers.PageController) *Routes {
	return &Routes{controller}
}

type Routes struct {
	controller *controllers.PageController
}

func (r *Routes) Index() rhodium.Route {
	return rhodium.Route{
		Method:  http.MethodGet,
		Path:    "/",
		Handler: r.controller.Index,
	}
}

func (r *Routes) Routes() []rhodium.Route {
	return []rhodium.Route{
		r.Index(),
	}
}
