package main

import "fmt"

func main() {
	/*
		an example which shows how a blank interface value is used to box values of any non-interface type.
	*/

	var box interface{} //interface type variable

	fmt.Println(box) //In Go, the zero values of any interface type are represented by the predeclared (nil)

	/*Nothing is boxed in a nil interface value. Assigning an untyped nil to an interface value will clear the dynamic value boxed in the interface value.*/

	n := 5 //non-interface type variable - dynamic value of 5 with dynamic type of int

	box = n //boxing 'n' into interface value

	fmt.Println(box)

	s := "Boxed String, a dynamic value " //non-interface type variable - dynamic value with dynamic type of string

	box = s

	fmt.Println(box)
}
