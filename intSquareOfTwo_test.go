package main

//	go test -v
//	go test -bench=.

import (
	"fmt"
	"testing"
)

func TestSliceOfBytes(t *testing.T) {
	ans := SliceOfBytes(2)
	if !ans {
		t.Errorf("IntMin(2, -2) = %v; want -2", ans)
	}
}

func TestSliceOfBytesTableDriven(t *testing.T) {
	var tests = []struct {
		a    int
		want bool
	}{
		{1, false},
		{2, true},
		{-256, true},
		{1024, true},
		{1023, false},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d", tt.a)
		t.Run(testName, func(t *testing.T) {
			ans := SliceOfBytes(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func BenchmarkSliceOfBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceOfBytes(i)
	}
}
