package main

import "fmt"

func MeanMedian(value []float64) (mean, median float64) {
	var total float64
	for i := 0; i < len(value); i++ {
		total += value[i]
		mean = total / float64(len(value))
	}

	if len(value)%2 != 0 {
		median = value[len(value)/2]
	} else {
		median = (value[(len(value)-1)/2] + value[(len(value)/2)]) / 2
	}

	return mean, median

}

func main() {

	fmt.Println(MeanMedian([]float64{1, 2, 3, 4}))
	fmt.Println(MeanMedian([]float64{1, 2, 3, 4, 5}))
	fmt.Println(MeanMedian([]float64{7, 8, 9, 13, 15}))
	fmt.Println(MeanMedian([]float64{10, 20, 30, 40, 50}))
	fmt.Println(MeanMedian([]float64{15, 20, 30, 60, 120}))
}
