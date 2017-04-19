// Construct a pipeline that connects an arbitrary number of goroutines with channels.
// What is the maximum number of pipleline stages you can create without running out of
// memory? How long does a value take to transit the entire pipeline?
package main

import (
	"flag"
	"fmt"
	"time"
)

var stages = flag.Int("stages", 100000, "number of pipeline stages")

func main() {
	flag.Parse()
	ch := make(chan struct{})
	ch1 := ch

	// create goroutine pipeline
	for i := 0; i < *stages; i++ {
		ch2 := make(chan struct{})
		go func(in <-chan struct{}, out chan<- struct{}) {
			out <- <-in

		}(ch1, ch2)
		ch1 = ch2

	}

	// record time for value to propagate through pipeline
	start := time.Now()
	ch <- struct{}{}
	<-ch1
	fmt.Printf("Duration: %v\n", time.Since(start))
}
