package main

import (
	"fmt"
	"sync"
)

//START OMIT
func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		//start a goroutine using the keyword go
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}

	wg.Wait() //wait until wg.Done() is called 10 times...
}
//DONE OMIT