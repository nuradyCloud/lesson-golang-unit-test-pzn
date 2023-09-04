package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldNur(t *testing.T) {
	result := HelloWorld("Nur")

	if result != "Hello Nur" {
		//t.Fail()
		t.Error("Result must be Hello Nur")
	}
	fmt.Println("TestHelloWorldNur done")
}

func TestHelloWorldAdy(t *testing.T) {
	result := HelloWorld("Ady")

	if result != "Hello Ady" {
		//t.FailNow()
		t.Fatal("Result must be Hello Ady")
	}

	fmt.Println("TestHelloWorldAdy done")
}
