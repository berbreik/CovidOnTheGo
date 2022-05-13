package assignment

import (
	"assignment/driversql"
	handler "assignment/http"
	"assignment/service"
	"assignment/store"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	sqlconf := driversql.Mysqlconfig{
		Driver:   "mysql",
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Password: "root",
		Dbname:   "covid",
	}

	db,err := driversql.ConnectToMySQL(sqlconf)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("error closing connection to sql server %v", err)

		}
	}(db)
	store := store.NewDatastore(db)
	serve := service.NewService(store)
	handlermain := handler.NewHandler(serve)
	
	route := mux.NewRouter()
	route.HandleFunc("/patient", handlermain.AddHandler).Methods("POST")
	route.HandleFunc("/patient/{id}", handlermain.GetHandler).Methods("GET")
	route.HandleFunc("/patient/{id}", handlermain.UpdateHandler).Methods("PUT")
	route.HandleFunc("/patient/{id}", handlermain.RemoveHandler).Methods("DELETE")
	route.HandleFunc("/patient", handlermain.GetAllHandler).Methods("GET")
	
	err = http.ListenAndServe(":8080", route)
	if err != nil {
		log.Fatal(err)
	}

}
