package mymath_test

import (
	"testing"

	"github.com/MahlerFive/go-unittest-talk/tabledriven/mymath"
)

func TestMultiply(t *testing.T) {
	type TestCase struct {
		Name            string
		Op1             int
		Op2             int
		ExpectedProduct int
	}

	testCases := []TestCase{
		{
			Name:            "multiply by 0",
			Op1:             0,
			Op2:             5,
			ExpectedProduct: 0,
		},
		{
			Name:            "multiply positives",
			Op1:             3,
			Op2:             5,
			ExpectedProduct: 15,
		},
		{
			Name:            "multiply negative by positive",
			Op1:             -3,
			Op2:             5,
			ExpectedProduct: -15,
		},
		{
			Name:            "multiply positive by negative",
			Op1:             3,
			Op2:             -5,
			ExpectedProduct: -15,
		},
		{
			Name:            "multiply negatives",
			Op1:             -3,
			Op2:             -5,
			ExpectedProduct: 15,
		},
	}

	for _, tc := range testCases {
		actualProduct := mymath.Multiply(tc.Op1, tc.Op2)
		if actualProduct != tc.ExpectedProduct {
			t.Errorf("%s: %d * %d = %d, expected %d", tc.Name, tc.Op1, tc.Op2, actualProduct, tc.ExpectedProduct)
		}
	}

}
