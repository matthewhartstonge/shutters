package main

import (
	// Standard Library Imports
	"fmt"

	// External Imports
	"github.com/matthewhartstonge/shutters"
)

func someDeeplyNestedFuncCalledFromSomewhereElse() {
	closeMe := func() {
		// Create a closure to kill off some resources gracefully...
		fmt.Println("I've gracefully closed some things! 😸")
	}
	shutters.Add(closeMe)

	fmt.Println("Processing hardly! or hardly processing... 🤔")
}

func main() {
	// Run the things! 🏃‍♂️🏃‍♀️
	someDeeplyNestedFuncCalledFromSomewhereElse()

	// Then a while later... 🤐
	shutters.Close()
}
