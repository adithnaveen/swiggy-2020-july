package main

import "fmt"

var isConnected bool = false

func main() {
	fmt.Printf("Connection Open : %v\n", isConnected)
	businessLogic()
	fmt.Printf("Connection Open : %v\n", isConnected)
}

func businessLogic() {
	defer disConnect()
	connect()
	fmt.Println("DB Connection Status businessLogic : ", isConnected)
	fmt.Println("Your Business Logic Goes Her ")
}
func connect(){
	isConnected = true
	fmt.Println("DB Connection is Open")
}
func disConnect() {
	isConnected = false
	fmt.Println("DB Connection Close ")
}