package routes

import (
	"backend/handlers"
	"backend/pkg/mysql"
	"backend/repositories"

	"github.com/gorilla/mux"
)

func DataRoutes(r *mux.Router) {
	dataRepository := repositories.RepositoryData(mysql.DB)
	h := handlers.HandlerData(dataRepository)

	r.HandleFunc("/datas", h.ShowData).Methods("GET")
	r.HandleFunc("/detail/{id}", h.GetDataByID).Methods("GET")
	r.HandleFunc("/add", h.CreateData).Methods("POST")
	r.HandleFunc("/detail/{id}", h.UpdateData).Methods("PATCH")
	r.HandleFunc("/detail/{id}", h.DeleteData).Methods("DELETE")
}
