package main

import "fmt"

/*
func sumInts(values []int) int {
	var result int
	for _, val := range values {
		result += val
	}
	return result
}

func sumFloats(values []float32) float32 {
	var result float32
	for _, val := range values {
		result += val
	}
	return result
}
*/

/*
func sum[T int | float32](values []T) T {
	var result T
	for _, val := range values {
		result += val
	}
	return result
}
*/

type Summable interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func sum[T Summable](values []T) T {
	var result T
	for _, val := range values {
		result += val
	}
	return result
}

func main() {
	ints := []int{4, 1, 5, 2, 3}
	// sumInts := sumInts(ints)
	// sumInts := sum[int](ints)
	sumInts := sum(ints)
	fmt.Println(sumInts)

	floats := []float32{3.5, 1.7, 4.9, 2.8, 5}
	// sumFloats := sumFloats(floats)
	// sumFloats := sum[float32](floats)
	sumFloats := sum(floats)
	fmt.Println(sumFloats)
}
