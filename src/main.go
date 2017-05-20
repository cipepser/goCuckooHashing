package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// c := goCuckooHashing.NewCuckoo()
	//
	// // insert the keys.
	// x := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// for _, key := range x {
	// 	cnt := 0
	// 	c.Insert(key, cnt)
	// }
	//
	// // look up for the key "1" and "10".
	// fmt.Println("key:1  ", c.Lookup(1))  // key:1   true
	// fmt.Println("key:10 ", c.Lookup(10)) // key:10  false
	//
	// // delete the key "3".
	// fmt.Println("before: ", *c) // before:  {[0 8 6 0 4 0 9 0 5 7] [0 1 0 0 0 2 0 0 3 0]}
	// c.Delete(3)
	// fmt.Println("after:  ", *c) // after:   {[0 8 6 0 4 0 9 0 5 7] [0 1 0 0 0 2 0 0 0 0]}

	num := 10
	rand.Seed(time.Now().UnixNano())
	ns := make([]int64, num)
	for i := 0; i < num; i++ {
		ns[i] = rand.Int63()
	}
	fmt.Println(ns)

}
