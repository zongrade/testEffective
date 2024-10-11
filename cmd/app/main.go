// Package main is starting package
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	_ "github.com/zongrade/testeffective/docs"
	db "github.com/zongrade/testeffective/internal/database"
	"github.com/zongrade/testeffective/internal/routes"
	"github.com/zongrade/testeffective/internal/tools"
)

// MigrationCurrentVersion set version of most high version.
const MigrationCurrentVersion = 2

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка получения env файла")
	}

	err := db.SetupConnection()
	if err != nil {
		if errors.Is(err, tools.ErrMigrationDirNotExist) {
			fmt.Printf("ошибка fs: %v", err)
		}

		fmt.Printf("ошибка бд: %v", err)

		return
	}

	defer func() {
		if err := db.DBcfg.Connection.Close(); err != nil {
			log.Printf("ошибка закрытия соединения: %v", err)

			return
		}

		log.Println("закрыто соединение с бд")
	}()

	log.Println("Подключение к базе данных успешно")

	if closing, err := setupMigrations(); err != nil {
		closing()
		log.Printf("ошибка миграции: %v", err)
	} else {
		defer func() {
			closing()
		}()
		createServer()
	}
}

func setupMigrations() (func(), error) {
	if err := goose.SetDialect("postgres"); err != nil {
		return nil, fmt.Errorf("ошибка установки диалекта: %w", err)
	}

	closeFunc := func() {
		if err := goose.Reset(db.DBcfg.Connection, db.DBcfg.MigrationDir); err != nil {
			log.Printf("Ошибки отката миграций: %v", err)

			return
		}

		log.Printf("миграции сброшены")
	}
	if version, err := goose.EnsureDBVersion(db.DBcfg.Connection); err != nil {
		return closeFunc, fmt.Errorf("ошибка получения версии миграций: %w", err)
	} else if version > 0 {
		log.Println("попытка сброса миграций в 0")
		closeFunc()
	}

	if err := goose.UpTo(db.DBcfg.Connection, db.DBcfg.MigrationDir, MigrationCurrentVersion); err != nil {
		return closeFunc, fmt.Errorf("ошибка применения миграций: %w", err)
	}

	log.Println("Миграции успешно применены!")

	return closeFunc, nil
}

func createServer() {
	// Создаем канал для перехвата сигналов
	idleConnsClosed := make(chan struct{})
	var readTimeout time.Duration
	var WriteTimeout time.Duration
	var IdleTimeout time.Duration
	var ReadHeaderTimeout time.Duration
	var ServerShutdownTimeout time.Duration

	if readTimeoutStr, err := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT")); err != nil {
		readTimeout = time.Duration(5)
	} else {
		readTimeout = time.Duration(readTimeoutStr)
	}

	if WriteTimeoutStr, err := strconv.Atoi(os.Getenv("SERVER_WRITE_TIMEOUT")); err != nil {
		WriteTimeout = time.Duration(10)
	} else {
		WriteTimeout = time.Duration(WriteTimeoutStr)
	}

	if IdleTimeoutStr, err := strconv.Atoi(os.Getenv("SERVER_IDLE_TIMEOUT")); err != nil {
		IdleTimeout = time.Duration(120)
	} else {
		IdleTimeout = time.Duration(IdleTimeoutStr)
	}

	if ReadHeaderTimeoutStr, err := strconv.Atoi(os.Getenv("SERVER_READ_HEADER_TIMEOUT")); err != nil {
		ReadHeaderTimeout = time.Duration(2)
	} else {
		ReadHeaderTimeout = time.Duration(ReadHeaderTimeoutStr)
	}

	if ServerShutdownTimeoutStr, err := strconv.Atoi(os.Getenv("SERVER_SHUTDOWN_TIMEOUT")); err != nil {
		ServerShutdownTimeout = time.Duration(0)
	} else {
		ServerShutdownTimeout = time.Duration(ServerShutdownTimeoutStr)
	}

	server := &http.Server{
		Addr:              ":" + os.Getenv("SERVER_PORT"),
		ReadTimeout:       readTimeout * time.Second,       // Максимальное время ожидания чтения запроса
		WriteTimeout:      WriteTimeout * time.Second,      // Максимальное время ожидания записи ответа
		IdleTimeout:       IdleTimeout * time.Second,       // Таймаут для ожидания простоя соединений
		ReadHeaderTimeout: ReadHeaderTimeout * time.Second, // Время ожидания получения заголовков
	}

	if ServerShutdownTimeout != 0 {
		// Автоматический таймер для graceful shutdown
		go func() {
			time.AfterFunc(ServerShutdownTimeout*time.Second, func() {
				log.Println("Таймер истек — выполняем graceful shutdown...")

				if err := server.Shutdown(context.Background()); err != nil {
					log.Fatalf("Ошибка при остановке сервера по таймеру: %v", err)
				}

				close(idleConnsClosed)
			})
		}()
	}

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)

		<-sigint

		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatalf("Ошибка при остановке сервера: %v", err)
		}

		close(idleConnsClosed)
	}()

	routes.SetupRoutes()

	log.Println("Сервер запущен на порту 8080")
	log.Println("http://localhost:8080")

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("ошибка запуска сервера: %v", err)
	}

	<-idleConnsClosed
	log.Println("Сервер остановлен")
}
