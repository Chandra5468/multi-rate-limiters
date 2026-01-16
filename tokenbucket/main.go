package main

func main() {

}

/* USE THIS IN MAIN OR ANY FUNCTION.
// create a channel that can hold 5 tokens (burst size)
requests := make(chan int, 5)

for i := 1; i <= 5; i++ {
	requests <- i
}
close(requests)

// This limiter will allow a token every 500 milliseconds (2 per second)
limiter := time.NewTicker(500 * time.Millisecond)
defer limiter.Stop()

// consume tokens from the requests channel, rate-limited by the limiter ticker
for req := range requests {
	<-limiter.C
	fmt.Printf("Request %d processed at %s \n", req, time.Now().Format("15:04:05.000"))
}
*/
