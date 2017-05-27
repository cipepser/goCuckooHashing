package goCuckooHashing

import (
	"math/rand"
	"testing"
	"time"
)

func benchmarkBucketizedCuckooHashingInsert(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		c := NewBucketizedCuckoo()

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

func BenchmarkBucketizedCuckooHashingInsert10000(b *testing.B) {
	benchmarkBucketizedCuckooHashingInsert(10000, b)
}
func BenchmarkBucketizedCuckooHashingInsert20000(b *testing.B) {
	benchmarkBucketizedCuckooHashingInsert(20000, b)
}
func BenchmarkBucketizedCuckooHashingInsert30000(b *testing.B) {
	benchmarkBucketizedCuckooHashingInsert(30000, b)
}
func BenchmarkBucketizedCuckooHashingInsert40000(b *testing.B) {
	benchmarkBucketizedCuckooHashingInsert(40000, b)
}
func BenchmarkBucketizedCuckooHashingInsert50000(b *testing.B) {
	benchmarkBucketizedCuckooHashingInsert(50000, b)
}
func BenchmarkBucketizedCuckooHashingInsert60000(b *testing.B) {
	benchmarkBucketizedCuckooHashingInsert(60000, b)
}
func BenchmarkBucketizedCuckooHashingInsert70000(b *testing.B) {
	benchmarkBucketizedCuckooHashingInsert(70000, b)
}
func BenchmarkBucketizedCuckooHashingInsert80000(b *testing.B) {
	benchmarkBucketizedCuckooHashingInsert(80000, b)
}
func BenchmarkBucketizedCuckooHashingInsert90000(b *testing.B) {
	benchmarkBucketizedCuckooHashingInsert(90000, b)
}
func BenchmarkBucketizedCuckooHashingInsert100000(b *testing.B) {
	benchmarkBucketizedCuckooHashingInsert(100000, b)
}
