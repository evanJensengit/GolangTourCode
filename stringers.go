package main

import (
	"fmt"
)

type Human struct {
	name   string
	weight int
	sex    string
}

func (h Human) String() string {
	return fmt.Sprintf("%v is %v lbs and is a %v ", h.name, h.weight, h.sex)
}

func testingStringer() {

	h := Human{
		name:   "Joe",
		weight: 400,
		sex:    "Male",
	}
	fmt.Println(h)
}
