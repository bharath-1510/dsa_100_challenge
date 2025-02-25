package day03

func IsArraySorted(arr []int) bool {
	i := 0
	for i < len(arr)-1 {
		if arr[i] > arr[i+1] {
			return false
		}
		i++
	}
	return true
}
