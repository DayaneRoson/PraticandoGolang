package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint16 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CreateUser creates an user on database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, erro := io.ReadAll(r.Body)
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

// GetUsers fetches all users saved on database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, error := banco.Connection()
	if error != nil {
		w.Write([]byte("Could not connect to database"))
		return
	}
	defer db.Close()

	rows, error := db.Query("select * from usuarios")
	if error != nil {
		w.Write([]byte("Coud not fetch users"))
		return
	}
	defer rows.Close()

	var usuarios []usuario
	for rows.Next() {
		var usuario usuario
		if error := rows.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); error != nil {
			w.Write([]byte("Could not scan user, format may be wrong"))
			return
		}

		usuarios = append(usuarios, usuario)
	}
	w.WriteHeader(http.StatusOK)
	if error := json.NewEncoder(w).Encode(usuarios); error != nil {
		w.Write([]byte("Could not parse users to json, format may be wrong"))
		return
	}
}

// GetUser fetches an user saved on database
func GetUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, error := strconv.ParseUint(parameters["id"], 10, 16)
	if error != nil {
		w.Write([]byte("Could not parse URI parameter"))
		return
	}

	db, error := banco.Connection()
	if error != nil {
		w.Write([]byte("Could not connect to database"))
		return
	}

	row, error := db.Query("select * from usuarios where id = ?", ID)
	if error != nil {
		w.Write([]byte("Could not execute query string"))
		return
	}

	defer row.Close()

	var usuario usuario
	if row.Next() {
		if error := row.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); error != nil {
			w.Write([]byte("Could not scan user"))
			return
		}
	}

	if error := json.NewEncoder(w).Encode(usuario); error != nil {
		w.Write([]byte("Could not parse user, format may be wrong"))
		return
	}
}
