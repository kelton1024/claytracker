package scores

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type ScoreService struct {
	repo *ScoreRepository
}

func NewScoreService(db *pgx.Conn) (*ScoreService, error) {
	repo, err := NewScoreRepository(db)
	if err != nil {
		return nil, err
	}

	return &ScoreService{repo: repo}, nil
}

func (ss *ScoreService) AddScore(ctx context.Context, scoreData *scoreAddRequest) error {
	fmt.Println(scoreData)
	return ss.repo.AddScore(ctx, scoreData)
}

func (ss *ScoreService) QueryScore(ctx context.Context, key string) ([]byte, error) {
	return nil, nil
}
