package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestPing(t *testing.T)  {
	url := "http://localhost:5656/api/ping"
	request, err := http.NewRequest("GET", url, nil)


	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}
	body, _ := ioutil.ReadAll(res.Body)
	 responseBody := strings.TrimSpace(string(body))
	 observedStatusCode := res.StatusCode

	fmt.Println("responseBody-> " , responseBody)
	fmt.Println("observedStatusCode-> ", observedStatusCode)
	 DisplayTestCaseResults("ping", responseBody, observedStatusCode, t)
}





func DisplayTestCaseResults(functionalityName string, responseBody string, observedStatusCode int, t *testing.T) {
		fmt.Println("in DisplayTestCaseResults, ", observedStatusCode)
		if observedStatusCode == 200 {
			t.Logf("-> %s, -> %v, -> %s ",  responseBody, observedStatusCode, functionalityName)
		} else {
			t.Errorf("%s, %v, %s ",  responseBody, observedStatusCode, functionalityName)

		}
	}