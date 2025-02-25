package day02

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

// Approach : Using Binary Search
func FindFirstandLastOccurenceUsingBinarySearch(arr []int, target int) (int, int) {
	
	first := -1
	low := 0
	high := len(arr) - 1
	// Searching for First Occurence
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			first = mid
			high = mid - 1
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}
	last := -1
	low = 0
	high = len(arr) - 1
	// Searching for Last Occurence
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			last = mid
			low = mid + 1
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}
	return first, last

}
