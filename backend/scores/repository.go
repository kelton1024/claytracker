package scores

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/jackc/pgx/v5"
)

type ScoreRepository struct {
	conn    *pgx.Conn
	sqlData map[string]string
}

func NewScoreRepository(db *pgx.Conn) (*ScoreRepository, error) {
	repo := &ScoreRepository{conn: db}
	err := repo.loadSQLFiles("scores/sql")
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (sr *ScoreRepository) AddScore(ctx context.Context, scoreData *scoreAddRequest) error {
	tx, err := sr.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	fmt.Println(scoreData)
	sql := `INSERT INTO scores_tracking (score, station_number) VALUES ($1, $2);`
	_, err = tx.Exec(ctx, sql, scoreData.Scores, scoreData.Station)
	if err != nil {
		rbErr := tx.Rollback(ctx)
		if rbErr != nil {
			err = fmt.Errorf("insert failed with error %v and rollback failed with error %v", err, rbErr)
		}
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (sr *ScoreRepository) QueryScore(ctx context.Context, key string) ([]byte, error) {
	return nil, nil
}

// TODO: Make this a helper function that all APIs can use
func (sr *ScoreRepository) loadSQLFiles(sqlBaseDir string) error {
	files, err := os.ReadDir(sqlBaseDir)
	if err != nil {
		return err
	}

	sr.sqlData = make(map[string]string)
	for _, file := range files {
		sql, err := os.ReadFile(path.Join(sqlBaseDir, file.Name()))
		if err != nil {
			return err
		}
		sr.sqlData["AddRange"] = string(sql)
	}
	return nil
}
