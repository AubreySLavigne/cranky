package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

var sum int

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	for i := 3; i < 12; i++ {
		num := newNumber(i)
		for j := 1; j < len(num.digits); j++ {
			num.slicePos = j
			findCrankyNumbers(num)
		}
	}

	log.Printf("Total: %d", sum)
}

func findCrankyNumbers(num number) {
	for i := 0; i < 10; i++ {
		newNum := num.copy()
		if newNum.propose(i) {
			if newNum.hasMore() {
				findCrankyNumbers(newNum)
			} else if newNum.check() {
				product := newNum.productInt()
				sum += newNum.asInt()
				log.Printf("%s >> %d (sum: %d)", newNum.summary(), product, sum)
			}
		}
	}
}
