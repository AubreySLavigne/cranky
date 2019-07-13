package main

import (
	"log"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

const batchSize = 1000000

func main() {

	wg = sync.WaitGroup{}

	// need to implement the counter to make see how many ongoing
	// processes we have

	wg.Add(runtime.NumCPU())
	in := make(chan int64)
	out := make(chan int64)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func(in chan int64, out chan int64) {
			for start := range in {
				for i := start; i < start+batchSize; i++ {
					num := proposed(i)
					if num.isCranky() {
						log.Print(num)
						// out <- int64(num)
					}
				}
			}

			wg.Done()
		}(in, out)
	}

	for i := 0; i < 100000000; i++ {
		in <- int64(i * batchSize)
	}

	close(in)

	wg.Wait()
}
