package main

import (
	"fmt"
	"math/rand"
)

// START OMIT
type Name string
type Names []string

func (names Names) Shuffle() {
	for i := range names {
		r := rand.Intn(len(names))
		names[i], names[r] = names[r], names[i]
	}
}

type Shuffler interface {
	Shuffle()
}

func Shuffle(s Shuffler) {
	fmt.Println(s)
	s.Shuffle()
	fmt.Println(s)
}

func main() {
	names := Names{"Jessy", "Matt", "Daniella"}
	Shuffle(names)
}

// DONE OMIT
