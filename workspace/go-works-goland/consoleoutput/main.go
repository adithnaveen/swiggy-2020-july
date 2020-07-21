package main

import "fmt"

func  main()  {

	str1 := "This is string 1"
	str2 := "This is second string"
	str3 := "hello there"
	aNumber := 33
	isTrue := true

	// if you dont want the return parameter the put _
	strLength, _ := fmt.Println(str1, str2, str3);

	//if err == nil {
	//	fmt.Println("The length of the string: ", strLength);
	//}

	fmt.Println("The length of the string: ", strLength)

	fmt.Printf("Value of aNumber %v\n", aNumber)
	fmt.Printf("Value of Boolean %v\n", isTrue)
	fmt.Printf("Value of aNumber as float: %.2f\n", float64(aNumber))

	fmt.Printf("Data Types : %T, %T, %T, %T and %T\n",
		str1, str2, str3, aNumber, isTrue);

	myString := fmt.Sprintf("Data Types : %T, %T, %T, %T and %T",
		str1, str2, str3, aNumber, isTrue)

	fmt.Println(myString)



}