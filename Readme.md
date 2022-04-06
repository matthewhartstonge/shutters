# Shutters :zipper_mouth_face:

`shutters` encapsulates tasks that need to be called on shutdown.

[![Go Reference](https://pkg.go.dev/badge/github.com/matthewhartstonge/shutters.svg)](https://pkg.go.dev/github.com/matthewhartstonge/shutters)
[![Go Report Card](https://goreportcard.com/badge/github.com/matthewhartstonge/shutters)](https://goreportcard.com/report/github.com/matthewhartstonge/shutters)

Ever had to gracefully shut things down, but then had multiple resources you needed to track in order to shut down
gracefully?

Maybe shutters is for you!

## Getting Started

`shutters` makes use of go mod, you can install it by using go get:

```shell
go get github.com/matthewhartstonge/shutters
```

## Examples

All examples are runnable from the [_examples](./_examples) folder.

### [main](./_examples/main/main.go)

Hopefully, this provides an example of the common use case:

```go
package main

import (
	"fmt"

	"github.com/matthewhartstonge/shutters"
)

func someDeeplyNestedFuncCalledFromSomewhereElse() {
	closeMe := func() {
		// Create a closure that can close the resources you need...
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
```

### [instanced](./_examples/instanced/main.go)

If you want to specifically have your own instance for some form of sub-process, you can do that too:

```go
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

	// todo(@lazybrogrammer): sorry not sorry, I'm too 10x to close the blinds.
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
```

### [gofunky](./_examples/gofunky/main.go)

Congratz, you made it this far!
As your reward you get some more code - an example with a go func for some fun:

```go
package main

import (
	// Standard Library Imports
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"

	// External Imports
	"github.com/matthewhartstonge/shutters"
)

func main() {
	// have some resource we want to ensure is closed on shutdown.
	isClosed := false

	defer func() {
		// close all the shutters at the end of main.
		shutters.Close()
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovering panic'd go thread due to:", r)
			}

			fmt.Println("my work here is done.. 😤")
			wg.Done()
		}()

		closeShutters := func() {
			// Create a closer function to alter whatever resources we require
			// to be gracefully shutdown.
			isClosed = true
			fmt.Println("the shutters have been gracefully closed?", isClosed)
		}

		// Bind in the shutter
		shutters.Add(closeShutters)
		fmt.Println("the shutters are closed?", isClosed)

		if i, _ := rand.Int(rand.Reader, big.NewInt(100)); i.Int64() > 50 {
			// let's totally panic randomly! That helps things!
			panic("lol I broke it")
		}

		// Jus gonna take some time off here to process my thoughts...
		time.Sleep(time.Second * 2)
	}()

	// Wait for our wonderful process to finish...
	wg.Wait()
}
```
