package main

import (
	"fmt"
)

//digital object like a "window" or an email
// type softwareObject struct {
// 	platform string
// 	state string

// }

// func (s softwareObject) changeState(state string) {

// }

// func (s softwareObject) changeLocation(loc string) {

// }

//physical object represents any object in the material world
type physicalObject struct {
	stateOptions, locationOptions []string
	state, location, name         string
	m                             map[string]string
}

//gets passed a string of that state the object is changing to
func (p *physicalObject) changeState(state string) {
	p.state = state
	return
}

func (p *physicalObject) changeLocation(loc string) bool {
	for _, e := range p.locationOptions {
		if e == loc {
			p.location = loc
			p.changeState(p.m[loc])
			return true
		}
	}
	fmt.Println("Enter valid location")
	return false

}

func (p *physicalObject) describe() {
	if p == nil {
		fmt.Println("<nil>")
		return
	}

	fmt.Printf("Name: %s \n  Current state: %s, \n  Current location: %s \n\n", p.name, p.state, p.location)
}

type Object interface {
	changeState(string)
	changeLocation(string) bool
	describe()
}

//for 12/26 methods tour.golang.org
type tempStruct struct {
	S string
}

type tempInter interface {
	M()
}

func (t *tempStruct) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}
func logicPlatformHardCodeExample() {
	mot1 := physicalObject{
		stateOptions:    []string{"Shop", "Road"},
		locationOptions: []string{"Home", "Work", "Shop"},
		name:            "Motorcycle",
		state:           "On road",
		location:        "Home",
		m:               make(map[string]string),
	}

	mot1.m["Home"] = "Road"
	mot1.m["Work"] = "Road"
	mot1.m["Shop"] = "Shop"

	mot1.describe()
	mot1.changeLocation("Shop")
	mot1.describe()

	// var o Object
	// var t1 *physicalObject
	// o = t1

	var tempI Object
	mot2 := &mot1
	tempI = mot2
	fmt.Println("Describing the interface Object that has the underlying type of physicalobject")
	fmt.Println("The variable I am using to get the underlying physical object is actually a reference to the memory address of the actual physicalObject")
	tempI.describe()

	var tempI2 tempInter
	var tempS *tempStruct
	tempI2 = tempS
	tempI2.M()

	var tempO Object
	var tempP *physicalObject
	tempO = tempP
	tempO.describe()

	//here I do a type assertion
	tAssert, ok := tempO.(*physicalObject)
	fmt.Println(tAssert, ok)
	checkType(34)
	checkType(tempO)

}

//type switches
func checkType(i interface{}) {
	switch var1 := i.(type) {
	case *physicalObject:
		fmt.Println(var1, " is Physical object ")
	case float64:
		fmt.Println(var1, " is Float object ")
	case string:
		fmt.Println(var1, " is string")
	case int:
		fmt.Println(var1, " is int")
	default:
		fmt.Println(var1, " dunno the type")
	}

}
