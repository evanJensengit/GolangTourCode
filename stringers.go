package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Human struct {
	name   string
	weight int
	sex    string
}

type rot13Reader struct {
	r io.Reader
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

func testingReader() {
	//NewReader is a method that creates a reader
	//a Reader interface can utilize read method
	r := strings.NewReader("this is something for reader to read reader reads the reader")
	fmt.Printf("Type of %v is %T", *r, *r)
	s1 := make([]byte, 8)
	n, err := r.Read(s1)
	fmt.Printf("%q %v", s1[:n], err)
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	l, err := r.r.Read(b)
	fmt.Println("In rot13Reader")
	for i := 0; i < l; i++ {
		fmt.Println("In for loop")
		b[i] += 4
		i++
	}
	return l, err
}

//apparently it is common to wrap a reader with another reader
func testingReaderOfReaderWithIO() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r) //here reader is implicitely called
}

func testingReaderOfReaderNoIO() {
	s := strings.NewReader("Something to be read")

	b := make([]byte, 12)
	s.Read(b)
	fmt.Printf("%q\n\n", b)
	r := rot13Reader{s}
	fmt.Printf("%q", r)
	r.Read(b)
	fmt.Printf("%q\n\n", b)

}
