package goCuckooHashing

type BucketizedCuckoo struct {
	T1, T2 [N][BCKSIZE]int64
}

func NewBucketizedCuckoo() *BucketizedCuckoo {
	return new(BucketizedCuckoo)
}

func (c *BucketizedCuckoo) Lookup(key int64) bool {
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

func (c *BucketizedCuckoo) Insert(key int64, cnt int) {
	for cnt < MAXLOOP {
		h1, h2 := hash(key)
		switch cnt % 2 {
		case 0:
			for i, bucket := range c.T1[h1] {
				if bucket == 0 {
					c.T1[h1][i] = key
					return
				} else if i == BCKSIZE-1 {
					key, c.T1[h1][i] = c.T1[h1][i], key
				}
			}

		case 1:
			for i, bucket := range c.T2[h2] {
				if bucket == 0 {
					c.T2[h2][i] = key
					return
				} else if i == BCKSIZE-1 {
					key, c.T2[h2][i] = c.T2[h2][i], key
				}
			}
		}
		cnt++
	}
}

func (c *BucketizedCuckoo) Delete(key int64) {
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
