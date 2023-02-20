package main

import (
	"fmt"
)

func place_length(arr []string) (arr1 []int) {
	length_arr := [7]int{}
	for i := 0; i < len(arr); i++ {
		l := len(arr[i])
		length_arr[i] = l
	}
	return length_arr[:]
}

func main() {
	places := [7]string{}
	fmt.Println("Enter the Places:")
	for i := range places {
		var x string
		fmt.Scan(&x)
		places[i] = x
	}
	length := place_length(places[:])
	fmt.Println("Places with their length")
	for i := 0; i < len(length); i++ {
		fmt.Printf("%v : %v\n", places[i], length[i])
	}
}
