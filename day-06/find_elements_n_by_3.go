package day06

import "sort"

func FindElementsNBy3Times(arr []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(arr); i++ {

		cnt := 0
		for j := i; j < len(arr); j++ {
			if arr[j] == arr[i] {
				cnt += 1
			}
		}

		if cnt > (len(arr) / 3) {

			if len(res) == 0 || arr[i] != res[0] {
				res = append(res, arr[i])
			}
		}

		if len(res) == 2 {
			sort.Ints(res)
			break
		}
	}
	return res
}

func FindElementsNBy3TimesHash(arr []int) []int {
	res := make([]int, 0)
	freq := make(map[int]int)
	for _, i := range arr {
		freq[i]++
	}
	cal := len(arr) / 3
	for k, v := range freq {
		if v > cal {
			res = append(res, k)
		}
	}
	sort.Ints(res)
	return res
}
