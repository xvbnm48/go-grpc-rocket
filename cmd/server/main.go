package main

import (
	"log"

	"github.com/xvbnm48/go-grpc-rocket/internal/db"
	"github.com/xvbnm48/go-grpc-rocket/internal/rocket"
)

func Run() error {
	// for initializing and starting
	// the server grpc
	rocketStore, err := db.New()
	if err != nil {
		return err
	}

	err = rocketStore.Migrate()
	if err != nil {
		log.Println("failed to migrations")
		return err
	}

	_ = rocket.New(rocketStore)
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
