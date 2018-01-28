/*
* Author: Yogi Hardi <yogi.hardi@gmail.com>
 */
package main

import (
	"fmt"
	"sync"
	"math/rand"
	"sort"
)

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
	/*
		    this should print something like this:
			1
			2
			3
			4
			5
			6
			7
			8
		    make sure that you close both channels and program should exit successfully at the end.

	*/
}

func merge(a, b <-chan int) <-chan int {
	// this should take a and b and return a new channel which will send all values from both a and b
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in a and b.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(2)
	go output(a)
	go output(b)

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()


	// sort
	list := make([]int, 8)
	i := 0
	for n := range out {
		list[i] = n
		i++
	}
	sort.Ints(list)

	result := make(chan int)
	go func() {
		for _,v := range list {
			result <- v
		}
		close(result)
	}()

	return result
}

func asChan(vs ...int) <-chan int {
	// this should reutrn a channel and send `vs` values randomly to it.
	out := make(chan int)

	go func() {
		perm := rand.Perm(len(vs))
		for i := range perm{
			out <- vs[i]
		}
		close(out)
	}()
	return out
}

