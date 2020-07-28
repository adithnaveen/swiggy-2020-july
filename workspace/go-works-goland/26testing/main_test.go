// +build unit

package main

import (
	"io/ioutil"
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
	for i:=0; i<b.N; i++ {
		Add(i, i)
	}
}








