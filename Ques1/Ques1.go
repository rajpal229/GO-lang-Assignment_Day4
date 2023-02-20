package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var s string
	fmt.Println("Enter comma seperated numbers: ")
	fmt.Scan(&s)
	str_arr := strings.Split(s, ",")
	fmt.Printf("This %v is Type - %v\n", str_arr, reflect.TypeOf(str_arr).Kind())
	// str_slice := str_arr[:]
	// fmt.Printf("This %v is Type - %v\n", str_slice, reflect.TypeOf(str_slice).Kind())
}
