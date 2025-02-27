package day06

import "sort"

func FindElementsNByKTimesHash(arr []int, k int) []int {
	res := make([]int, 0)
	freq := make(map[int]int)
	for _, i := range arr {
		freq[i]++
	}
	cal := len(arr) / k
	for k, v := range freq {
		if v > cal {
			res = append(res, k)
		}
	}
	sort.Ints(res)
	return res
}