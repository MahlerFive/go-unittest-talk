package api_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/MahlerFive/go-unittest-talk/mocks/api"
)

func TestCreateCustomer(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		TestName      string
		Customer      api.Customer
		StoreError    bool
		ExpectedError bool
	}

	testCases := []TestCase{
		{
			TestName: "normal case",
			Customer: api.Customer{
				Name: "jeff",
				Age:  33,
			},
			StoreError:    false,
			ExpectedError: false,
		},
		{
			TestName: "store error",
			Customer: api.Customer{
				Name: "jeff",
				Age:  33,
			},
			StoreError:    true,
			ExpectedError: true,
		},
		{
			TestName: "negative age",
			Customer: api.Customer{
				Name: "jeff",
				Age:  -100,
			},
			StoreError:    false,
			ExpectedError: true,
		},
	}

	for _, tc := range testCases {
		// Capture 'tc' variable
		tc := tc

		// Run subtest
		t.Run(tc.TestName, func(t *testing.T) {
			t.Parallel()

			// Use a mock store to have Create() return what we want
			store := NewMockCustomerStore()
			store.MockCreate = func(c api.Customer) error {
				if tc.StoreError {
					return errors.New("store error")
				}
				return nil
			}

			customerAPI := api.NewCustomerAPI(store)
			err := customerAPI.CreateCustomer(tc.Customer)
			if err == nil && tc.ExpectedError {
				t.Error("Expected error, but got none")
			}
			if err != nil && !tc.ExpectedError {
				t.Errorf("Got error %s, but expected none", err)
			}
		})
	}
}

func TestGetCustomer(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		TestName         string
		ID               int
		StoreError       bool
		ExpectedCustomer *api.Customer
		ExpectedError    bool
	}

	testCases := []TestCase{
		{
			TestName:   "normal case",
			ID:         123,
			StoreError: false,
			ExpectedCustomer: &api.Customer{
				ID:   123,
				Name: "jeff",
				Age:  33,
			},
			ExpectedError: false,
		},
		{
			TestName:         "store error",
			ID:               123,
			StoreError:       true,
			ExpectedCustomer: nil,
			ExpectedError:    true,
		},
		{
			TestName:         "missing customer",
			ID:               789,
			StoreError:       false,
			ExpectedCustomer: nil,
			ExpectedError:    false,
		},
	}

	for _, tc := range testCases {
		// Capture 'tc' variable
		tc := tc

		// Run subtest
		t.Run(tc.TestName, func(t *testing.T) {
			t.Parallel()

			// Use a mock store to have Get() return what we want
			store := NewMockCustomerStore()
			store.MockGet = func(id int) (*api.Customer, error) {
				var err error
				if tc.StoreError {
					err = errors.New("store error")
				}
				return tc.ExpectedCustomer, err
			}

			customerAPI := api.NewCustomerAPI(store)
			customer, err := customerAPI.GetCustomer(tc.ID)
			if err == nil && tc.ExpectedError {
				t.Error("Expected error, but got none")
			}
			if err != nil && !tc.ExpectedError {
				t.Errorf("Got error %s, but expected none", err)
			}
			if !reflect.DeepEqual(customer, tc.ExpectedCustomer) {
				t.Errorf("expected customer %+v, actual customer %+v", tc.ExpectedCustomer, customer)
			}
		})
	}
}
