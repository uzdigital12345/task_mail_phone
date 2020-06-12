package purchase

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	PostgresHost     = "localhost"
	PostgresPort     = 5432
	PostgresDatabase = "alief_tech"
	PostgresUser     = "delever"
	PostgresPassword = "delever"
)

type Sqlx struct {
	db *sqlx.DB
}

// NewSqlx ...
func NewSqlx(db *sqlx.DB) PurchaseStorageI {

	return &Sqlx{
		db: db,
	}
}

func (cm *Sqlx) Create(purchase *Purchase) (*Purchase, error) {
	tx, err := cm.db.Begin()

	if err != nil {
		return nil, err
	}

	id, err := uuid.NewRandom()
	insertNew :=
		`INSERT INTO
		purchases
		(
			id,
			product,
			price
		)
		VALUES
		($1, $2, $3)`
	_, err = tx.Exec(
		insertNew,
		id,
		purchase.product,
		purchase.price,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	c, err := cm.GetPurchase(purchase.id)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cm *Sqlx) GetPurchase(id string) (*Purchase, error) {
	var (
		purchase Purchase
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
				product,
				price
		FROM purchases
		WHERE `+column+`=$1 `, id,
	)

	err = row.Scan(
		&purchase.id,
		&purchase.product,
		&purchase.price,
	)
	if err != nil {
		return nil, err
	}
	return &purchase, nil
}

func (cm *Sqlx) GetPurchases() ([]*Purchase, error) {
	var purchases []*Purchase

	query := `
		SELECT  id,
				product,
				price
		FROM purchases `
	rows, err := cm.db.Queryx(query)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var c Purchase
		err = rows.Scan(
			&c.id,
			&c.product,
			&c.price,
		)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, &c)
	}
	return purchases, nil
}
