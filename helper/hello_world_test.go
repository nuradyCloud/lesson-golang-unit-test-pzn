package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestSubTest(t *testing.T) {
	t.Run("Nur", func(t *testing.T) {
		result := HelloWorld("Nur")
		require.Equal(t, "Hello Nur", result, "Result must be 'Hello Nur'")
	})

	t.Run("Ady", func(t *testing.T) {
		result := HelloWorld("Ady")
		require.Equal(t, "Hi Ady", result, "Result must be 'Hello Ady'")
	})
}

func TestMain(m *testing.M) {
	//before unit test
	fmt.Println("BEFORE UNIT TEST")

	m.Run()
	//after unit test
	fmt.Println("AFTER UNIT TEST")

}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cant run on Mac OS")
	}
	result := HelloWorld("Nur")
	require.Equal(t, "Hello Nur", result, "Result must be 'Hello Nur'")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Nur")

	assert.Equal(t, "Hello Nur", result, "Result must be 'Hello Nur'")
	fmt.Println("Test Hello world with assert is done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Nur")

	require.Equal(t, "Hello Nur", result, "Result must be 'Hello Nur'")
	fmt.Println("Test Hello world with require is done")
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
