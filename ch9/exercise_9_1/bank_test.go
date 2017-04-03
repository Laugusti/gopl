package bank

import (
	"testing"
)

func TestWithdraw(t *testing.T) {
	if Withdraw(0) == false {
		t.Errorf("Withdraw failed! Balance: %d, Withdraw: %d", Balance(), 0)
	}
	if Withdraw(1) == true {
		t.Errorf("Withdraw succeeded! Balance: %d, Withdraw: %d", Balance(), 1)
	}
	Deposit(100)
	if Withdraw(1) == false {
		t.Errorf("Withdraw failed! Balance: %d, Withdraw: %d", Balance(), 1)
	}
	if Withdraw(99) == false {
		t.Errorf("Withdraw failed! Balance: %d, Withdraw: %d", Balance(), 99)
	}
	if Withdraw(1) == true {
		t.Errorf("Withdraw succeeded! Balance: %d, Withdraw: %d", Balance(), 1)
	}
	if Withdraw(0) == false {
		t.Errorf("Withdraw failed! Balance: %d, Withdraw: %d", Balance(), 0)
	}
}
