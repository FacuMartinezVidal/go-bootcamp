package goroutines

import (
	"fmt"
	"time"
)

func Goroutines() {
	//don't forget about select and defer
	done := make(chan bool)
	go greet("Gophers", done)
	go slowGreet("Gophers", done)
	//with <- we expect a value from the channel

	for range done {
	}
}

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello", phrase)
	//with <- we send a value to the channel
	doneChan <- true
	//close the channel if you know what function is the last one to write to it
	close(doneChan)
}
