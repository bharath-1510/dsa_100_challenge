package day06

import (
	"sort"
)

func MostFrequentElementArrayBruteForce(arr []int) int {
	maxCnt := -1
	pos := 0
	for i := 0; i < len(arr); i++ {

		cnt := 0
		for j := 0; j < len(arr); j++ {
			if arr[j] == arr[i] {
				cnt += 1
			}
		}
		if cnt > maxCnt {
			maxCnt = cnt
			pos = i
		}
	}
	return arr[pos]
}
func MostFrequentElementArray(arr []int) int {
	cnt := -1
	freq := make(map[int]int)
	ans := 0
	for _, i := range arr {
		freq[i]++
		if freq[i] > cnt {
			cnt = freq[i]
			ans = i
		}

	}

	return ans
}

func MostFrequentElementArrayUsingSorting(arr []int) int {
	maxCnt := -1
	cnt := 1
	pos := -1
	sort.Ints(arr)
	for index := 0; index < len(arr)-1; index++ {
		if arr[index] == arr[index+1] {
			cnt++
		} else {
			if maxCnt < cnt {
				maxCnt = cnt
				pos = index
			}
			cnt = 1
		}
	}
	if maxCnt < cnt {
		pos = len(arr) - 1
	}
	return arr[pos]
}
