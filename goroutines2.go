package main

import (
	"fmt"
)

// START OMIT
func iterate(until int) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < until; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func main() {
	iterCh := iterate(10)

	//ranges until channel is closed
	for i := range iterCh {
		fmt.Println(i)
	}
}
// DONE OMIT
