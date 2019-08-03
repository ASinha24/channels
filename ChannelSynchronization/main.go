package main

import (
	"fmt"
	"time"
)

//Here’s an example of using a blocking receive to wait for a goroutine to finish. When waiting for multiple goroutines to finish, you may prefer to use a WaitGroup.
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(5 * time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool)
	go worker(done)

	<-done //a blocking receive to wait for finish

}
