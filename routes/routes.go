package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	ProfileRoutes(r)
	AuthRoutes(r)
	ProductRoutes(r)
	TopingRoutes(r)
	TransactionRoutes(r)
	CartRoutes(r)
}
