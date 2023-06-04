package Controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "github.com/YuriBertoldi/Go-APIRest/conexaodb"
	models "github.com/YuriBertoldi/Go-APIRest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var ListPersonalidade []models.Personalidade

	db.DataBase.Find(&ListPersonalidade)

	json.NewEncoder(w).Encode(ListPersonalidade)
}

func RetornaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var personalidade models.Personalidade

	db.DataBase.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func CriarPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&personalidade)
	db.DataBase.Create(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}

func EditarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var personalidade models.Personalidade
	db.DataBase.First(&personalidade, id)
	json.NewDecoder(r.Body).Decode(&personalidade)
	db.DataBase.Save(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}

func DeletarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	var personalidade models.Personalidade
	db.DataBase.Delete(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}
