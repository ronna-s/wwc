package main

import (
	"fmt"
)

// START OMIT
func iterate(until int) chan int {
	ch := make(chan int) //create an unbuffered channel (of size 0)

	//start a goroutine that writes the numbers from 1 to 10 to the channel
	go func() {
		for i := 0; i < until; i++ {
			ch <- i
		}
		close(ch) //close the channel when done.
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
