package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
	r := Rectangle{5, 3}
	q := &Square{6}
	shapes := []Shaper{r, q}
	for _, s := range shapes {
		fmt.Printf("类型 = %T, 面积 = %f\n", s, s.Area())
		switch t := s.(type){
		case Rectangle:
			fmt.Printf("Type = %T, area = %f\n", t, t.Area())
		case *Square:
			fmt.Printf("Type = %T, area = %f\n", t, t.Area())
		default:
			fmt.Printf("Unknown type")
		}
	}

	var sh Shaper = q
	if v, ok := sh.(*Square); ok{
		fmt.Printf("v = %T\n", v)
	}
}
