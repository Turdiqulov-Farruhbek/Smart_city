package postgres

import (
	"database/sql"
	
	_ "github.com/lib/pq"
)


func ConnectDb() (*sql.DB, error){
	psql := "user=postgres password=root dbname=imtihon_auth_db sslmode=disable"
	db, err := sql.Open("postgres", psql)
	if err != nil {
		return nil, err
	}
	return db, nil
}
