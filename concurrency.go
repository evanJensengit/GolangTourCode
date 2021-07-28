package main

import (
	"fmt"
	"sync"
	"time"
)

////--------youtube tutorials https://www.youtube.com/watch?v=yNOe3STbtGE

//runs a for loop with a sleep of half a second
func concurExample() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
		time.Sleep(time.Second / 2)
	}
}

//accepts strings and runs a for loop printing string
func concurPass(s string) {
	i := 0
	for i < 10 {
		fmt.Println(i, s)
		i++
		time.Sleep(time.Millisecond * 500)
	}
}

//instantiates go routines
func callConcurExample() {
	//function content from https://www.youtube.com/watch?v=LvgVSSpwND8&t=340s "concurrency in go" by Jake wright
	var wg sync.WaitGroup //waitGroup variable is able to communicate with cpu saying
	wg.Add(1)             //"wait for these processes to finish before exiting program"
	// go concurPass("Dog")
	// go concurPass("Bird")
	go func() { //anonymous function
		//this runs in the background so how can we pass the value
		concurExample()
		go concurPass("Dog")
		concurPass("Bird")
		wg.Done() //this makes the wg decrement
	}()
	wg.Wait() //makes program wait

}

func channelsYoutubeTut() {
	out1 := make(chan string)
	go processWithChannels("order", out1)
	var themsg string

	for i := 0; true; i++ {
		msg, open := <-out1 //reading data out of channel ,this will block until it actually gets data to read
		themsg += msg + " "

		if !open {
			break
		}
		fmt.Println(i, msg)
	}
	fmt.Println(themsg)

	// we can also do it this way
	for msg := range out1 {
		msg = <-out1
		fmt.Println(msg)
	}
}

func processWithChannels(item string, out chan string) {
	defer close(out)
	defer fmt.Println("this is the end of the processWithChannel function")
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second / 2)
		out <- item //passes item to the channel

	}
	//close(out) //manually close channel when done
}

//-----tour go
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

//
func usingChannels() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int, 3)  // make a channel
	go sum(s[:len(s)/2], c) //pass in first part of sice
	go sum(s[len(s)/2:], c) //pass in second part of slice
	x, y := <-c, <-c        // receive from c

	fmt.Println(x, y, x+y)
}

//-----tour go


// for val := range ch {
// 	val = <- ch
// 	fmt.Println(val)
// }