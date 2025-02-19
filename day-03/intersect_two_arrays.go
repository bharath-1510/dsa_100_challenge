// Bruteforce Approach
func IntersectTwoArrayWithoutDuplicatesMerge(arr1 []int, arr2 []int) []int {
	var res []int
	i := 0
	j := 0
	for i < len(arr1) && j < len(arr2) {
		if i < len(arr1)-1 && arr1[i] == arr1[i+1] {
			i++
			continue
		}
		if arr1[i] == arr2[j] {
			res = append(res, arr1[i])
			i++
			j++
		} else if arr1[i] < arr2[j] {
			i++
		} else {
			j++
		}
	}
	return res
}

// Brute Force (Duplicate Values)
func IntersectTwoArrayWithDuplicatesMerge(arr1 []int, arr2 []int) []int {
	var res []int
	i := 0
	j := 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] == arr2[j] {
			res = append(res, arr1[i])
			i++
			j++
		} else if arr1[i] < arr2[j] {
			i++
		} else {
			j++
		}
	}
	return res
}




// Using Map
func IntersectTwoArray(arr1 []int, arr2 []int) []int {
	flag := make(map[int]bool)
	var res []int
	for _, val := range arr1 {
		flag[val] = true
	}

	for _, val := range arr2 {
		if flag[val] == true {
			res = append(res, val)
			flag[val] = false
		}
	}
	return res
}

// Using Map (Duplicate Values) 
func IntersectTwoArrayWithDuplicates(arr1 []int, arr2 []int) []int {
	flag := make(map[int]int)
	var res []int
	for _, val := range arr1 {
		flag[val]++
	}

	for _, val := range arr2 {
		if flag[val] > 0 {
			res = append(res, val)
			flag[val]--
		}
	}
	return res
}
