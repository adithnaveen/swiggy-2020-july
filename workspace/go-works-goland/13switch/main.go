package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(6)+1
	result:=""
	// first day is 1
	// last day is 7
	switch  dow{
		case 1:
			result = "its sunday "
		case 7:
			result ="its Saturday"
		default:
			result ="its a week day"
	}
	fmt.Println(dow)
	fmt.Println(result)

	//-------------

	x:=-30
	switch  {
	case x<0 :
		fmt.Println("The value is less than 0 ")
		fallthrough
	case x==0 :
		fmt.Println("The value is ==0 ")
	default:
		fmt.Println("The value is more than zero")
	}


}
