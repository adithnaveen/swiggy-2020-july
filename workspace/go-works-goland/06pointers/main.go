package main

import "fmt"

func main() {
	var p *int

	if p!= nil {
		fmt.Println("Value of P : ", *p);
	} else {
		fmt.Println("Sorry p is nil ")
	}

	var v int = 44
	p = &v

	if p!= nil {
		fmt.Println("Value of P : ", *p);
	} else {
		fmt.Println("Sorry p is nil ")
	}

	var value1 float64 = 33.2
	pointerF := &value1
	fmt.Println("Value 1 ", *pointerF)

	*pointerF = *pointerF / 4
	fmt.Println("Value1 P  ", *pointerF)
	fmt.Println("Value1 ", value1)

}