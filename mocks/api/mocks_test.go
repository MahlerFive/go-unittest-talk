package api_test

import (
	"github.com/MahlerFive/go-unittest-talk/mocks/api"
)

type MockCustomerStore struct {
	// Have a public member for each function we want to mock
	// The tests can then easily set any of these members as necessary
	MockCreate func(c api.Customer) error
	MockGet    func(id int) (*api.Customer, error)
}

func NewMockCustomerStore() *MockCustomerStore {
	return &MockCustomerStore{}
}

func (s *MockCustomerStore) Create(c api.Customer) error {
	// Call the overridden member
	return s.MockCreate(c)
}

func (s *MockCustomerStore) Get(id int) (*api.Customer, error) {
	// Call teh overridden member
	return s.MockGet(id)
}
