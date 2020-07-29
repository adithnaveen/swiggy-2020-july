
package main

/* you can enable tags with (put this as first statement)
// +build unit

 */
import (
	//"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type AddResult struct {
	x int
	y int
	expected int
}

var AddResults =[] AddResult {
	{10,20, 30},
	{100,200,300},
	{1,2,3},
}

// go test
func TestAdd(t *testing.T)  {
	for _, test := range AddResults {
		result :=Add(test.x, test.y)
		if result != test.expected {
			t.Fatalf("Expected Result Not Found... ")
		}
	}
}

func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/test.data")
	if err != nil {
		t.Fatalf("File Not Found : %v", err)
	}
	expected:="Hello World"

	if string(data) != expected {
		t.Fatalf("Data Mismatch from file... ")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(i, i)
	}
}


	func TestHttpRequest(t *testing.T) {

		handler := func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "{\"status\" : \"good\"}")
		}

		// GET/DELETE - nil
		// POST/PUT - JSON
		req := httptest.NewRequest("GET", "https://jsonplasdfsdfceholder.typcode.com/todos/1", nil)

		w := httptest.NewRecorder()

		handler(w, req)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		if 200 != resp.StatusCode {
			t.Fatal("status code not okay ")
		}
	}


