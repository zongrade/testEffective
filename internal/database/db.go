// Package db provides usage of database
package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/zongrade/testeffective/internal/tools"
)

// DBconfig placeholder.
type DBconfig struct {
	MigrationDir string
	Connection   *sql.DB
}

// DBcfg конфиг БД.
var DBcfg *DBconfig

// SetupConnection настроивает коннект к БД.
func SetupConnection() error {
	DBcfg = &DBconfig{}

	if err := tools.IsMigrationDirExist(); err != nil {
		return fmt.Errorf("подключение к базе данных не произведено: %w", err)
	}

	fmt.Println(tools.PathMigrationDir())
	DBcfg.MigrationDir = tools.PathMigrationDir()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("ошибка подключения к базе данных: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return fmt.Errorf("ошибка при проверке подключения: %w", err)
	}

	DBcfg.Connection = db

	return nil
}
