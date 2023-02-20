package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	int_slice := make([]int, 3)
	i := 0
	for {
		var x string
		fmt.Println("Enter a number(ENter 'X' to exit")
		fmt.Scan(&x)
		if x == "X" {
			break
		}
		num, err := strconv.Atoi(x)
		if err != nil {
			fmt.Println("Invalid input\nEnter an Integer or \"X\" to exit")
		}
		if i <= 2 {
			int_slice[i] = num
		} else {
			int_slice = append(int_slice, num)
		}
		i++
	}
	sort.Ints(int_slice)
	fmt.Println(int_slice)
}
