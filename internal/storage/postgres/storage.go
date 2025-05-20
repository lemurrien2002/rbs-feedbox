package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storagepostgres struct {
	db *sql.DB
}

func NewStoragepostgres(dsn string) *Storagepostgres {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		// panic("Ошибка подкл" + err.Error())
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	if err = db.Ping(); err != nil {
		// panic("Недоступн база" + err.Error())
		log.Fatalf("Недоступна база данных: %v", err)
	}
	// fmt.Println("Подключено")
	log.Println("Подключение к базе данных успешно установлено")
	return &Storagepostgres{db: db}
}
