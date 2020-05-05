package unittest

import (
	"math"
	"testing"
)

func TestAdd32(t *testing.T) {
	tests := []struct{ a, b, c int32 }{
		{1, 2, 3},
		{0, 0, 0},
		{math.MaxInt32, 1, math.MinInt32},
	}
	for _, test := range tests {
		if actual := Add32(test.a, test.b); actual != test.c {
			t.Errorf("add() error on a=%d and b=%d, expecting %d, actual got %d",
				test.a, test.b, test.c, actual)
		}
	}
}

func BenchmarkAdd32(b *testing.B) {
	var a1, a2 int32 = 1, 2
	var res int32 = 3
	for i := 0; i < b.N; i++ {
		if actual := Add32(a1, a2); actual != res {
			b.Errorf("add() error on a=%d and b=%d, expecting %d, actual got %d",
				a1, a2, res, actual)
		}
	}
}
