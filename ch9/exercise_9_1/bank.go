// Add a function Withdraw(amount int) bool to the bank1 program.
package bank

type withdraw struct {
	amount int
	ch     chan<- bool
}

var deposits = make(chan int)         // send amount to deposit
var balances = make(chan int)         // receive balance
var withdrawals = make(chan withdraw) // withdraw amount

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdrawals <- withdraw{amount, ch}
	return <-ch
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case wd := <-withdrawals:
			if wd.amount > balance {
				wd.ch <- false
			} else {
				balance -= wd.amount
				wd.ch <- true
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
