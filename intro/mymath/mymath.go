package mymath

// Bad version
func Multiply(op1 int, op2 int) int {
	product := 0
	for i := 0; i < op2; i++ {
		product += op1
	}
	return product
}

// Good version
// func Multiply(op1 int, op2 int) int {
// 	return op1 * op2
// }
