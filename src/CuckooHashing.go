package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/big"
)

const (
	N       int64 = 10 // テーブルの大きさ
	MAXLOOP int   = 100
)

func hash(key int64) (h1, h2 int64) {
	hasher := md5.New()
	hasher.Write([]byte(string(key)))

	h := hex.EncodeToString(hasher.Sum(nil))

	h1_t, _ := new(big.Int).SetString(h[:int(len(h)/2)], 16)
	h2_t, _ := new(big.Int).SetString(h[int(len(h)/2):], 16)
	h1 = h1_t.Rem(h1_t, big.NewInt(N)).Int64()
	h2 = h2_t.Rem(h2_t, big.NewInt(N)).Int64()

	return h1, h2
}

type Cuckoo struct {
	T1, T2 [N]int64
}

func NewCuckoo() *Cuckoo {
	return new(Cuckoo)
}

func (c *Cuckoo) lookup(key int64) bool {
	h1, h2 := hash(key)

	if c.T1[h1] == key || c.T2[h2] == key {
		return true
	} else {
		return false
	}
}

func (c *Cuckoo) insert(key int64, cnt int) {
	for cnt < MAXLOOP {
		h1, h2 := hash(key)
		switch cnt % 2 {
		case 0:
			if c.T1[h1] == 0 {
				c.T1[h1] = key
				return
			} else {
				key, c.T1[h1] = c.T1[h1], key
			}
		case 1:
			if c.T2[h2] == 0 {
				c.T2[h2] = key
				return
			} else {
				key, c.T2[h2] = c.T2[h2], key
			}
		}
		cnt++
	}
	panic("fail to insert. you must reconstruct the Table...")
}

func (c *Cuckoo) delete(key int64) {
	h1, h2 := hash(key)

	if c.T1[h1] == key {
		c.T1[h1] = 0
	}

	if c.T2[h2] == key {
		c.T2[h2] = 0
	}
	return
}

func main() {
	c := NewCuckoo()

	// insert the keys.
	cnt := 0
	x := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, key := range x {
		c.insert(key, cnt)
	}

	// look up for the key "1" and "10".
	fmt.Println("key:1  ", c.lookup(1))  // key:1   true
	fmt.Println("key:10 ", c.lookup(10)) // key:10  false

	// delete the key "3".
	fmt.Println("before: ", *c) // before:  {[0 8 6 0 4 0 9 0 5 7] [0 1 0 0 0 2 0 0 3 0]}
	c.delete(3)
	fmt.Println("after:  ", *c) // after:   {[0 8 6 0 4 0 9 0 5 7] [0 1 0 0 0 2 0 0 0 0]}
}
