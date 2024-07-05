package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

// InitDB устанавливает соединение с базой данных.
func InitDB() error {
	// Получаем учетные данные базы данных из переменных окружения.
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Строим строку подключения.
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbname)

	// Устанавливаем соединение с базой данных.
	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("не удалось подключиться к базе данных: %v", err)
	}

	// Пингуем базу данных и проверяем соединение.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		return fmt.Errorf("не удалось пинговать базу данных: %v", err)
	}

	return nil
}

// GetDB возвращает указатель на объект sql.DB.
func GetDB() *sql.DB {
	return db
}
