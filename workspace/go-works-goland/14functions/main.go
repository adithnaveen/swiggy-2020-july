package main

import (
	"fmt"
	"strings"
)

func main()  {
	DoSomeWork()
	fmt.Println(AddValues(10, 20))
	fmt.Println(AddAllValues(12,23,34,45,6,7,8))
	fullName, length := FullName("Harry", "Peter")
	fmt.Printf("Full Name is: %s and length is: %v\n", fullName, length)

	full, length := FullNameNakedReturn("Steve", "Jobs")
	fmt.Printf("Full Name is: %s and length is: %v\n", full, length)
}

func DoSomeWork() {
	fmt.Println("Doing Your Work... ")

}

func  AddValues ( value1 , value2 int) int   {
	return value1 + value2
}
//in go overloading is not supported
//func  AddValues ( value1 , value2 int) int   {
//	return value1 + value2
//}

func AddAllValues(values ... int) int  {
	sum:=0

	for i:= range  values {
		sum += values[i]
	}
	return  sum
}

// in go you can have multi return
// the functions first argument returns full name
// second argument returns lenght
func FullName (f, l string) (string, int) {
	full:= f + " " + l
	length := len(full)
	return full, length
}

// you can name the return variable
// in this kind of functions we are not returning
// these functions are called naked function since they dont
// return anything inside
func FullNameNakedReturn(f, l string) (full string, length int)  {
	// since the variable decl is already taken care
	// := is not allowed
	// use only =
	full = f + " " + l
	strings.ToUpper()
	length = len(full)
	return
}
















