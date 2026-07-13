// This program constructs a pipeline of an arbitrary number of goroutines
// connected by channels
package main

import (
	"fmt"
	"time"
)

const stages = 100

// pipeline builds a chain of n stage goroutines. Each stage simply
// forwards whatever it receives to the next stage, so the value written
// to in comes out of out unchanged after passing through n goroutines
// and n channels.
func pipeline(n int) (in chan<- int, out <-chan int) {
	first := make(chan int)
	cur := first
	for i := 0; i < n; i++ {
		next := make(chan int)
		go stage(cur, next)
		cur = next
	}
	return first, cur
}

func stage(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v
	}
	close(out)
}

func main() {

	start := time.Now()
	in, out := pipeline(stages)
	fmt.Printf("built a %d-stage pipeline in %s\n", stages, time.Since(start))

	sendTime := time.Now()
	in <- 1
	close(in)
	v := <-out
	fmt.Printf("value %d transited %d stages in %s\n", v, stages, time.Since(sendTime))
}
