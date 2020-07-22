package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileName:="sample.txt"

	content, err := ioutil.ReadFile(fileName)

	checkError(err)

	res:= string (content)
	fmt.Println(res)

}

func checkError(err error)  {
	if err != nil {
		panic(err)
	}
}