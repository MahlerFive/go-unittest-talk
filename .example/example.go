package mymath_test

import "testing"

func TestMyUnit(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		TestName       string
		Input          int
		ExpectedOutput int
	}

	testCases := []TestCase{
		{
			TestName:       "test normal case",
			Input:          1,
			ExpectedOutput: 3,
		},
		{
			TestName:       "test edge case",
			Input:          0,
			ExpectedOutput: 0,
		},
	}

	for _, tc := range testCases {
		actualOutput := mypackage.MyUnit(tc.Input)
		if actualOutput != tc.ExpectedOutput {
			t.Errorf("%s: actual %d, expected %d", tc.TestName, actualOutput, tc.ExpectedOutput)
		}
	}

}
