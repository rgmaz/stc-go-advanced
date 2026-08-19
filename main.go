package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Ada", Age: 36}

	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		val := v.Field(i)
		fmt.Printf("%s: %v\n", field.Name, val.Interface())
	}
}
