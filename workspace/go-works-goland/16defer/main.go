package main

import (
	"fmt"
)

func main()  {
	//if you put defer this will be execute last
	defer fmt.Println("Connection Close")
	fmt.Println("Connection Open")
	//// some problem here
	//panic("SomeError")

	myFunc()

	defer fmt.Println("Statement1 ")
	defer fmt.Println("Statement2 ")

	fmt.Println("Do the work ")
}

func myFunc()  {
	defer fmt.Println("I'm defer in myFunc()")
	fmt.Println("Not Defer in myFunc")
}
