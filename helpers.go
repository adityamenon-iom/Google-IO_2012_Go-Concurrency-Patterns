package main

func main() {
	var value int

	// Declaring and initializing.
	var c chan int
	c = make(chan int)
	// or
	c := make(chan int) // HL

	// Sending on a channel.
	c <- 1 // HL

	// Receiving from a channel.
	// The "arrow" indicates the direction of data flow.
	value = <-c // HL

	_ = value
}
