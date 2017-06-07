package basics

import "sync/atomic"

type Balance struct {
	amount int64
}

func (b *Balance) Amount() int64 {
	return b.amount
}

// Oops, it seems that this Add function is not thread safe.
// Can you do something about it?
func (b *Balance) Add(delta int64) {
	atomic.AddInt64(&b.amount, delta)
}
