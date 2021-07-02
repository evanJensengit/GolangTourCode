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
	fmt.Println("RANGE LOOP")
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
	//--FUNCTIONS AS VALUES. More types 24/27
	//functions are values also and you can pass functions to other functions

	//hypot is a function that takes in two float64 variables and returns the square root of x^2 and y^2
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	//compute takes in a function
	//varFLoat := compute(hypot(6,2))
	//^this doesnt work because the compute func gets passed a func variable. If we were to
	//pass it the function "hypot" with the values i.e hypot(6,2) we would essentially be passing
	//a float64.
	//Since we pass it the empty hypot fcn it is of func type and waits to jump into hypot until it has values
	//given to it
	varFloat := compute(hypot)
	fmt.Println(varFloat)

	//CLOSURES
	//closures are function values that access variables outside of the body of their function
	//that they keep track of or are "bound to" outside when they
	//are used. the adder function returns a closure.

	//CLOSURE EXERCISE
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	//METHODS
	usePersonMethods()
	//method has special receiver argument
	//method is a function with a receiver argument

	//value methods take only the value with the specific type, pointer
	//methods take the value or a pointer

	//always use a pointer receiver

	//interfaces have method signatures. objects that use the method implement the interface
}

//this function gets passed a function variable
func compute(fn func(float64, float64) float64) float64 {

	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
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

type Vertex struct {
	X, Y float64
}

//method Abs() receivers a Vertex. So a Vertex can call Abs()
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//So I can create a struct
type Person struct {
	name        string
	theirHabits []Habit
}

func (p Person) details() {
	fmt.Print("Hi my name is", p.name, "and I like to ")

	for _, e := range p.theirHabits {
		if !e.isEmpty() {
			fmt.Printf("%v,", e)
		}
	}
}
func (h Habit) isEmpty() bool {
	if h.name == "" {
		return true
	}
	return false
}

type Habit struct {
	name string
}

func usePersonMethods() {
	PersonAHabits := make([]Habit, 10)
	habit1 := "Working"

	aHabit := Habit{habit1}

	PersonAHabits[0] = aHabit
	PersonA := Person{"Bob", PersonAHabits}
	PersonA.details()

}
