package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// a buffered channel of size 3
	dataQueue := make(chan int, 3)
	wg := sync.WaitGroup{}
	wg.Add(1)
	// Consumer: Process one item every second (very slow)
	go func() {
		defer wg.Done()
		for item := range dataQueue {
			fmt.Printf("✅ Consumed: %d\n", item)
			time.Sleep(1 * time.Second)
		}
	}()

	// Producer: Tries to send 10 items very quickly
	for i := 1; i <= 10; i++ {
		select {
		case dataQueue <- i:
			fmt.Printf("📤 Sent to buffer: %d\n", i)
		default:
			// This branch runs immediately if the channel is full. Dropping
			fmt.Printf("❌ Buffer full! Dropped: %d\n", i)
		}
		// Small sleep to simulate fast production
		time.Sleep(100 * time.Millisecond)
	}
	close(dataQueue)
	wg.Wait()
}
