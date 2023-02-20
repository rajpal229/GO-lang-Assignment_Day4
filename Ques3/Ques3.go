package main

import (
	"fmt"
	"math"
	"strconv"
)

func average(arr []float64) float64 {
	var sum float64
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	avg := sum / float64(len(arr))
	return avg
}

func std_dev(arr []float64) float64 {
	avg := average(arr)
	var sum float64
	for i := 0; i < len(arr); i++ {
		sum = sum + math.Pow(math.Abs(arr[i]-avg), 2)
	}
	sd := math.Sqrt(sum / float64(len(arr)))
	return sd
}

func main() {
	var float_slice []float64
	for {
		var inp string
		fmt.Println("Enter a floating point number(Enter q or Q to stop):")
		fmt.Scan(&inp)
		if inp == "q" || inp == "Q" {
			break
		}
		num, err := strconv.ParseFloat(inp, 64)
		if err != nil {
			fmt.Println("Invalid Input\nPlease enter a floating point number or q/Q to quit")
			continue
		}
		float_slice = append(float_slice, num)
	}
	fmt.Println("The slice of floating point number:", float_slice)
	fmt.Printf("\nAverage of slice is = %v\nStandard Deviation of slice is %.5f", average(float_slice), std_dev(float_slice))
}
