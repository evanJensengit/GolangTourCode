package main

import (
	"fmt"
	"sync"
	"time"
)

func concurExample() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
		time.Sleep(time.Millisecond * 500)
	}
}

func concurPass(s string) {
	i := 0
	for i < 10 {
		fmt.Println(i, s)
		i++
		time.Sleep(time.Millisecond * 500)
	}
}
func callConcurExample() {
	var wg sync.WaitGroup
	wg.Add(1)
	go concurPass("Dog")
	go concurPass("Bird")
	go func() {
		concurExample()
		wg.Done()
	}()
	wg.Wait()

	time.Sleep(time.Microsecond * 10000)
}
