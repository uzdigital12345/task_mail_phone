package purchase

import (
	_ "github.com/lib/pq"
)
type Purchase struct{
	id string
	product string
	price float64
}
type PurchaseStorageI interface {
	Create(purchase *Purchase) (*Purchase, error)
	GetPurchase(id string) (*Purchase, error)
	GetPurchases() ([]*Purchase, error)
}
