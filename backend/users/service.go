package users

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type UserService struct {
	repo *UserRepository
}

func NewUserService(db *pgx.Conn) (*UserService, error) {
	repo, err := NewUserRepository(db)
	if err != nil {
		return nil, err
	}

	return &UserService{repo: repo}, nil
}

// Handle data validation here
func (svc *UserService) AddUser(ctx context.Context, userData *userAddRequest) error {
	err := svc.validate(userData)
	if err != nil {
		return err
	}

	log.Print("adding a new user to the database")
	err = svc.repo.AddUser(ctx, userData)
	if err != nil {
		return err
	}
	return nil
}

// TODO: Validate against JSON Schema
// Note: I think we should be able to take in
// any and just call the appropriate schema to validate
func (svc *UserService) validate(data any) error {
	return nil
}
