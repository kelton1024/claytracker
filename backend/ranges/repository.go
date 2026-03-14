package ranges

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/jackc/pgx/v5"
)

type RangeRepository struct {
	conn    *pgx.Conn
	sqlData map[string]string
}

func NewRangeRepository(db *pgx.Conn) (*RangeRepository, error) {
	repo := &RangeRepository{conn: db}
	err := repo.loadSQLFiles("ranges/sql")
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (rr *RangeRepository) AddRange(ctx context.Context, rangeData *rangeAddRequest) error {
	tx, err := rr.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	sql := rr.sqlData["AddRange"]
	_, err = tx.Exec(ctx, sql, rangeData.Name, rangeData.Address1, rangeData.Address2, rangeData.City, rangeData.State, rangeData.Zipcode, rangeData.Lat, rangeData.Long)
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

// TODO: Make this a helper function that all APIs can use
func (rr *RangeRepository) loadSQLFiles(sqlBaseDir string) error {
	files, err := os.ReadDir(sqlBaseDir)
	if err != nil {
		return err
	}

	rr.sqlData = make(map[string]string)
	for _, file := range files {
		sql, err := os.ReadFile(path.Join(sqlBaseDir, file.Name()))
		if err != nil {
			return err
		}
		rr.sqlData["AddRange"] = string(sql)
	}
	return nil
}
