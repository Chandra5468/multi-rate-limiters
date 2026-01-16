package main

import (
	"fmt"
	"time"
)

// What happens if an unbuffered channel is used instead of a buffered channel

/*
In this scenario, the sender (the Ticker) and the receiver (your Loop) must meet at the exact same moment.
If your loop is even slightly late, the Ticker waits, pushing the entire schedule forward.

This is called accumulative drift.
*/

func main() {
	// An unbuffered channel (size 0)
	// This requires a perfect handshake to send/receive
	unbufferedC := make(chan time.Time)

	// Fake ticker go routine
	go func() {
		for {
			time.Sleep(2 * time.Second)
			fmt.Println("[Ticker] Ready to send tick...")
			unbufferedC <- time.Now() // This will BLOCK until the loop is ready
			fmt.Println("[Ticker] Tick delivered!")
		}
	}()

	// Processing loop
	for i := 1; i <= 3; i++ {
		t := <-unbufferedC
		fmt.Printf("Request %d processed at %s\n", i, t.Format("15:04:05.000"))

		// SIMULATED DELAY: Imagine your database or logic takes 200ms
		// During this time, the Ticker is blocked and cannot start its next 500ms countdown!
		time.Sleep(200 * time.Millisecond)
		fmt.Println("-------------------------------------------")
	}
}
