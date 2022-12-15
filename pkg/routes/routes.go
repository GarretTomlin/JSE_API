package routes

import (
	"JSE_API/pkg/controllers"
	"github.com/julienschmidt/httprouter"
)

type Routes struct {
	router *httprouter.Router
}

// New creates a new instance of the Routes struct
func New() *Routes {
	router := httprouter.New()
	return &Routes{router: router}
}

// GetHttpRouter returns the underlying httprouter.Router
func (r *Routes) GetHttpRouter() *httprouter.Router {
	return r.router
}

func (r *Routes) ApiRoutes() {
	router := r.router

	TradeSummary := controllers.TradeSummary{}

	//Summary Routes
	router.GET("/TradeSummary", TradeSummary.GetStockAdvancing)

	//health checks
	router.GET("/healthcheck", controllers.HealthCheck)
}
