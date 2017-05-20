package goCuckooHashing

import (
	"math/rand"
	"testing"
	"time"
)

func benchmarkCuckooHashingInsert(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := NewCuckoo()

		// generate pseudo-nubers
		rand.Seed(time.Now().UnixNano())
		keys := make([]int64, n)
		for i := 0; i < n; i++ {
			keys[i] = rand.Int63()
		}

		// insert the keys.
		for _, key := range keys {
			cnt := 0
			c.Insert(key, cnt)
		}
	}

}

func BenchmarkCuckooHashingInsert10(b *testing.B)   { benchmarkCuckooHashingInsert(10, b) }
func BenchmarkCuckooHashingInsert30(b *testing.B)   { benchmarkCuckooHashingInsert(30, b) }
func BenchmarkCuckooHashingInsert50(b *testing.B)   { benchmarkCuckooHashingInsert(50, b) }
func BenchmarkCuckooHashingInsert70(b *testing.B)   { benchmarkCuckooHashingInsert(70, b) }
func BenchmarkCuckooHashingInsert100(b *testing.B)  { benchmarkCuckooHashingInsert(100, b) }
func BenchmarkCuckooHashingInsert300(b *testing.B)  { benchmarkCuckooHashingInsert(300, b) }
func BenchmarkCuckooHashingInsert500(b *testing.B)  { benchmarkCuckooHashingInsert(500, b) }
func BenchmarkCuckooHashingInsert700(b *testing.B)  { benchmarkCuckooHashingInsert(700, b) }
func BenchmarkCuckooHashingInsert1000(b *testing.B) { benchmarkCuckooHashingInsert(1000, b) }
func BenchmarkCuckooHashingInsert3000(b *testing.B) { benchmarkCuckooHashingInsert(3000, b) }
