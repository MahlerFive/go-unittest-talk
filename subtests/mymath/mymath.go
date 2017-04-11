package mymath

import "time"

func Multiply(op1 int, op2 int) int {
	// simulate a slow function
	time.Sleep(time.Second * 1)

	return op1 * op2
}

func Add(op1 int, op2 int) int {
	// simulate a slow function
	time.Sleep(time.Second * 1)

	return op1 + op2
}

func Subtract(op1 int, op2 int) int {
	// simulate a slow function
	time.Sleep(time.Second * 1)

	return op1 - op2
}
