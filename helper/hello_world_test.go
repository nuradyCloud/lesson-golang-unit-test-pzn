package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ady")

	if result != "Hello Ady" {
		panic("Result is not 'Hello Ady'")
	}
}

func TestHelloWorldNur(t *testing.T) {
	result := HelloWorld("Nur")

	if result != "Hello Nur" {
		panic("Result is not 'Hello Nur'")
	}
}
