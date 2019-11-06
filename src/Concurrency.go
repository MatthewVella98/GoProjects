package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	/*Can't make the code infinitely fast, as it depends up to your cpu architecutre
	and the number of cores.
	go count("Sheep")
	go count("Fish")
	Running all the code in go routines. Nothing is logged to console.
	When the main go routine finishes, the program exists regardless of what other go routines are doing.
	previously, the main go routine never finished since it had its own infinite loop.
	Since no more code exists (except 2 go routines only), the program will terminate.

	if you sleep for 2 seconds, it outputs for 2 seconds and terminate.
	You can use scanln, so that the goroutines run and the program doesnt terminate since it awaits the call from the keyboard.
	This is not very useful though. The best is to use the wait group.
	*/

	go count("Sheep") // If this is not specified as a goroutine, it will run forever without counting fish.
	count("Fish")
	//------------------------------- Go Routines With Wait group

	var wg sync.WaitGroup //This is a counter of goroutines.
	wg.Add(1)

	go func() { // Anonymouse function -> Runs immediately.
		countTill5("Sheep")
		wg.Done()
	}()

	wg.Wait() // This will block until 0 (Until all the goroutines are finished).

	//------------------------------ Channels - A way for goroutines to communicate with each other.
	/* The count function so far has been outputting directly to the terminal. But what if we want to communicate
	back with the main go routine? Channels. */

}

func count(thing string) { // Infinite loop function.
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func countTill5(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func countUsingChannels(thing string, c chan string) {
	for i := 0; i <= 5; i++ {
		c <- thing // Send 'thing' over the channel.
	}
}
