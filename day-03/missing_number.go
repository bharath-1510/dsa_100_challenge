package day03

// Brute force
func FindMissingNumber(arr []int) int {

	for i := 1; i <= len(arr)+1; i++ {
		flag := false
		for _, val := range arr {
			if val == i {
				flag = true
				break
			}
		}
		if !flag {
			return i
		}
	}
	return 0
}

// Using Formula n*(n-1)/2
func FindMissingNumberUsingFormula(arr []int) int {
	n := len(arr) + 2
	sum := 0
	for i := 0; i < len(arr); i++ {
		if i < len(arr)-1 && arr[i] == arr[i+1] {
			n--
			continue
		}
		sum = sum + arr[i]
	}
	res := n * (n - 1) / 2
	return res - sum
}

//Using Map
func FindMissingNumberUsingHash(arr []int) int {
	flag := make(map[int]bool)
	for _, val := range arr {
		flag[val] = true
	}
	for i := 1; i <= len(arr)+1; i++ {
		if !flag[i] {
			return i
		}
	}
	return 0
}
