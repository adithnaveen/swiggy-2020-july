package main

import (
	"fmt"
	"net/http"
)


//type Handler interface {
//	ServeHTTP(ResponseWriter, *Request)
//}

type Hello struct {

}

// fmt.println("hi")
/// fmt.Fprint(writer, "hi");
func (h Hello) ServeHTTP(writer http.ResponseWriter, req *http.Request)  {
	fmt.Fprint(writer, "<h2> Hello Welcome from Go Lang</h2>")
}


func main() {
	var h Hello
	err:= http.ListenAndServe("127.0.0.1:7788", h)
	checkError(err)
}



func checkError(err error)  {
	if err != nil {
		panic(err)
	}
}
