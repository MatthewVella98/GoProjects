package lego

import (
	"sync"

	"../util"
)

// -----------------------------------------------------------------------------
// Implements a 1-place buffer, allowing the in and out channels to be decoupled
// by a single message.
// -----------------------------------------------------------------------------
func Id(in, out chan int) {
	for {
		i := <-in
		out <- i
	}
}

// -----------------------------------------------------------------------------
// Computes the successor of an integer.
// -----------------------------------------------------------------------------
func Succ(in, out chan int) {
	for {
		i := <-in
		out <- i + 1
	}
	
}

// -----------------------------------------------------------------------------
// Eats up the value that is presented at the input channel.
// -----------------------------------------------------------------------------
func Sink(in chan int) {
	for {
		<-in
	}

}

// -----------------------------------------------------------------------------
// Removes the first input on the channel and allowing the rest (the tail) to
// pass though. 
// -----------------------------------------------------------------------------
func Tail(in, out chan int) {
	for i := 0; i < 1; i++ {
		// remove the first
		<-in
	}

	j := <-in
	out <- j

}

// -----------------------------------------------------------------------------
// Prepends the integer n to the head of the output the channel.
// -----------------------------------------------------------------------------
func Prefix(n int, in, out chan int) {
	for i := 0; true; i++ {
		if i == 0 {
			out <- n
		}

		j := <-in
		out <- j
	}

}

// -----------------------------------------------------------------------------
// Adds two numbers presented at the input channels in the *predetermined* order
// of in_x followed by in_y.
// -----------------------------------------------------------------------------
func PlusSerial(in_x, in_y, out chan int) {
	for {
		x := <-in_x
		y := <-in_y
		out <- x + y
	}
}

// -----------------------------------------------------------------------------
// Adds two numbers presented at the input channels, the first to be read either
// at in_x or at in_y. This is a wrong implementation of Plus because in certain
// cases we can add two successive numbers from in_x or two successive numbers
// from in_y, whereas our intended behaviour is that of adding one number from
// in_x and the other from in_y together.
// -----------------------------------------------------------------------------
func PlusWrong(in_x, in_y, out chan int) {
	for {
		x, y := 0, 0
		for i := 0; i < 2; i++ {
			select {
			case x = <-in_x:
			case y = <-in_y:
			}
		}
		out <- x + y
	}
}

// -----------------------------------------------------------------------------
// Adds two numbers presented at the input channels in *any* order.

// This implementation, although a corret one as it works as intended by either
// adding in_x + in_y or in_y + in_x, is naive, because as seen in the select
// statement, we are enumerating all the possible interleavings in which inputs
// can appear on the input channels in_x and in_y:
//   1. in_x can be read first, followed by in_y
//   2. in_y can be read first, followed by in_x.
// The number of possible permutations (i.e. not combinations, since order
// matters) in which inputs can be interleaved grows according to the number of
// input channels. This number is substantial and can quickly lead to
// unmanageable code!
// -----------------------------------------------------------------------------
func PlusNaive(in_x, in_y, out chan int) {
	for {
		x, y := 0, 0
		select {
		case x = <-in_x:
			y = <-in_y
		case y = <-in_y:
			x = <-in_x
		}
		out <- x + y
	}
}

// -----------------------------------------------------------------------------
// Adds two numbers presented at the input channels in *any* order.

// This implementation employs the sync.WaitGroup.
// -----------------------------------------------------------------------------
func PlusWG(in_x, in_y, out chan int) {
	var wg sync.WaitGroup
	for {
		x, y := 0, 0
		wg.Add(2)
		go func() {
			defer wg.Done()
			x = <-in_x
		}()
		go func() {
			defer wg.Done()
			y = <-in_y
		}()
		wg.Wait()
		out <- x + y
	}
}

// -----------------------------------------------------------------------------
// Adds two numbers presented at the input channels in *any* order.

// This implementation employs the Par (or fork-join) concurreny abstraction
// pattern, and works exactly like the one using sync.WaitGroup.
// -----------------------------------------------------------------------------
func Plus(in_x, in_y, out chan int) {
	for {
		x, y := 0, 0
		util.Par(func() {
			x = <-in_x
		}, func() {
			y = <-in_y
		})
		out <- x + y
	}
}

// -----------------------------------------------------------------------------
// Replicates the number presented at the input channel to out_x and out_y.

// This implementation employs the sync.WaitGroup.
// -----------------------------------------------------------------------------
func DeltaWG(in, out_x, out_y chan int) {
	var wg sync.WaitGroup //waitgroup - counter
	for {
		i := <-in
		wg.Add(2) //waiting for 2 go routines
		go func() {
			out_x <- i //sending over out_x channel
			wg.Done()  //decrements counter by 1
		}()
		go func() {
			out_y <- i
			wg.Done()
		}()
		wg.Wait() //waiting until all go routines are finished
	}
}

// -----------------------------------------------------------------------------
// Replicates the number presented at the input channel to out_x and out_y.

// This implementation employs the Par (or fork-join) concurreny abstraction
// pattern, and works exactly like the one using sync.WaitGroup.
// -----------------------------------------------------------------------------
func Delta(in, out_x, out_y chan int) {
	for {
		i := <-in
		util.Par(func() {
			out_x <- i
		}, func() {
			out_y <- i
		})
	}
}

// -----------------------------------------------------------------------------
// Produces the list of integers greater than zero by outputting the next
// number each time this is read from the output channel.
// -----------------------------------------------------------------------------
func Nos(out chan int) {
	x := make(chan int) //creating 3 integer channels
	y := make(chan int)
	z := make(chan int)

	go Prefix(0, z, x)
	go Delta(x, y, out)
	go Succ(y, z)
}

// -----------------------------------------------------------------------------
// Produces the running total of the numbers presented in the input channel in
// (an integrator).
// -----------------------------------------------------------------------------
func Int(in, out chan int) {
	x := make(chan int)
	y := make(chan int)
	z := make(chan int)

	go Prefix(0, y, z)
	go Plus(in, z, x)
	go Delta(x, out, y)
}

// -----------------------------------------------------------------------------
// Outputs the numbers that are seperated by a difference of two.
// -----------------------------------------------------------------------------
func Pairs(in, out chan int) {
	x := make(chan int)
	y := make(chan int)
	z := make(chan int)

	go Delta(in, x, z)
	go Tail(x, y)
	go Plus(y, z, out)
}

// -----------------------------------------------------------------------------
// Produces the list of numbers from the fibonacci sequence starting from 0.
// -----------------------------------------------------------------------------
func Fib(out chan int) {
	w := make(chan int)
	x := make(chan int)
	y := make(chan int)
	z := make(chan int)

	go Prefix(1, y, z)
	go Prefix(0, w, x)
	go Delta(x, y, out)
	go Pairs(x, y)
}

// -----------------------------------------------------------------------------
// Produces the list of square numbers starting from 1.
// -----------------------------------------------------------------------------
func Squares(out chan int) {
	x := make(chan int)
	y := make(chan int)

	go Nos(x)
	go Int(x, y)
	go Pairs(y, out)
}
