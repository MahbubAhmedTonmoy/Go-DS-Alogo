package fibonacciseries

import "math"

func RecursiveFibonacci(nth uint) uint {
	if nth < 1 {
		return nth
	}
	return RecursiveFibonacci(nth-1) + RecursiveFibonacci(nth-2)
}
func BrutforceFibonacci(n uint) uint {
	a := uint(0)
	b := uint(1)
	result := uint(0)
	if n < 2 {
		result = +n
		a, b = b, result
	}
	for i := uint(2); i <= n; i++ {
		result = a + b
		a, b = b, result
	}
	return result
}

// Formula This function calculates the n-th fibonacci number using the [formula](https://en.wikipedia.org/wiki/Fibonacci_number#Relation_to_the_golden_ratio)
// Attention! Tests for large values fall due to rounding error of floating point numbers, works well, only on small numbers
func Formula(n uint) uint {
	sqrt5 := math.Sqrt(5)
	phi := (sqrt5 + 1) / 2
	powPhi := math.Pow(phi, float64(n))
	return uint(powPhi/sqrt5 + 0.5)
}
