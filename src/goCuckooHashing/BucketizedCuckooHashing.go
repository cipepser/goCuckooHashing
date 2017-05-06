package goCuckooHashing

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/big"
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

type BucketizedCuckoo struct {
	T1, T2 [N][BCKSIZE]int64
}

func NewBucketizedCuckoo() *BucketizedCuckoo {
	return new(BucketizedCuckoo)
}

func (c *BucketizedCuckoo) lookup(key int64) bool {
	h1, h2 := hash(key)
	
	for _, bucket := range c.T1[h1] {
		if bucket == key {
			return true
		}
	}

	for _, bucket := range c.T2[h2] {
		if bucket == key {
			return true
		}
	}
	
	return false
}

func (c *BucketizedCuckoo) insert(key int64, cnt int) {
	for cnt < MAXLOOP {
		h1, h2 := hash(key)
		switch cnt % 2 {
		case 0:
			for i, bucket := range c.T1[h1] {
				if bucket == 0 {
					c.T1[h1][i] = key
					return
				} else if i == BCKSIZE - 1 {
					key, c.T1[h1][i] = c.T1[h1][i], key
				}
			}

		case 1:
			for i, bucket := range c.T2[h2] {
				if bucket == 0 {
					c.T2[h2][i] = key
					return
				} else if i == BCKSIZE - 1 {
					key, c.T2[h2][i] = c.T2[h2][i], key
				}
			}
		}
		cnt++
	}
	panic("fail to insert. you must reconstruct the Table...")
}

func (c *BucketizedCuckoo) delete(key int64) {
	h1, h2 := hash(key)

	for i, bucket := range c.T1[h1] {
		if bucket == key {
			c.T1[h1][i] = 0
		}
	}

	for i, bucket := range c.T1[h1] {
		if bucket == key {
			c.T2[h2][i] = 0
		}
	}
	
	return
}

func BucketizedCuckoo() {
	c := NewBucketizedCuckoo()

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
