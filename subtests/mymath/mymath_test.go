package mymath_test

import (
	"testing"

	"github.com/MahlerFive/go-unittest-talk/parallel/mymath"
)

func TestMultiply(t *testing.T) {
	t.Parallel()

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
			Name:            "multiply negatives",
			Op1:             -3,
			Op2:             -5,
			ExpectedProduct: 15,
		},
	}

	for _, tc := range testCases {
		// Capture 'tc' variable
		tc := tc

		// Run subtest
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			actualProduct := mymath.Multiply(tc.Op1, tc.Op2)
			if actualProduct != tc.ExpectedProduct {
				t.Errorf("%s: %d * %d = %d, expected %d", tc.Name, tc.Op1, tc.Op2, actualProduct, tc.ExpectedProduct)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name        string
		Op1         int
		Op2         int
		ExpectedSum int
	}

	testCases := []TestCase{
		{
			Name:        "add 0",
			Op1:         0,
			Op2:         5,
			ExpectedSum: 5,
		},
		{
			Name:        "add positives",
			Op1:         3,
			Op2:         5,
			ExpectedSum: 8,
		},
		{
			Name:        "add negatives",
			Op1:         -3,
			Op2:         -5,
			ExpectedSum: -8,
		},
	}

	for _, tc := range testCases {
		// Capture 'tc' variable
		tc := tc

		// Run subtest
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			actualSum := mymath.Add(tc.Op1, tc.Op2)
			if actualSum != tc.ExpectedSum {
				t.Errorf("%s: %d + %d = %d, expected %d", tc.Name, tc.Op1, tc.Op2, actualSum, tc.ExpectedSum)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name               string
		Op1                int
		Op2                int
		ExpectedDifference int
	}

	testCases := []TestCase{
		{
			Name:               "subtract 0",
			Op1:                5,
			Op2:                0,
			ExpectedDifference: 5,
		},
		{
			Name:               "subtract positives",
			Op1:                5,
			Op2:                3,
			ExpectedDifference: 2,
		},
		{
			Name:               "subtract negatives",
			Op1:                -3,
			Op2:                -5,
			ExpectedDifference: 2,
		},
	}

	for _, tc := range testCases {
		// Capture 'tc' variable
		tc := tc

		// Run subtest
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			actualDifference := mymath.Subtract(tc.Op1, tc.Op2)
			if actualDifference != tc.ExpectedDifference {
				t.Errorf("%s: %d + %d = %d, expected %d", tc.Name, tc.Op1, tc.Op2, actualDifference, tc.ExpectedDifference)
			}
		})
	}
}
