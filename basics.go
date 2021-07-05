package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	testingStringer()
} //END MAIN
//reading into slices

func somePointers() {
	i := 12
	b := &i
	fmt.Println(b)
	var i1 byte = 5
	fmt.Println(i1)
	i1 += 14
	fmt.Println(i1)

}

//READERS
//what readers do is they read in data
func slicesAndReader() {
	sb := make([]byte, 8)
	read := strings.NewReader("Sup reader boi!")
	nb, err := read.Read(sb)
	fmt.Println("Number of bytes populated", nb)
	fmt.Println("What the bytes represent:", sb[:nb])
	fmt.Println("Error= ", err)

}

func usingRange() {
	slice := make([]int, 10)
	for i := range slice {
		slice[i] = i
	}
	fmt.Println(slice)
}

//what is the difference between declaring a type and declaring a struct?
type F float64

//practice type for stringers section
type arr [5]float64

type arrNoString [5]float64

func (a arr) String() string {
	return fmt.Sprintf("|  %g|%g|%g|%g|%g  |", a[0], a[1], a[2], a[3], a[4])
}

//What does an interface do?
//interfaces hold any values that implement the methods that are listed in the interface
//interface values have underlying types that they represent
// interface I represents any object that utilizes the details() method call
type I interface {
	details()
}

type P interface {
}

type H interface {
	isEmpty() bool
}

//basically if there are two different types that use the same method
//then we can use an interface to generalize those two different types into one
//broader type
//we can set an interface type equal to another type and then use
//methods defined in the interface on that data type

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

//STRINGERS
//with stringers I can create a method with the receiver type and then the method
//String() string and it will alter the print statement

type Robot struct {
	ID   int
	name string
	mark int
}

func (r *Robot) details() {
	fmt.Print("Hi my name is ", r.name, ", My ID number is ", r.ID, ".\n")
	fmt.Println("I am a mark", r.mark, "robot\n")
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

func (p *Person) details() {
	fmt.Print("Hi I am a human my name is ", p.name, " and I like to ")

	for _, e := range p.theirHabits {
		if !e.isEmpty() {
			fmt.Print(e.name)
		}
	}
	fmt.Println("\n")
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

func goingThroughTourGo() {
	fmt.Println("HELLO")
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

	//SWITCH STATEMENT
	//can also use strings
	x := 4

	switch x {
	case 1:
		print("1")
	case 4:
		print("4") //prints 4

	}

	//with no parameter to switch statement it evaluates booleans
	switch {
	case x > 5:
		print(" x > 5")
	case x == 5:
		print(" x == 5")

	case x < 5:
		print(" x < 5")
		//prints x < 5

	}
	//^ this is a good way to do long if else statements

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
		fmt.Print(f(), " ")
	}
	fmt.Println()

	//METHODS
	usePersonMethods()
	//method has special receiver argument
	//method is a function with a receiver argument
	//we can only declare a receiver with types we have defined in the package,
	//not types defined outside the package (like int)
	//value methods take only the value with the specific type, pointer
	//methods take the value or a pointer

	//always use a pointer receiver since then we can directly modify the variables value
	//all methods on a given type need to have the same type of receiver (either value or pointer)

	//interfaces have method signatures. objects that use the method implement the interface

	//BENEFIT of using a method is that we can call it with a reference or a pointer or a value
	//whereas with a function that takes in a value you can only call that function with the specific
	//type of value (pointer or value - not both)
	h1 := Habit{"Playing with cats"}
	var i2 H = h1
	fmt.Printf("%T", i2)

	var inter1 P = F(math.Pi)

	fmt.Printf("%T", inter1)
	fmt.Println("HI_")

	PersonAHabits := make([]Habit, 10)
	habit1 := "Working"

	aHabit := Habit{habit1}

	PersonAHabits[0] = aHabit
	PersonA := Person{"Bob", PersonAHabits}
	PersonA.details()
	var i3 I = &PersonA
	i3.details()
	i4 := make([]I, 2)
	i4[0] = &PersonA
	i4[1] = &Robot{name: "Joe",
		ID:   234,
		mark: 12,
	}
	fmt.Println("------PRINTING THE INTERFACE-------")
	for _, e := range i4 {
		e.details()
	}

	//empty interface
	//type assertion with interface what does that do?

	//STRINGERS
	a := arr{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", a) //I can change how this is printed with Stringers
	b1 := arrNoString{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", b1)
	fmt.Println("HELLO")
	slicesAndReader()
	usingRange()
	somePointers()
}
