package main

import "testing"

// go test -coverprofile=c.out
//go tool cover -html c.ouy
func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{30000, 40000, 50000},
	}

	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d, %d); got %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}

//go test -bench . -cpuprofile cpu.out
//go tool pprof cpu.out
func BenchmarkTriangle(b *testing.B) {
	x, y := 3, 4
	ans := 5

	for i := 0; i < b.N; i++ {
		actual := calcTriangle(x, y)
		if actual != ans {
			b.Errorf("calcTriangle(%d, %d); got %d; expected %d", x, y, ans, actual)
		}
	}
}
