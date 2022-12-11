package routes

import (
	"JSE_API/pkg/controllers"
	"github.com/julienschmidt/httprouter"
)
type Routes struct {
	router *httprouter.Router
}

func (r *Routes) ApiRoutes() {
	router := r.router

	 TradeSummary := controllers.TradeSummary{}

	//User Routes
	router.GET("/TradeSummary", TradeSummary.GetStockAdvancing)

}
