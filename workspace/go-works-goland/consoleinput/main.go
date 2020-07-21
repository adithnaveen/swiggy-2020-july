package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {

	//var s string
	// if you have a long string like (hi how are you) then only hi is considered
	//fmt.Scanln(&s)
	//fmt.Println(s)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Your Name")
	name, _ := reader.ReadString('\n')
	fmt.Println("Your Name is " , name)

	// accept salary -> float
	fmt.Println("Enter Your Salary")
	strSalary, _ := reader.ReadString('\n')
	salary, _  := strconv.ParseFloat(strings.TrimSpace(strSalary), 64);
	fmt.Println("Your Salary is " , salary)

	// accept number of months since joined -> int
	fmt.Println("Enter Number of Months")
	strMonths, _ := reader.ReadString('\n')
	months, _  := strconv.ParseInt(strings.TrimSpace(strMonths), 10, 64)
	fmt.Println("Number Of Months You Worked: " , months)









	// calculate salary * noOfMonths



}