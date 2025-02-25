package day05
// Bruteforce Approach
func RotateByKInRight(arr []int, k int) []int {
	k = k % len(arr)
	res := make([]int, len(arr))
	copy(res, arr)
	for i := 0; i < k; i++ {
		temp := res[len(arr)-1]
		for j := len(res) - 1; j > 0; j-- {
			res[j] = res[j-1]
		}
		res[0] = temp
	}
	return res
}

// Using Temp Array
func RotateByTempRightArray(arr []int, k int) []int {
	k = k % len(arr)
	res := make([]int, len(arr))
	j := 0

	for i := len(arr) - k; i < len(arr); i++ {
		res[j] = arr[i]
		j++
	}

	for i := 0; i < len(arr)-k; i++ {
		res[j] = arr[i]
		j++
	}
	return res
}
