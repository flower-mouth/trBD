package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"time"
	"trBD/internal/configuration"
)

type Client interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Close(context.Context) error
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

func NewClient(ctx context.Context, maxAttempts int, sc configuration.StConfig) (conn *pgx.Conn) {
	var err error
	connectionUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)
	err = AttemptDatabaseConnection(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		conn, err = pgx.Connect(ctx, connectionUrl)
		if err != nil {
			return err
		}

		return nil
	}, maxAttempts, 5*time.Second)

	if err != nil {
		log.Printf("Error connecting to database")
	}

	return conn
}

var DBClient Client

func init() {
	// Создаем клиент для работы с базой данных
	sc := configuration.StConfig{
		Username: "postgres",
		Password: "password",
		Host:     "localhost",
		Port:     "5432",
		Database: "trbd",
	}

	// Инициализируем клиент базы данных
	DBClient = NewClient(context.Background(), 3, sc)

	// Проверяем подключение к базе данных
	err := AttemptDatabaseConnection(func() error {
		conn, err := DBClient.Begin(context.Background())
		if err != nil {
			return err
		}
		defer conn.Rollback(context.Background())

		// Проверяем, что таблица Participants существует
		_, err = conn.Exec(context.Background(), "SELECT 1 FROM Participants LIMIT 1")
		if err != nil {
			return err
		}

		return nil
	}, 3, 5*time.Second)

	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
}

func AttemptDatabaseConnection(fn func() error, attempts int, delay time.Duration) (err error) {
	for attempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attempts--

			continue
		}
		return nil
	}
	return
}
