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
