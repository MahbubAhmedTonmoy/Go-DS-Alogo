package balancedparentheses

import "testing"

func TestIsBalancedSimple(t *testing.T) {
	input := "{[()]}"

	got := BalanceChecker(input)
	want := true

	if got != want {
		t.Errorf("\nInput: %s\nGot: %v\nWant: %v\n", input, got, want)
	}
}

func TestIsBalancedFalty(t *testing.T) {
	input := "{([()]}"

	got := BalanceChecker(input)
	want := false

	if got != want {
		t.Errorf("\nInput: %s\nGot: %v\nWant: %v\n", input, got, want)
	}
}

func TestIsBalancedHandlesEmpty(t *testing.T) {
	input := ""

	got := BalanceChecker(input)
	want := true

	if got != want {
		t.Errorf("\nInput: %s\nGot: %v\nWant: %v\n", input, got, want)
	}
}

func TestIsBalancedHandlesOneChar(t *testing.T) {
	input := "{"

	got := BalanceChecker(input)
	want := false

	if got != want {
		t.Errorf("\nInput: %s\nGot: %v\nWant: %v\n", input, got, want)
	}
}

func TestIsBalancedHandlesNonBracketsCorrectly(t *testing.T) {
	input := "aaaa"

	got := BalanceChecker(input)
	want := false

	if got != want {
		t.Errorf("\nInput: %s\nGot: %v\nWant: %v\n", input, got, want)
	}
}

func TestIsBalancedHandlesOrdering(t *testing.T) {
	input := "([)]"

	got := BalanceChecker(input)
	want := false

	if got != want {
		t.Errorf("\nInput: %s\nGot: %v\nWant: %v\n", input, got, want)
	}
}

func BenchmarkIsBalanced(b *testing.B) {
	input := "{[()]}"

	for i := 0; i < b.N; i++ {
		BalanceChecker(input)
	}
}
