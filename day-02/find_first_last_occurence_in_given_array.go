// Approach : Using Two Pointer
func FindFirstandLastOccurence(arr []int, target int) (int, int) {
	first := -1
	last := -1

	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			first = i
			for i < len(arr) && arr[i] == target {
				i++
			}
			last = i - 1
		}
	}
	return first, last
}
