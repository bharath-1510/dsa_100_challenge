func FindSecondSmallestAndSecondLargest(arr []int) (int, int) {
	if len(arr) == 1 {
		return arr[0], arr[0]
	}
	max := math.MinInt
	pos := 0
	max2 := math.MinInt
	min := math.MaxInt
	min2 := math.MaxInt
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			pos = i
		}
	}
	for i := 0; i < len(arr); i++ {
		if i != pos && arr[i] > max2 {
			max2 = arr[i]
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
			pos = i
		}
	}
	for i := 0; i < len(arr); i++ {
		if i != pos && arr[i] < min2 {
			min2 = arr[i]
		}
	}
	return min2, max2
}
