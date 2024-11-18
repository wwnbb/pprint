package pprint_test

import (
	pp "github.com/wwnbb/pprint"
	"testing"
)

type TestStruct struct {
	Name string
	Age  int
}

// Complex data structure
type ComplexStruct struct {
	Name     string
	Age      int
	Children []TestStruct
}

func TestPprint(t *testing.T) {
	testStruct := TestStruct{
		Name: "John",
		Age:  30,
	}
	pp.PrettyPrint(testStruct)
}

func TestPprintComplex(t *testing.T) {
	complexStruct := ComplexStruct{
		Name: "John",
		Age:  30,
		Children: []TestStruct{
			{Name: "Jane", Age: 25},
			{Name: "Jack", Age: 20},
		},
	}
	pp.PrettyPrint(complexStruct)
}
