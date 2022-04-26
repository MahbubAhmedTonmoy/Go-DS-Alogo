package fibonacciseries

import "testing"

func testCases() []struct {
	name     string
	n        uint
	expected uint
} {
	tests := []struct {
		name     string
		n        uint
		expected uint
	}{
		{"Fibonacci 0-th number == 0", 0, 0},
		{"Fibonacci 1-th number == 1", 1, 1},
		{"Fibonacci 2-th number == 1", 2, 1},
		{"Fibonacci 3-th number == 2", 3, 2},
		{"Fibonacci 4-th number == 3", 4, 3},
		{"Fibonacci 5-th number == 5", 5, 5},
		{"Fibonacci 6-th number == 8", 6, 8},
		{"Fibonacci 7-th number == 13", 7, 13},
		{"Fibonacci 8-th number == 21", 8, 21},
		{"Fibonacci 9-th number == 34", 9, 34},
		{"Fibonacci 10-th number == 55", 10, 55},
		{"Fibonacci 90-th number == 2880067194370816120", 90, 2880067194370816120},
	}
	return tests
}

func TestRecursiveFibonacci(t *testing.T) {
	tests := testCases()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := RecursiveFibonacci(test.n)
			if actual != test.expected {
				t.Errorf("Return value = %v, want %v", actual, test.expected)
			}
		})
	}
}

func TestBrutforceFibonacci(t *testing.T) {
	tests := testCases()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := BrutforceFibonacci(test.n)
			if actual != test.expected {
				t.Errorf("Return value = %v, want %v", actual, test.expected)
			}
		})
	}
}

func TestFormula(t *testing.T) {
	tests := testCases()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Formula(test.n)
			if actual != test.expected {
				t.Errorf("Return value = %v, want %v", actual, test.expected)
			}
		})
	}
}
