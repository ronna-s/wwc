package main

import "sync"

type Message string

type Shuffler interface {
	Shuffle()
}

type Client chan Message

type Broadcaster struct {
	clients []Client
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		cli := make(chan Message)

	}()

}
