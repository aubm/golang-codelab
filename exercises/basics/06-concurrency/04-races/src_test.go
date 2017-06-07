package basics_test

import (
	"testing"
	"time"

	. "github.com/aubm/golang-codelab/exercises/basics/06-concurrency/04-races"
)

func TestBalanceAdd(t *testing.T) {
	balance := &Balance{}
	doTransactions := func() {
		go balance.Add(1)
		go balance.Add(-1)
		go balance.Add(4)
		go balance.Add(3)
		go balance.Add(-7)
	}
	for i := 0; i < 100000; i++ {
		go doTransactions()
	}
	time.Sleep(2 * time.Second)
	if balance.Amount() != 0 {
		t.Errorf("Balance is not concurrent safe!")
	}
}
