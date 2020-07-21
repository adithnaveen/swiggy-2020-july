package main

import "fmt"

func  main()  {
	var names [5] string

	names[0] = "Rohan"
	names[1] = "Hemanth"
	names[2] = "Milan"
	names[3] = "Adithya"
	names[4] = "Karthik"

	fmt.Println(names)

	fmt.Println(names[0])

	var numbers =  [3] int {1,2,3};

	fmt.Println(numbers)
	fmt.Println(numbers[0])
	fmt.Println(len(numbers))

	// you cannot add any more values
	newNumber := 3
	numbers[newNumber] = 34
	fmt.Println(numbers)
}
