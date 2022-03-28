package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hrqcds/go-rest-api/controllers"
	"github.com/hrqcds/go-rest-api/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/new_personalidades", controllers.NovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.PersonalidadeEspecifica).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.UpdatePersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r))) // configurando cors
}
