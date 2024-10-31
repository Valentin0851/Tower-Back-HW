package hw4

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

// creating worker
func worker4(ctx context.Context, worker_id int, ch chan int, wg *sync.WaitGroup) {
	var varialbe int
	for {
		// waiting for signal from main or cancelation
		select {

		// if the context is canceled, return
		case <-ctx.Done():
			fmt.Printf("Worker %d: Canceled\n", worker_id)
			wg.Done()
			return

		// if there's a data in the channel, print it
		case varialbe = <-ch:
			fmt.Printf("Worker %d: %d\n", worker_id, varialbe)
		}
	}
}

func Ex4() {
	ctx, cancel := context.WithCancel(context.Background())

	// waiting for in goroutine to finish
	sigChannel := make(chan os.Signal, 1)
	defer close(sigChannel)
	signal.Notify(sigChannel, os.Interrupt)

	go func() {
		<-sigChannel
		// calling cancelation on pressing ctrl+c
		cancel()
	}()

	// creating a count of workers variable and reading it
	var count_of_workers int
	fmt.Printf("Enter the number of workers: ")
	fmt.Scan(&count_of_workers)

	// creating one channel for all the workers with buffer to ensure a deadlock will not occur
	ch := make(chan int, count_of_workers)
	defer close(ch)

	// creating a waitgroup to wait for all workers to finish
	var wg sync.WaitGroup
	wg.Add(count_of_workers)

	// creating workers in goroutines
	for i := 0; i < count_of_workers; i++ {
		go worker4(ctx, i, ch, &wg)
	}

	// infinite loop
	for varialbe := 0; ; varialbe++ {
		select {

		// if the context is canceled, return
		case <-ctx.Done():
			// waiting for all workers to finish
			wg.Wait()
			return

		// sending data to the channel
		default:
			ch <- varialbe
		}
	}

}
