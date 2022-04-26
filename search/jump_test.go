package search

import "testing"

func TestJumpSearch(t *testing.T) {
	for _, test := range searchTests {
		actualValue, actualError := JumpSearch(test.data, test.target, len(test.data)-1)
		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.target, test.expected, actualValue)
		}
		if actualError != test.expectedError {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected error '%s', get error '%s'", test.name, test.data, test.target, test.expectedError, actualError)
		}
	}
}

func BenchmarkJumpSearch(b *testing.B) {
	testCase := generateBenchmarkTestCase()
	b.ResetTimer() // exclude time taken to generate test case
	for i := 0; i < b.N; i++ {
		_, _ = LinearSearch(testCase, i)
	}
}
