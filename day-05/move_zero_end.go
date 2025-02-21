// Bruteforce Approach
func MoveZeroEnd(arr []int) []int {
	pos := 0
	for i := range arr {
		if arr[i] != 0 {
			arr[pos] = arr[i]
			pos++
		}
	}
	for pos < len(arr) {
		arr[pos] = 0
		pos++
	}
	return arr
}

// Using Swap
func MoveZeroEndUsingSwap(arr []int) []int {
	pos := 0
	for i := range arr {
		if arr[i] != 0 {
			temp := arr[pos]
			arr[pos] = arr[i]
			arr[i] = temp
			pos++
		}
	}
	return arr
}
