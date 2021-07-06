package main

import (
	"fmt"
	"time"
)

func concurExample() {
	i := 0
	for {
		fmt.Println(i)
		i++
		time.Sleep(time.Millisecond * 500)
	}
}

func callConcurExample() {
	concurExample()
	concurExample()
}
