package day07

import "sort"

func TwoSumUsingSortingWithTwoPointer(arr []int, target int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	sort.Ints(res)
	l , r := 0, len(arr)-1
	for l<r {
		sum := res[l] + res[r]
		if sum == target {
			return []int{res[l],res[r]}
		} else if sum < target {
			l++
		} else {
			r--
		}

	}
	return nil
}

func TwoSumUsingHash(arr []int, target int) []int {
	res := make(map[int]int)
	for i, num := range arr {
		res[num] = i
	}
	for k, _ := range res {
		search := target - k
		if find(arr, search) {
			return []int{search, k}
		}
	}
	return nil
}
func find(arr []int, search int) bool {
	for i := range arr {
		if i == search {
			return true
		}
	}
	return false
}
func TwoSumBruteForce(arr []int, target int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			sum := arr[i] + arr[j]
			if sum == target {
				return []int{arr[i], arr[j]}
			}
		}
	}
	return nil
}
