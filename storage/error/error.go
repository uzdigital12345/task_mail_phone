package error

import (
	"github.com/jmoiron/sqlx"
)

type Sqlx struct {
	db *sqlx.DB
}

// NewSqlx ...
func NewSqlx(db *sqlx.DB) ErrorStorageI {
	return &Sqlx{
		db: db,
	}
}

func (cm *Sqlx) SaveError(Error string)  error {
	tx, err := cm.db.Begin()

	if err != nil {
		return err
	}

	insertNew :=
		`INSERT INTO
		errors
		(
			error
		)
		VALUES
		($1)`
	_, err = tx.Exec(
		insertNew,
		Error,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
