package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	type Product struct {
		Name       string
		PriceCents int
	}
	v := reflect.ValueOf(&Product{"Pizza Speciale", 1299})
	uppercaseStructFields(v.Elem())
	fmt.Println(v.Elem())

	printStructFieldDetails(v.Elem().Type())
}

func uppercaseStructFields(v reflect.Value) {
	if v.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < v.NumField(); i++ {
		var field reflect.Value = v.Field(i)
		if field.Kind() == reflect.String {
			field.SetString(strings.ToUpper(field.String()))
		}
	}
}

func printStructFieldDetails(t reflect.Type) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field %d: %s %s\n", i, field.Name, field.Type)
	}
}
