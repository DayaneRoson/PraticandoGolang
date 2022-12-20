package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConnection := "golang:User@course@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConnection)

	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()
	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println("Conexão está aberta!")
	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()
	fmt.Println(linhas)
}
