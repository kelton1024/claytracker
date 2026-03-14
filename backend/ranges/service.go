package ranges

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type RangeService struct {
	repo *RangeRepository
}

func NewRangeService(db *pgx.Conn) (*RangeService, error) {
	repo, err := NewRangeRepository(db)
	if err != nil {
		return nil, err
	}

	return &RangeService{repo: repo}, nil
}

// Handle data validation here
func (svc *RangeService) AddRange(ctx context.Context, rangeData *rangeAddRequest) error {
	err := svc.validate(rangeData)
	if err != nil {
		return err
	}

	log.Print("adding a new range to the database")
	err = svc.repo.AddRange(ctx, rangeData)
	if err != nil {
		return err
	}
	return nil
}

// TODO: Validate against JSON Schema
// Note: I think we should be able to take in
// any and just call the appropriate schema to validate
func (svc *RangeService) validate(data any) error {
	return nil
}
