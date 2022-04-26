package factorial

import "testing"

func testCases() []struct {
	name   string
	number int
	result int
} {
	var tests = []struct {
		name   string
		number int
		result int
	}{
		{"5!", 5, 120},
		{"3!", 3, 6},
		{"6!", 6, 720},
		{"12!", 12, 479001600},
		{"1!", 1, 1},
		{"0!", 0, 1},
		// {"-1!", -1, -1},
		// {"-12!", -12, -479001600},
	}
	return tests
}
func TestBruteForceFactorial(t *testing.T) {
	tests := testCases()
	for _, tv := range tests {
		t.Run(tv.name, func(t *testing.T) {
			result := IterativeFactorial(tv.number)
			if result != tv.result {
				t.Errorf("Wrong result! Expected:%d, returned:%d ", tv.result, result)
			}
		})
	}
}

func TestRecursiveFactorial(t *testing.T) {
	tests := testCases()
	for _, tv := range tests {
		t.Run(tv.name, func(t *testing.T) {
			result := RecursiveFactorial(tv.number)
			if result != tv.result {
				t.Errorf("Wrong result! Expected:%d, returned:%d ", tv.result, result)
			}
		})
	}
}

func BenchmarkBruteForceFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IterativeFactorial(10)
	}
}

func BenchmarkRecursiveFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecursiveFactorial(10)
	}
}
