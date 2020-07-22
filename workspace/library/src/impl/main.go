package main

import (
	"fmt"
	"stringutil"
)

func main() {
	n1, l1 := stringutil.FullName("Naveen", "Kumar")
	fmt.Printf("FullName: %v, Number of charss: %v\n", n1, l1)

	n2, l2 := stringutil.FullNameNakedReturn("Kanchan", "Kumar")
	fmt.Printf("FullName: %v, Number of charss: %v\n", n2, l2)
}
