package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type  Posts struct {
	Title string
	UserId int
}

func main() {
   url:="https://jsonplaceholder.typicode.com/posts"
	content := ContentFromServer(url)
	fmt.Println(PostFromJson(content))
}

func PostFromJson(content string)  [] Posts {
	posts := make([] Posts, 0, 20)
	decoder:= json.NewDecoder(strings.NewReader(content))

	_, err := decoder.Token()
	checkError(err)

	var post Posts

	for decoder.More() {
		err:= decoder.Decode(&post)
		checkError(err)
		posts = append(posts, post)
	}

	return posts
}

func ContentFromServer(url string) string  {
	resp, err:= http.Get(url)
	checkError(err)
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	return string(bytes)
}

func checkError(err error)  {
	if err != nil {
		panic(err)
	}
}
