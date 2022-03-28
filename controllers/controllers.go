package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hrqcds/go-rest-api/database"
	"github.com/hrqcds/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var p []models.Personalidade

	database.DB.Find(&p)

	json.NewEncoder(w).Encode(p)
}

func PersonalidadeEspecifica(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r) // parametros

	id := vars["id"] // especifciando

	var p models.Personalidade

	database.DB.First(&p, id)

	json.NewEncoder(w).Encode(p)

}

func NovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade

	json.NewDecoder(r.Body).Decode(&personalidade)

	database.DB.Create(&personalidade)

	json.NewEncoder(w).Encode(personalidade)
}

func DeletePersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var p models.Personalidade

	database.DB.First(&p, id)

	database.DB.Delete(&p)

	json.NewEncoder(w).Encode(p)
}

func UpdatePersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var p models.Personalidade

	database.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)

	json.NewEncoder(w).Encode(p)
}
