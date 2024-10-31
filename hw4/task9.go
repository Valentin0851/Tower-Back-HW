package hw4

import (
	"context"
	"fmt"
	"sync"
)

// function for sending doubles to channel
func double(value_channel chan int, double_channel chan int, ctx context.Context) {
	for {
		select {

		// if there's a value in the channel, double it and send it to the channel
		case value := <-value_channel:
			double_channel <- value * 2

		// if the context is canceled, return
		case <-ctx.Done():
			return
		}
	}
}

// function for printing doubles from channel
func print_double(double_channel chan int, wg *sync.WaitGroup, ctx context.Context) {
	for {
		select {

		// if there's a value in the channel, print it
		case value := <-double_channel:
			fmt.Printf("%v is a double\n", value)
			wg.Done()

		// if the context is canceled, return
		case <-ctx.Done():
			return
		}
	}
}

func Ex9() {
	// creating a waitgroup
	var wg sync.WaitGroup
	wg.Add(10)

	// creating channels
	value_channel := make(chan int)
	double_channel := make(chan int)

	// creating a context
	ctx, cancel := context.WithCancel(context.Background())

	// creating goroutines for double and print_double
	go double(value_channel, double_channel, ctx)
	go print_double(double_channel, &wg, ctx)

	// sending some values to the channels
	for i := 0; i < 10; i++ {
		value_channel <- i
	}

	// waiting for goroutines to finish
	wg.Wait()

	// canceling the context
	cancel()

}
