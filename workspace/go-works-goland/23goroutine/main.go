package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := [] string {
		"https://golang.org/", // string
		"https://jsonplaceholder.typicode.com/posts", // json
		"https://httpbin.org/xml",  // xml
	}

	var wg sync.WaitGroup
	for _, url := range  urls{
		wg.Add(1)
		go func(url string) {
			returnType(url) // async
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func returnType(url string)  {
	resp, err:= http.Get(url);
	checkError(err)
	defer resp.Body.Close() ;

	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)

}

func checkError(err error)  {
	if err != nil {
		panic(err)
	}
}
