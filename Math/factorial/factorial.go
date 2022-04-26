package factorial

func IterativeFactorial(n int) int {
	tempN := n
	if n < 0 {
		tempN *= -1 // make number positive
	}
	result := 1
	for i := 2; i <= tempN; i++ {
		result *= i
	}
	if n < 0 {
		result *= -1 //// make number as it comes
	}
	return result
}

func RecursiveFactorial(n int) int {
	if n <= 1 {
		return 1
	} else {
		return n * RecursiveFactorial(n-1)
	}
}
