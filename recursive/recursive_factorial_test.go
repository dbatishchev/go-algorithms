package recursive

import (
	"testing"
)

var tests = []struct {
	in  int
	out int
}{
	{1, 1},
	{2, 2},
	{5, 120},
	{10, 3628800},
}

func TestFactorial(t *testing.T) {
	for _, tt := range tests {
		r := Factorial(tt.in)
		if r != tt.out {
			t.Errorf("Factorial(%d) => %d, want %d", tt.in, r, tt.out)
		}
	}
}