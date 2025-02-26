package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	N int
}

func main() {
	myStruct := MyStruct{}
	vStruct := reflect.ValueOf(&myStruct) // pass-by-reference!
	process(vStruct.Elem().Field(0))      // de-ref and access field
	fmt.Printf("%+vStruct\n", myStruct)
}

func process(v reflect.Value) {
	fmt.Println(v.CanSet())
	v.SetInt(456)
}
