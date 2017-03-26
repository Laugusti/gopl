package main

import "fmt"
import "os"
import "time"

func main() {
	abort := make(chan struct{})
	tick := time.Tick(1 * time.Second)

	fmt.Println("Commencing countdown. Press return to abort.")
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
