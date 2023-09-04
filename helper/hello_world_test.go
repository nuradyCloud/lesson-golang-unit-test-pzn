package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Nur")

	assert.Equal(t, "Hello Nur", result, "Result must be 'Hello Nur'")
	fmt.Println("Test Hello world with assert is done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Nur")

	require.Equal(t, "Hello Nur", result, "Result must be 'Hello Nur'")
	fmt.Println("Test Hello world with assert is done")
}

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
