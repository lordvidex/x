package ptr

import (
	"fmt"
	"strconv"
	"testing"
)

func ExampleToObj() {
	type test struct {
		a, b int
		c, d string
	}
	y := &test{a: 4, b: 0, c: "hello", d: ""}
	var z *test
	fmt.Println(ToObj(y), ToObj(z))
	// Output: {4 0 hello } {0 0  }
}

func ExampleToString() {
	str := "hello"
	fmt.Println(ToString(&str))
	// Output: hello
}

func ExampleToString_empty() {
	fmt.Println(ToString(nil))
	// Output:
}

var globStr string

func BenchmarkToString(b *testing.B) {
	var res string
	b.Run("ToObj", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			str := String("Hello" + strconv.Itoa(i))
			res = ToObj(str)
		}
		globStr = res
	})

	b.Run("ToString", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			str := String("Hella" + strconv.Itoa(i))
			res = ToString(str)
		}
		globStr = res
	})
}
