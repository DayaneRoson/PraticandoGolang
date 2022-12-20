package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver de conexão mysql
)

// Connection abre a conexão com o banco de dados
func Connection() (*sql.DB, error) {
	stringConnection := "golang:User@course@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConnection)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
