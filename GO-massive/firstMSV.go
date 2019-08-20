package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x[4])

	fmt.Println(x)

	var y [5]float64
	y[0] = 11
	y[1] = 22
	y[2] = 33
	y[3] = 44
	y[4] = 56

	fmt.Println(y)

	var total float64 = 0.0
	for i := 0; i < 5; i++ {
		total += y[i]
	}
	fmt.Println(total)
	fmt.Println(total / 5)
	fmt.Println(total / 5.0) //search average

	var max float64 = 0 //search max element
	max = y[0]
	for i := 0; i < 5; i++ {
		if y[i] > max {
			max = y[i]
		}
	}
	fmt.Println(max)

	var total_2 float64 = 0
	for i := 0; i < len(y); i++ {
		total_2 += y[i]
	}
	fmt.Println(len(y))
	fmt.Println(total_2 / float64(len(y))) //convert int to float64 because total_2 is float64

}
