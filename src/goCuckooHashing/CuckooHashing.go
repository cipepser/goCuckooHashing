package goCuckooHashing

import (
	"crypto/md5"
	"encoding/hex"
	"math/big"
)

func hash(key int64) (h1, h2 int64) {
	hasher := md5.New()
	hasher.Write([]byte(string(key)))

	h := hex.EncodeToString(hasher.Sum(nil))

	t1, _ := new(big.Int).SetString(h[:int(len(h)/2)], 16)
	t2, _ := new(big.Int).SetString(h[int(len(h)/2):], 16)
	h1 = t1.Rem(t1, big.NewInt(N)).Int64()
	h2 = t2.Rem(t2, big.NewInt(N)).Int64()

	return h1, h2
}

type Cuckoo struct {
	T1, T2 [N]int64
}

func NewCuckoo() *Cuckoo {
	return new(Cuckoo)
}

func (c *Cuckoo) Lookup(key int64) bool {
	h1, h2 := hash(key)

	if c.T1[h1] == key || c.T2[h2] == key {
		return true
	} else {
		return false
	}
}

func (c *Cuckoo) Insert(key int64, cnt int) {
	for cnt < MAXLOOP {
		h1, h2 := hash(key)
		switch cnt % 2 {
		case 0:
			if c.T1[h1] == 0 {
				c.T1[h1] = key
				return
			}
			key, c.T1[h1] = c.T1[h1], key

		case 1:
			if c.T2[h2] == 0 {
				c.T2[h2] = key
				return
			}

			key, c.T2[h2] = c.T2[h2], key
		}
		cnt++
	}
	// panic("fail to insert. you must reconstruct the Table...")
}

func (c *Cuckoo) Delete(key int64) {
	h1, h2 := hash(key)

	if c.T1[h1] == key {
		c.T1[h1] = 0
	}

	if c.T2[h2] == key {
		c.T2[h2] = 0
	}
	return
}
