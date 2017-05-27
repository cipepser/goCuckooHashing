package goCuckooHashing

import (
	"math/rand"
	"testing"
	"time"
)

func benchmarkCuckooHashingInsert(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		c := NewCuckoo()

		// generate pseudo-nubers
		rand.Seed(time.Now().UnixNano())
		keys := make([]int64, n)
		for j := 0; j < n; j++ {
			keys[j] = rand.Int63()
		}

		b.StartTimer()
		// insert the keys.
		for _, key := range keys {
			cnt := 0
			c.Insert(key, cnt)
		}
	}

}

func BenchmarkCuckooHashingInsert10000(b *testing.B)  { benchmarkCuckooHashingInsert(10000, b) }
func BenchmarkCuckooHashingInsert20000(b *testing.B)  { benchmarkCuckooHashingInsert(20000, b) }
func BenchmarkCuckooHashingInsert30000(b *testing.B)  { benchmarkCuckooHashingInsert(30000, b) }
func BenchmarkCuckooHashingInsert40000(b *testing.B)  { benchmarkCuckooHashingInsert(40000, b) }
func BenchmarkCuckooHashingInsert50000(b *testing.B)  { benchmarkCuckooHashingInsert(50000, b) }
func BenchmarkCuckooHashingInsert60000(b *testing.B)  { benchmarkCuckooHashingInsert(60000, b) }
func BenchmarkCuckooHashingInsert70000(b *testing.B)  { benchmarkCuckooHashingInsert(70000, b) }
func BenchmarkCuckooHashingInsert80000(b *testing.B)  { benchmarkCuckooHashingInsert(80000, b) }
func BenchmarkCuckooHashingInsert90000(b *testing.B)  { benchmarkCuckooHashingInsert(90000, b) }
func BenchmarkCuckooHashingInsert100000(b *testing.B) { benchmarkCuckooHashingInsert(100000, b) }
