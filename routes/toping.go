package routes

import (
	"waysbuck/handlers"
	"waysbuck/pkg/middleware"
	"waysbuck/pkg/mysql"
	"waysbuck/repositories"

	"github.com/gorilla/mux"
)

func TopingRoutes(r *mux.Router) {
	topingRepository := repositories.RepositoryToping(mysql.DB)
	h := handlers.HandlerToping(topingRepository)

	r.HandleFunc("/topings", h.FindTopings).Methods("GET")
	r.HandleFunc("/toping/{id}", h.GetToping).Methods("GET")
	// r.HandleFunc("/toping", h.CreateToping).Methods("POST")
	r.HandleFunc("/toping", middleware.Auth(middleware.UploadFile(h.CreateToping))).Methods("POST")
	r.HandleFunc("/toping/{id}", middleware.Auth(middleware.UploadFile(h.UpdateToping))).Methods("PATCH")
	r.HandleFunc("/toping/{id}", middleware.Auth(h.DeleteToping)).Methods("DELETE")
}
