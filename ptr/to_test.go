package ptr

import "fmt"

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
