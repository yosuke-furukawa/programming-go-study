package bank

import (
	"fmt"
	"testing"
)

func TestSync(t *testing.T) {
	for i := 0; i < 10000; i++ {
		Deposit(i)
		if Balance() != i {
			t.Errorf("amount is not correct")
		}
		w := Withdraw(i)
		if !w {
			t.Errorf("amount is not correct")
		}
		if Balance() != 0 {
			t.Errorf("amount is not correct")
		}

	}

}

func TestGoroutine(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
