// Bruteforce Approach
func RotateByKInLeft(arr []int, k int) []int {
	k = k % len(arr)
	res := make([]int, len(arr))
	copy(res, arr)
	for i := 0; i < k; i++ {
		temp := res[0]
		for j := 0; j < len(res)-1; j++ {
			res[j] = res[j+1]
		}
		res[len(arr)-1] = temp
	}
	return res
}

// Using Temp Array
func RotateByTempArray(arr []int, k int) []int {
	res := make([]int, len(arr))
	k = k % len(arr)
	j := 0
	for i := k; i < len(arr); i++ {
		res[j] = arr[i]
		j++
	}
	for i := 0; i < k; i++ {
		res[j] = arr[i]
		j++
	}
	return res
}

// Using Reverse
func RotateByReversal(arr []int, k int) []int {
	k = k % len(arr)
	reverse(arr, 0, k-1)
	reverse(arr, k, len(arr)-1)
	reverse(arr, 0, len(arr)-1)
	return arr
}

func reverse(arr []int, i int, j int) {
	start := i
	end := j
	for start < end {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
	}
}
