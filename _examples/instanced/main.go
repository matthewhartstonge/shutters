package main

import (
	// Standard Library Imports
	"fmt"

	// External Imports
	"github.com/matthewhartstonge/shutters"
)

func aFunction() {
	blinds := shutters.NewManager()

	closeMe := func() {
		// Create a closure to kill off some resources gracefully...
		fmt.Println("aFunction has gracefully closed some things! 😸")
	}
	blinds.Add(closeMe)

	fmt.Println("aFunction here! Processing hardly! or hardly processing... 🤔")

	// todo: close the blinds... (oops..)
}

func anotherFunction() {
	shutter := shutters.NewManager()

	closeMe := func() {
		// Create a closure to kill off some resources gracefully...
		fmt.Println("anotherFunction has gracefully closed some other things! 😸")
	}
	shutter.Add(closeMe)

	fmt.Println("anotherFunction here! Still processing hardly! otherwise hardly processing... 🤔")

	shutter.Close()
}

func main() {
	closeMe := func() {
		// Create a closure to kill off some resources gracefully...
		fmt.Println("Main shutter gracefully closing some things! 😸")
	}
	shutters.Add(closeMe)

	// Run the things! 🏃‍♂️🏃‍♀️
	aFunction()
	anotherFunction()

	// Then a while later... 🤐
	shutters.Close()
}
