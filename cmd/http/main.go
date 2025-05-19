package main

import (
	"fmt"
	"net/http"
	"os"

	"rbs-feedbox/internal/service"
	"rbs-feedbox/internal/storage/postgres"
	httproutes "rbs-feedbox/internal/transport"
)

func main() {
	// Формируем строку для подключения к базе данных
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	// Подключаемся к БД
	storage := postgres.NewStoragepostgres(dsn)
	// Инициализируем структуру сервисов (бл)
	svc := service.New(storage)

	httproutes.Register(svc)
	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
