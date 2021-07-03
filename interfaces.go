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
	fmt.Printf("Name: %s \n  Current state: %s, \n  Current location: %s \n\n", p.name, p.state, p.location)
}

type Object interface {
	changeState()
	changeLocation() bool
	describe()
}

func main() {
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

}
