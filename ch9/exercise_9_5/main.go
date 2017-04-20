// Write a program with two goroutines that send messages back and forth over
// two unbuffered channels in ping-pong fashion. How many communications per
// second can the program sustain?
package main

import (
	"flag"
	"fmt"
	"time"
)

var seconds = flag.Int("seconds", 5, "number of seconds for goroutine communication")

func main() {
	flag.Parse()
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch2 <- 1
		for val := range ch1 {
			ch2 <- val + 1
		}
	}()

	var count int
	for start := time.Now(); time.Since(start) < time.Duration(*seconds)*time.Second; {
		count = <-ch2
		ch1 <- count + 1

	}
	close(ch1)
	close(ch2)

	fmt.Printf("communications per second = %f\n", float64(count)/float64(*seconds))
}
