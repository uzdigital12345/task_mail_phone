package customer

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Sqlx struct {
	db *sqlx.DB
}

// NewSqlx ...
func NewSqlx(db *sqlx.DB) CustomerStorageI {
	return &Sqlx{
		db: db,
	}
}

func (cm *Sqlx) Create(customer *Customer) (*Customer, error) {
	tx, err := cm.db.Begin()

	if err != nil {
		return nil, err
	}

	id, err := uuid.NewRandom()
	insertNew :=
		`INSERT INTO
		customers
		(
			id,
			phone,
			email
		)
		VALUES
		($1, $2, $3)`
	_, err = tx.Exec(
		insertNew,
		id,
		customer.phone,
		customer.email,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	c, err := cm.GetCustomer(customer.id)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cm *Sqlx) GetCustomer(id string) (*Customer, error) {
	var (
		customer Customer
		column   string
	)
	_, err := uuid.Parse(id)
	if err != nil {
		column = " phone"
	} else {
		column = " id"
	}
	row := cm.db.QueryRow(`
		SELECT  id,
				phone,
				email,
				is_sent
		FROM customers
		WHERE `+column+`=$1 `, id,
	)

	err = row.Scan(
		&customer.id,
		&customer.phone,
		&customer.email,
		&customer.isSent,
	)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (cm *Sqlx) GetCustomers() ([]*Customer, error) {
	var customers []*Customer

	query := `
		SELECT  id,
				phone,
				email,
				is_sent
		FROM customers `
	rows, err := cm.db.Queryx(query)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var c Customer
		err = rows.Scan(
			&c.id,
			&c.phone,
			&c.email,
			&c.isSent,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, &c)
	}
	return customers, nil
}
