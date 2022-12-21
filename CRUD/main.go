package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//CRUD - CREATE, READ, UPDATE, DELETE
	//CREATE - POST
	//READ- GET
	//UPDATE- PUT
	//DELETE - DELETE
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", servidor.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", servidor.DeleteUser).Methods(http.MethodDelete)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
