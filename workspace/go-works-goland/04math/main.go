package main

import (
	"fmt"
	"math"
	"math/big"
)

func main()  {
	i1, i2, i3 := 10, 20, 30
	intSum := i1 + i2 + i3

	fmt.Println("Sum of Number " , intSum)

	f1, f2, f3 := 23.4, 43.5, 677.6
	floatSum := f1 + f2 + f3

	fmt.Println("Sum of floats ", floatSum)

	var b1, b2, b3, bigSum big.Float
	b1.SetFloat64(1.1)
	b2.SetFloat64(2.2)
	b3.SetFloat64(3.6)

	bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
	fmt.Printf("BigSum value : %.2f\n", &bigSum)

	circleRadius:=14.5
	circumference := circleRadius * math.Pi
	fmt.Printf("Circumference: %.2f", circumference)



}
