package main

import (
	"fmt"
	"math/rand"
	"time"

	"./goCuckooHashing"
)

func main() {
	n := 10
	c := goCuckooHashing.NewCuckoo()

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
	fmt.Println(c)

}
