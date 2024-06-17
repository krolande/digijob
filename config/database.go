package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Permet la connexion à la BD
func DBConn() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:@tcp(localhost:3307)/digiJob")
	if err != nil {
		panic(err.Error())
	}
	// ce return veut dire qu'il doit renvoyer toutes les variables de return ?parseTime=true
	return
}
