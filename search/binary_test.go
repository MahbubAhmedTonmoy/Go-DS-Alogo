package search

import "testing"

func TestBinary(t *testing.T) {
	for _, test := range searchTests {
		actualValue, actualError := Binary(test.data, test.target, 0, len(test.data)-1)
		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.target, test.expected, actualValue)
		}
		if actualError != test.expectedError {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected error '%s', get error '%s'", test.name, test.data, test.target, test.expectedError, actualError)
		}
	}
}
func TestBinaryIterative(t *testing.T) {
	for _, test := range searchTests {
		actualValue, actualError := BinaryIterative(test.data, test.target)
		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.target, test.expected, actualValue)
		}
		if actualError != test.expectedError {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected error '%s', get error '%s'", test.name, test.data, test.target, test.expectedError, actualError)
		}
	}
}
func BenchmarkBinary(b *testing.B) {
	testCase := generateBenchmarkTestCase()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Binary(testCase, i, 0, len(testCase)-1)
	}
}
func BenchmarkBinaryIterative(b *testing.B) {
	testCase := generateBenchmarkTestCase()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = BinaryIterative(testCase, i)
	}
}
