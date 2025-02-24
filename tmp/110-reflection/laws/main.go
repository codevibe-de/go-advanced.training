package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	type Centimeter int
	var length Centimeter = 192
	//var n int = length // Cannot use 'length' (type Centimeter) as the type int
	_ = length

	var a any = length
	_ = a

	var r io.Reader        // static type is io.Reader
	r = os.Stdin           // actual type is *os.File
	r = bufio.NewReader(r) // actual type is *bufio.Reader

	var t reflect.Type = reflect.TypeOf(length)
	var v reflect.Value = reflect.ValueOf(length)
	fmt.Println(t, v) // Output: main.Centimeter 192

	var k reflect.Kind = t.Kind()
	fmt.Printf("Type is %v, which is based on %v\n", t, k)

	var length2 any = v.Interface()
	fmt.Println(length2)

	var length3 Centimeter = length2.(Centimeter)
	fmt.Printf("Recovered value %v of type %T\n", length3, length3)

	var n int = 23
	v2 := reflect.ValueOf(n)
	//v2.SetInt(42) // panic

	fmt.Println(v2.CanSet()) // Output: false

	v3 := reflect.ValueOf(&n)
	fmt.Printf("Type: %s, CanSet: %v\n", v3.Type(), v3.CanSet())
	// Output: Type: *int, CanSet: false
	//v3.SetInt(42) // also panics!

	v4 := v3.Elem()
	fmt.Printf("Type: %s, CanSet: %v\n", v4.Type(), v4.CanSet())
	// Output: Type: int, CanSet: true
	v4.SetInt(42) // works, finally
}
