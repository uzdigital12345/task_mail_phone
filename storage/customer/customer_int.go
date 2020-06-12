package customer

import (
	_ "github.com/lib/pq"
)
type Customer struct{
	id string
	phone string
	email string
	isSent bool
}
type CustomerStorageI interface {
	Create(customer *Customer) (*Customer, error)
	GetCustomer(id string) (*Customer, error)
	GetCustomers() ([]*Customer, error)
}
