package DB

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "go-movies"
)

func SetupDB() *sql.DB {
	info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	DB, err := sql.Open("postgres", info)
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	checkErr(err)

	return DB
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	} else {
		fmt.Println("conexion exitosa")
	}
}
