package users

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/jackc/pgx/v5"
	"github.com/alexedwards/argon2id"
	"runtime"
)

var encParams = &argon2id.Params {
	Memory:      64 * 1024,
	Iterations:  1,
	Parallelism: uint8(runtime.NumCPU()),
	SaltLength:  16,
	KeyLength:   32,
}

type UserRepository struct {
	conn    *pgx.Conn
	sqlData map[string]string
}

func NewUserRepository(db *pgx.Conn) (*UserRepository, error) {
	repo := &UserRepository{conn: db}
	err := repo.loadSQLFiles("users/sql")
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (ur *UserRepository) AddUser(ctx context.Context, userData *userAddRequest) error {
	tx, err := ur.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}


	sql := ur.sqlData["AddUser"]
	// Hash the incoming password
	passHash,err := argon2id.CreateHash(userData.Password, encParams)
	if err != nil {
		return err
	}
	//TODO: from here to carrots can be removed after verification that its working. This is also how to verify hashes for login.
	match,err := argon2id.ComparePasswordAndHash(userData.Password,passHash)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("Passowrd verificatin error")
	}
	// ^^^^^^^^^^
	_, err = tx.Exec(ctx, sql, userData.Username, userData.FirstName, userData.LastName, userData.Email, userData.Address1, userData.Address2, userData.City, userData.State, userData.Zipcode, passHash)
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
func (ur *UserRepository) loadSQLFiles(sqlBaseDir string) error {
	files, err := os.ReadDir(sqlBaseDir)
	if err != nil {
		return err
	}

	ur.sqlData = make(map[string]string)
	for _, file := range files {
		sql, err := os.ReadFile(path.Join(sqlBaseDir, file.Name()))
		if err != nil {
			return err
		}
		ur.sqlData["AddUser"] = string(sql)
	}
	return nil
}
