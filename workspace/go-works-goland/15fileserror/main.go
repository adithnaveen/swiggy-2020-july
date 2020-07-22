package main

import (
	"fmt"
	"os"
)

// simple HTTP error
type HttpError struct {
	status int
	method string
}
type SSLError struct {
	statusCode int
	message string
}
// HttpError implements Error method
func (httpError *HttpError) Error() string {
	return fmt.Sprintf("Something went wrong with the %v request. Server returned %v status.",
		httpError.method, httpError.status)
}

func (sslError *SSLError) Error() string {
	return fmt.Sprintf("Some thing went wrong %v, with Code %v", sslError.message, sslError.statusCode)
}

// mock function: sends HTTP request
func GetUserEmail(userId int) (string, error) {

	if userId <100 {
		return "email Found: naveen@test.com", nil
	}
	// request failed, return an error
	//  &HttpError{403, "GET"}
	return "", &HttpError{401, "GET"} // return is error
}

func main() {

	// get user email address
	if email, err := GetUserEmail(101); err != nil {
		fmt.Println(err) // print error

		// if user is unauthorized, perform session cleaning
		if errVal := err.(*HttpError); errVal.status == 403 {
			fmt.Println("Performing session cleaning...")
		}
	} else {
		// do something with the `email`
		fmt.Println("User email is", email)
	}
}



func main1() {
	file, err := os.Open("sample.txt")

	if err == nil {
		fmt.Println(file)
	} else  {
		fmt.Println(err.Error())
	}



}
