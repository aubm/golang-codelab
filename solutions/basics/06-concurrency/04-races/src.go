package basics

import "sync/atomic"

type Balance struct {
	amount int64
}

func (b *Balance) Amount() int64 {
	return b.amount
}

func (b *Balance) Add(delta int64) {
	atomic.AddInt64(&b.amount, delta)
}
