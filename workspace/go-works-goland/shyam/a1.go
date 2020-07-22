package main

import (
	"bufio"
	"fmt"
	"os"
)

type Animal interface {
	eat() string
	sleep() int
	breath() string
}

func (e Emp) eat() string {
	return "Rice,fruits etc"
}

func (e Emp) sleep() int {
	return 6
}

func (e Emp) breath() string {
	return "Nasal Respiratory System"
}


type Emp struct {
	empid string
	empname string
	salary int
	address *Address
}

type Address struct {
	hno string
	street string
	city string
	pin int
}


func main(){
	reader := bufio.NewReader(os.Stdin)
	//emps := []*Emp{}
	emps := make([]*Emp,0,5)
	fmt.Println(cap(emps))
	i:=0

	for {
		//emp := new(Emp)
		emp := &Emp{}

		fmt.Println("Enter Empid")
		emp.empid ,_ = reader.ReadString('\n')

		fmt.Println("Enter Employee Name")
		emp.empname ,_ = reader.ReadString('\n')

		fmt.Println("Enter Employee Salary")
		fmt.Scan(&emp.salary)

		adr := &Address{}

		fmt.Println("Enter house number")
		adr.hno ,_ = reader.ReadString('\n')

		fmt.Println("Enter Street name")
		adr.street ,_ = reader.ReadString('\n')

		fmt.Println("Enter Employee City Name")
		adr.city ,_ = reader.ReadString('\n')

		fmt.Println("Enter pin number")
		//adr.pin ,_ = reader.ReadString('\n')
		fmt.Scan(&adr.pin)

		emp.address = adr

		emps = append(emps,emp)
		//if i%5 == 0{
		//	emps = append(emps,emp)
		//} else {
		//	emps[i] = emp
		//}

		fmt.Println("To enter more emplyees : (0/1) ")
		response := 0
		fmt.Scan(&response)
		if response == 0 {
			break
		}
		i++
	}

	for i=0;i<len(emps);i++ {

		fmt.Println("Empid = ",emps[i].empid,"Empname = ",emps[i].empname,"Salary = ",emps[i].salary)

		fmt.Println("House Number = ",emps[i].address.hno,"Street name = ",emps[i].address.street,"City Name = ",emps[i].address.city,"Pin Number = ",emps[i].address.pin)

		fmt.Println("Eating Habit : ",emps[i].eat(),"\nAvg sleeping Duration : ",emps[i].sleep(),"\nBreathing Process : ",emps[i].breath())

	}

	fmt.Println("At the last capacity :",cap(emps))

}