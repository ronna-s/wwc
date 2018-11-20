package main

import (
	"errors"
	"fmt"
)

//START OMIT
func CanRecover() error {
	return errors.New("something went wrong...")
}

func CantRecover() error {
	return errors.New("something went awefully wrong...")
}

func main() {
	// Note the if-declaration syntax (similar to for loops).
	if err := CanRecover(); err != nil {
		fmt.Println(err)
	}

	if err := CantRecover(); err != nil {
		panic(err)
	}

	CantRecover()
}

//DONE OMIT
