package main

import (
	"fmt"
)

func test() {
	fmt.Println("Print")
}

func takeInInt(c int) {
	fmt.Println(c)
}

var ofTest = test

func main() {
	fmt.Println("Hello World")
	x := test
	x()
	ofTest()
	x1 := takeInInt
	x1(4)

}
