package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/xvbnm48/go-grpc-rocket/internal/rocket"
)

type Store struct {
	db *sqlx.DB
}

// New: return a new store or error
func New() (Store, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbTable := os.Getenv("DB_TABLE")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUsername, dbPassword, dbTable, dbSSLMode)

	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		return Store{}, err
	}

	return Store{
		db: db,
	}, nil
}

func (s Store) GetRocketByID(id string) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

func (s Store) InsertRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

func (s Store) DeleteRocket(id string) error {
	return nil
}
