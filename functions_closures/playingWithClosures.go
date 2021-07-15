package main

import (
	"fmt"
)

func test() {
	fmt.Println("Print")
}

var ofTest = test

func main() {
	fmt.Println("Hello World")
	x := test
	x()
	ofTest()

}
