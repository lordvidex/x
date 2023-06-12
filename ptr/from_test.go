package ptr

import (
	"fmt"
	"strconv"
	"testing"
	"github.com/stretchr/testify/assert"
)

func ExampleBool() {
	x := true
	y := Bool(x)
	fmt.Println(*y)
	// Output: true
}

func ExampleFloat64() {
	x := 1.0
	y := Float64(x)
	fmt.Println(*y)
	// Output: 1
}

func ExampleString() {
	x := "hello"
	y := String(x)
	fmt.Println(*y)
	// Output: hello
}

func ExampleObj() {
	x := 1
	y := Obj(x)
	fmt.Println(*y)
	// Output: 1
}

func ExampleObj_struct() {
	type t struct {
		x int
		_ *string
	}
	x2 := t{
		x: 1,
	}
	y2 := Obj(x2)
	fmt.Println(*y2)
	// Output: {1 <nil>}
}

func TestObj(t *testing.T) {
	testcases := []struct {
		name string
		obj any
		expect any
	} {

	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := Obj(tc.obj)
			assert.Equal(t, tc.expect, got)
		})
	}
}

var glob *string

func BenchmarkObjLiteral(b *testing.B) {
	// test the speed difference between Obj string and String
	var x *string
	b.Run("Obj", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			str := "hello" + strconv.Itoa(i)
			x = Obj(str)
		}
		glob = x
	})
	b.Run("String", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			str := "hella" + strconv.Itoa(i)
			x = String(str)
		}
		glob = x
	})
}