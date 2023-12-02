package main

import "fmt"

type poly interface {
	area() int
}

type rect struct {
	length  int
	breadth int
}

type square struct {
	side int
}

func (r rect) area() int {
	return r.length * r.breadth
}

func (s square) area() int {
	return s.side * s.side
}
func which(r, s poly) bool {
	return r.area() > s.area()

}

func main() {
	fmt.Println("interfaces tutorial")
	rect_obj := rect{2, 3}
	sqr_obj := square{2}
	print(which(rect_obj, sqr_obj))

}
