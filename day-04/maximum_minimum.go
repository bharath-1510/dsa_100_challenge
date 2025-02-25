package day04

import "math"

func FindMaximumAndMinimum(arr []int) (int, int) {
	max := math.MinInt
	min := math.MaxInt
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min, max
}
