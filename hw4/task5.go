package hw4

import (
	"fmt"
	"sync"
	"time"
)

func reciever(ch chan int, done chan bool, wg *sync.WaitGroup) {
	for {
		select {
		case <-done:
			wg.Done()
			return
		case v := <-ch:
			fmt.Printf("Received %d\n", v)
		}
	}
}

func sender(ch chan int, done chan bool, wg *sync.WaitGroup) {
	for i := 0; ; i++ {
		select {
		case <-done:
			wg.Done()
			return
		default:
			ch <- i
		}
	}
}

func Ex5() {

	// getting number of seconds to work
	fmt.Printf("Enter the number of seconds: ")
	var n int
	fmt.Scan(&n)

	wg := sync.WaitGroup{}
	wg.Add(2)

	// creating channels
	ch := make(chan int)
	defer close(ch)
	done := make(chan bool, 2)
	defer close(done)

	// creating goroutines
	go sender(ch, done, &wg)
	go reciever(ch, done, &wg)

	// sleeping for n seconds
	time.Sleep(time.Duration(n) * time.Second)

	// sending done signals for 2 goroutines
	done <- true
	done <- true

	// waiting for goroutines to finish
	wg.Wait()

}
