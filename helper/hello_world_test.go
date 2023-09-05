package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkSub(b *testing.B) {
	b.Run("Ady", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Ady")
		}
	})
	b.Run("Nur Ady", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Nur Ady")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ady")
	}
}

func BenchmarkHelloWorldName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Nur Ady")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		req      string
		expected string
	}{
		{
			name:     "Nur",
			req:      "Nur",
			expected: "Hello Nur",
		},
		{
			name:     "Ady",
			req:      "Ady",
			expected: "Hello Ady",
		},
		{
			name:     "Pamungkas",
			req:      "Pamungkas",
			expected: "Hello Pamungkas",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.req)
			require.Equal(t, test.expected, result)
		})
	}
}

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
