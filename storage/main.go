package storage

import (
	"github.com/jmoiron/sqlx"

	"github.com/uzdigital12345/task_mail_phone/storage/customer"
	er "github.com/uzdigital12345/task_mail_phone/storage/error"
	"github.com/uzdigital12345/task_mail_phone/storage/purchase"
)

type StorageI interface {
	Customer() customer.CustomerStorageI
	Purchase() purchase.PurchaseStorageI
	Error() er.ErrorStorageI
}

type StoragePg struct {
	customer customer.CustomerStorageI
	purchase purchase.PurchaseStorageI
	er       er.ErrorStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &StoragePg{
		customer: customer.NewSqlx(db),
		purchase: purchase.NewSqlx(db),
		er:       er.NewSqlx(db),
	}
}

func (s StoragePg) Customer() customer.CustomerStorageI {
	return s.customer
}

func (s StoragePg) Purchase() purchase.PurchaseStorageI {
	return s.purchase
}

func (s StoragePg) Error() er.ErrorStorageI {
	return s.er
}
