package main

import (
	"fmt"
)

func main() {
	for i := 10; i < 1000000000; i++ {
		num := proposed(i)
		if num.isCranky() {
			fmt.Println(num)
		}
	}
}
