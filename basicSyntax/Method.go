package main

import "fmt"

//Rect - struct definition
type Rect struct {
	width, height int
}

//Rect area() method
func (r Rect) area() int {
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area() // call method
	fmt.Println(area)
}