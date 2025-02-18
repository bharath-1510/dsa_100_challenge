// Approach : Using Two Pointer
func stockSellBuy(arr []int) int {
	max := 0
	slow := 0
	fast := 1
	for fast < len(arr) {
		temp := arr[fast] - arr[slow]
		if temp < 0 {
			slow = fast
		}
		fast++
		if temp > max {
			max = temp
		}
	}
	return max
}
