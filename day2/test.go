package main

import (
    "fmt"
    "reflect"
)

type Test struct {
    Test string
}

func (test *Test) setProperty(propName string, propValue string) *Test {
	reflect.ValueOf(test).Elem().FieldByName(propName).Set(reflect.ValueOf(propValue))
	return test
}

func _main() {
    var test Test
    test.setProperty("Test", "Updated value") 
    fmt.Println(test.Test) 
}
