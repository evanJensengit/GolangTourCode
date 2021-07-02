package main

import (
	"fmt"
	"math"
)

func main() {
	//SLICES
	var s []int = []int{2, 3, 4, 5}
	b := []int{2, 3, 4, 5}
	fmt.Println(s, cap(s))
	fmt.Println(len(s))
	c := append(b, 12)
	b = append(c, 13)
	fmt.Println(c, b, "\n")

	//can make a slice without doing the slice literal
	slice := make([]int, 10)
	slice[0] = 1
	fmt.Println(slice)
	//slices can contain any type

	//RANGE LOOP
	for a, b := range slice {
		fmt.Println(a, b)
	}
	for _, b := range slice {
		fmt.Print(b)
	}

	fmt.Println()
	for b := range slice {
		fmt.Print(b)
	}

	//FOR LOOP
	for i := 12; i > 0; i-- {
		fmt.Println(i)
	}

	i := 0
	//WHILE LOOP
	for i < 10 {
		//this is a while loop
		fmt.Println(i, " Balls")
		i++
	}
	//exercise for tour of go lang SLICES
	dx := 21
	dy := 13

	slice1 := make([][]uint8, dx)
	fmt.Println(len(slice1))
	//index, element

	for i := range slice1 {
		slice1[i] = make([]uint8, dy)
	}
	for h := range slice1 {
		for v := range slice1[h] {
			slice1[h][v] = uint8(h * v)
		}
	}

	var m map[string]int
	m = make(map[string]int)
	m["a"] = 4

	//functions are values also and you can pass functions to other functions

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	//method has special receiver argument
	//method = function with receiver argument

	//value methods take only the value with the specific type, pointer
	//methods take the value or a pointer

	//always use a pointer receiver

	//interfaces have method signatures. objects that use the method implement the interface
}

func fibonacci() func() int {
	x := 0
	y := 1
	sum := 0
	xturn := true
	return func() int {
		if xturn {
			xturn = false
			sum = x + sum
			x = sum
			return sum
		}
		xturn = true
		sum = y + sum
		y = sum
		return sum
	}
}

func compute(fn func(float64, float64) float64) float64 {

	return fn(3, 4)
}
