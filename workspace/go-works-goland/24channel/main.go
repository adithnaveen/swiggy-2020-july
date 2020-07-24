package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	c:= make(chan string)
	c1:= make(chan string)

	go DoMyWork("Hemanth", c)
	go DoMyWork1("Harry", c1)

	for msg:= range  c {
		fmt.Println(msg)
	}
	for msg:= range  c1 {
		fmt.Println(msg)
	}

}

func DoMyWork(name string, c chan string){
	for i:=1; i<=4; i++ {
		c <- "Doing Work of "+ name + " executed " + strconv.Itoa( i)
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}


func DoMyWork1(name string, c chan string){
	for i:=1; i<=10; i++ {
		c <- "Doing Work of "+ name + " executed " + strconv.Itoa( i)
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}

func IsClosed (ch <- chan string)   bool  {
	select {
		case <- ch :
				return true
	default:

	}
	return false
}

