package day04

// Approach 1 : Using Auxillary Array
func reverseArray_using_aux_array(arr []int) {
	res := make([]int, len(arr))
	start := 0
	for i := len(arr) - 1; i >= 0; i-- {
		res[start] = arr[i]
		start++
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = res[i]
	}
}

// Approach 2 :  Using Two Pointer
func reverseArrayTwoPointer(arr []int) {
	start := 0
	end := len(arr) - 1
	for start <= end {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
	}
}
