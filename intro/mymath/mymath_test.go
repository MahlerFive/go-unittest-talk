package mymath_test

import (
	"testing"

	"github.com/MahlerFive/go-unittest-talk/intro/mymath"
)

func TestMultiply(t *testing.T) {
	op1 := 0
	op2 := 5
	expectedProduct := 0
	actualProduct := mymath.Multiply(op1, op2)
	if actualProduct != expectedProduct {
		t.Errorf("%d * %d = %d, expected %d", op1, op2, actualProduct, expectedProduct)
	}

	op1 = 3
	op2 = 5
	expectedProduct = 15
	actualProduct = mymath.Multiply(op1, op2)
	if actualProduct != expectedProduct {
		t.Errorf("%d * %d = %d, expected %d", op1, op2, actualProduct, expectedProduct)
	}

	op1 = -3
	op2 = 5
	expectedProduct = -15
	actualProduct = mymath.Multiply(op1, op2)
	if actualProduct != expectedProduct {
		t.Errorf("%d * %d = %d, expected %d", op1, op2, actualProduct, expectedProduct)
	}

	op1 = 3
	op2 = -5
	expectedProduct = -15
	actualProduct = mymath.Multiply(op1, op2)
	if actualProduct != expectedProduct {
		t.Errorf("%d * %d = %d, expected %d", op1, op2, actualProduct, expectedProduct)
	}

	op1 = -3
	op2 = -5
	expectedProduct = 15
	actualProduct = mymath.Multiply(op1, op2)
	if actualProduct != expectedProduct {
		t.Errorf("%d * %d = %d, expected %d", op1, op2, actualProduct, expectedProduct)
	}
}
