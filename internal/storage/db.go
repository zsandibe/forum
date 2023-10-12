package storage

import (
	"database/sql"
	"forum/internal/config"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// type Database struct {
// 	db *sql.DB
// }

func NewSqlite3(config config.Config) (*sql.DB, error) {
	// Подключение к DB
	db, err := sql.Open(config.DB.Driver, config.DB.Dsn)
	// fmt.Println(err)
	if err != nil {
		// fmt.Println("OK")
		return nil, err
	}
	// Проверка соеденений с DB
	if err = db.Ping(); err != nil {
		return nil, err
	}
	// Создает таблицы
	if err = CreateTables(db, config); err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTables(db *sql.DB, config config.Config) error {
	file, err := os.ReadFile(config.Migrate)
	if err != nil {
		return err
	}
	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		_, err := db.Exec(request)
		if err != nil {
			return err
		}
	}
	return err
}
