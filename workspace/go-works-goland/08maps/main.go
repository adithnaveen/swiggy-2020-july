package main

import "fmt"

func main() {
	states := make(map[string] string)
	fmt.Println(states)

	states["KA"] = "BLR"
	states["TN"] = "CHN"
	states["AP"] = "HYD"

	fmt.Println(states)

	states["KA"] = "MYS"
	fmt.Println(states)

	city:= states["KA"]
	fmt.Println("Capital of Karnataka :", city)

	for k, v := range states{
		fmt.Printf("%v, %v \n", k, v)
	}
	delete(states, "KA")
	fmt.Println()

	for k, v := range states{
		fmt.Printf("%v, %v \n", k, v)
	}

	// to print only the keys
	keys := make([] string, len(states))
	i:=0

	for k:= range  states {
		keys[i] = k
		i++
	}

	fmt.Println(keys)

	fmt.Println("Keys are : ")
	// to show with the for loop
	for i:= range  keys {
		fmt.Println(keys[i])
	}

}
