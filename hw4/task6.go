package hw4

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// function that returns by receiving any value from cnannel
func closing_by_channel(close chan bool, wg *sync.WaitGroup) {
	for {
		select {

		// reading from channel
		case <-close:
			fmt.Printf("\n[!] Goroutine with channel has finished\n\n")
			wg.Done()
			return

		// if it's empty, do some work
		default:
			time.Sleep(1 * time.Second)
			fmt.Printf("Goroutine with channel is working\n")
		}
	}
}

// function that returns by receiving done signal from the context
func closing_by_context_cancel(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {

		// reading from context
		case <-ctx.Done():
			fmt.Printf("\n[!] Goroutine with context with cancel has finished\n\n")
			wg.Done()
			return

		// if it's not done, do some work
		default:
			time.Sleep(1 * time.Second)
			fmt.Printf("Goroutine with context with cancel is working\n")
		}
	}
}

// function that ends by passing some time
func closing_by_context_timeout(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {

		// if time is out, returning
		case <-ctx.Done():
			fmt.Printf("\n[!] Goroutine with context with timeout has finished\n\n")
			wg.Done()
			return

		// if it's not done, do some work
		default:
			time.Sleep(1 * time.Second)
			fmt.Printf("Goroutine with context with timeout is working\n")
		}
	}
}

func Ex6() {
	// creating a waitgroup to wait for all workers to finish
	var wg sync.WaitGroup
	wg.Add(3)

	// creating a channel and calling the function in goroutine
	ch := make(chan bool)
	go closing_by_channel(ch, &wg)

	// creating a context with cancel and calling the function in goroutine
	ctx_w_cancel, cancel := context.WithCancel(context.Background())
	go closing_by_context_cancel(ctx_w_cancel, &wg)

	// creating a context with timeout and calling the function in goroutine
	ctx_w_timeout, cnsl := context.WithTimeout(context.Background(), 2*time.Second)
	defer cnsl()
	go closing_by_context_timeout(ctx_w_timeout, &wg)

	// waiting some time
	time.Sleep(5 * time.Second)

	// cancelling the context
	cancel()
	// sending a value to the channel
	ch <- true

	// waiting for all workers to finish
	wg.Wait()
}
