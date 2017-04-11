package api

import (
	"github.com/pkg/errors"
)

type CustomerStore interface {
	Create(c Customer) error
	Get(id int) (*Customer, error)
}

type CustomerAPI struct {
	store CustomerStore
}

type Customer struct {
	ID   int
	Name string
	Age  int
}

// Take the store as an interface so we can provide a mock when testing
func NewCustomerAPI(s CustomerStore) *CustomerAPI {
	return &CustomerAPI{
		store: s,
	}
}

func (api *CustomerAPI) CreateCustomer(customer Customer) error {
	// Verify age
	if customer.Age < 0 {
		return errors.New("can not create customer with negative age")
	}

	// Create customer in the store
	if err := api.store.Create(customer); err != nil {
		return errors.Wrapf(err, "error creating customer %+v", customer)
	}

	return nil
}

func (api *CustomerAPI) GetCustomer(id int) (*Customer, error) {
	// Get customer from the store
	customer, err := api.store.Get(id)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting customer %d", id)
	}

	return customer, nil
}
