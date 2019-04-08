package routes

import (
	"go-api/controllers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", controllers.GetHome).Methods("GET","OPTIONS")
	r.HandleFunc("/directorios", controllers.GetAgenda).Methods("GET","OPTIONS")
	r.HandleFunc("/directorios/{uid}", controllers.GetAgendaId).Methods("GET","OPTIONS")
	r.HandleFunc("/directorios", controllers.PostAgenda).Methods("POST","OPTIONS")
	r.HandleFunc("/directorios/{uid}", controllers.PutAgenda).Methods("PUT","OPTIONS")
	r.HandleFunc("/directorios/{uid}", controllers.DeleteAgenda).Methods("DELETE",)

	return r
}
