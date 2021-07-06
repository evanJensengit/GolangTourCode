package main

import (
	"fmt"
	"image"
	"image/color"
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
	//to enter the Read function with *rot13Reader receiver I need to call r.Read
	r.Read(b)
	fmt.Printf("%q\n\n", b)

}

func testingImages() {
	m := image.NewRGBA(image.Rect(0, 40, 100, 100))
	fmt.Printf("%T", m)
	fmt.Println(m.Bounds())
	fmt.Println(m.At(12, 0).RGBA())
	fmt.Println(m)
	m.ColorModel()
}

type Image struct {
	w, h int
}

func (im Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.w, im.h)
}

func (im Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}

}
func imageWorks() {
	m := Image{64, 128}
	fmt.Println(m)
	//pic.ShowImage(m)
}
