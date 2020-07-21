package main

import "fmt"

// when you have interface then there is not relationship
// between interface and struct
type  Animal interface {
	// method name
	// return type
	Speak() string
}

// this will not work because Speak takes an argument
//type  Animal interface {
//	Speak(string) string
//}


// function decl
// func functionName (params) (return value/s) {}

func (e Emp) Speak() string {
	return "Hello... Hello..."
}
func (d Dog) Speak() string {
	return "Wof.. Wof... "
}

func (c Cat) Speak() string {
	return "Meow.. Meow... "
}

// structure or noun
type Dog struct {}
type Emp struct {}
type Cat struct {}



func main() {
	dog1 := Animal(Dog{})
	fmt.Println(dog1)
	fmt.Println(dog1.Speak())


	animals := [] Animal{Emp{}, Dog{}, Cat{}}

	for _, animal := range  animals {
		fmt.Println(animal.Speak())
	}


}