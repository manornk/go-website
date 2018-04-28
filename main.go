package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manornk/go-website/utils"

	_ "github.com/go-sql-driver/mysql" // MySQL driver

	"github.com/jmoiron/sqlx"
)

type dbHandler struct {
	db *sqlx.DB
}

func main() {

	db, err := utils.InitMySQL("database_name")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	h := &dbHandler{db}

	router := mux.NewRouter()

	router.HandleFunc("/", h.ReadHomePage).Methods("GET")

	router.HandleFunc("/{slug}", h.ReadPage).Methods("GET")

	router.HandleFunc("/{year}/{month}/{day}/{slug}", ReadPost).Methods("GET")

	http.ListenAndServe(":8080", router)
}
