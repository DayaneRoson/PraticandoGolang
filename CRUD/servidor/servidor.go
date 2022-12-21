package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type usuario struct {
	ID    uint16 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Coud not read json"))
		return
	}
	var usuario usuario
	if erro = json.Unmarshal(body, &usuario); erro != nil {
		w.Write([]byte("Coud not parse json"))
		return
	}
	db, erro := banco.Connection()
	if erro != nil {
		w.Write([]byte("Coud not connect to database"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Coud not create statement"))
		return
	}
	defer statement.Close()

	insertion, error := statement.Exec(usuario.Nome, usuario.Email)
	if error != nil {
		w.Write([]byte("Could not execute query string"))
		return
	}

	idInserted, error := insertion.LastInsertId()
	if error != nil {
		w.Write([]byte("Could not fetch id"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Query executed sucessfully. ID: %d", idInserted)))
}
