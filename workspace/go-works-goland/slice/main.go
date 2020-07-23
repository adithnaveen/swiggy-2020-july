package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = [] string {"Red", "Blue", "Green"}

	fmt.Println(colors)
	colors = append(colors, "Yellow")
	fmt.Println(colors)


	// leave out first element
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	// to remove last element
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	colors =  colors[:len(colors)-1]
	fmt.Println(colors)

	ints := make([]int, 5, 5)
	fmt.Println("Initial Capacity: " , cap(ints))
	ints[0] = 33;
	ints[1] = 435;
	ints[2] = 66;
	ints[3] = 77;
	ints[4] = 89;
	fmt.Println("After adding 5 elements Capacity: " , cap(ints))
	fmt.Println(ints)
	ints = append(ints, 100)
	fmt.Println(ints)
	fmt.Println("After Adding 10  Capacity: " , cap(ints))

	// asc order
	sort.Ints(ints)
	fmt.Println("In Ascending Order: ", ints)

	// in descending order []int
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	
	fmt.Println("In Descending Order: ", ints)
}
