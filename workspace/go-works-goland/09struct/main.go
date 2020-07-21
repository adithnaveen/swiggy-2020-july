package main

import "fmt"

type Dog struct {
	 // declare the variables if it is starting with Upper case then public

	Breed string
	Weight int
}

func main() {
	dog1  := Dog{Breed: "Beagle", Weight: 3}
	fmt.Println(dog1)

	fmt.Printf("Breed : %v\nWeight:%v\n", dog1.Breed, dog1.Weight)

}
