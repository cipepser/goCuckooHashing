package goCuckooHashing

import (
	"testing"
)

func BenchmarkCuckooHashing(b *testing.B) {
	for i:= 0; i <b.N; i++ {
		CuckooHashing()
	}
}