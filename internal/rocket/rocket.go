//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/xvbnm48/go-grpc-rocket/internal/rocket Store

package rocket

import "context"

// Rocket : Rocket is a struct that contains
type Rocket struct {
	ID     string
	Name   string
	Type   string
	Flight int
}

// Store: define a method for store a new rocket
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// Service: our rocket service for updating rocket inventory
type Service struct {
	Store Store
}

// New: for return a new instance of our rockets service
func New(store Store) Service {
	return Service{
		Store: store,
	}
}

// GetRocketByID: retrieve a rocket by its ID
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}

	return rkt, nil
}

// InsertRocket: insert a new rocket to the store
func (s Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}

	return rkt, nil
}

// DeleteRocket: delete rocket from inventory
func (s Service) DeleteRocket(id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}

	return nil
}
