package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func InitDB() {
	var err error
	dsn := "postgres://user:password@localhost:5432/mydb?sslmode=disable"
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
}

func SaveBatch(data []string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO records (data) VALUES ($1)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, record := range data {
		_, err = stmt.Exec(record)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
