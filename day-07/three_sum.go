package day07

import "sort"

func TripletSumUsingHashing(arr []int, target int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(arr)-1; i++ {
		temp := make(map[int]int)
		tempSum := target - arr[i]
		for j := i + 1; j < len(arr); j++ {
			if _, ok := temp[tempSum-arr[j]]; ok {
				res = append(res, []int{arr[i], arr[j], tempSum - arr[j]})
			}
			temp[arr[j]] = j
		}
	}
	return res
}

func TripletSumBruteForce(arr []int, target int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				sum := arr[i] + arr[j] + arr[k]
				if sum == target {
					res = append(res, []int{arr[i], arr[j], arr[k]})
				}
			}
		}
	}
	return res
}

func TripletSumUsingSorting(arr []int, target int) [][]int {
	res := make([][]int, 0)
	temp := make([]int, len(arr))
	copy(temp, arr)
	sort.Ints(temp)
	for i := 0; i < len(arr)-2; i++ {
		l, r := i+1, len(arr)-1
		for l < r {
			tempSum := temp[i] + temp[l] + temp[r]
			if tempSum == target {
				res = append(res, []int{temp[i], temp[l], temp[r]})
				l++
				r--
			} else if tempSum < target {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
