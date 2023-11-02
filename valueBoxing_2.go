package main

import "fmt"

type Box interface {
	dimension() int
}

type Carton struct {
	height, width, depth int
}

func (c *Carton) dimension() int {
	return c.height * c.width * c.depth
}

func main() {
	var c Carton
	fmt.Println(c)

	cc := &Carton{
		1, 2, 3,
	}
	fmt.Println("struct", cc)
	fmt.Println("struct", cc.dimension())

	//value of type *Carton is boxed into an interface value of type Box
	var b Box = &Carton{1, 2, 3}
	fmt.Println(b)

	// i is a blank interface value, value of type *Carton is boxed into an blank interface
	var i interface{} = &Carton{1, 2, 3}
	fmt.Println(i)

	i = b
	fmt.Println(i)

	fmt.Println(b.dimension())

}
