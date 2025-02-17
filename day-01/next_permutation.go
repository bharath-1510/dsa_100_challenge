// Approach 1 : By Generating all permutations

// Original Array : {2, 4, 1, 7, 5, 0}

// Find all the permutation of array by recursion and swapping the numbers
// Store it in a 2d array

// Sort the 2D array

// find the input array 
// if it is last in the array , then update arr as first permutation
// else update the next permutation after the permutation found

// Result Array : {2, 4, 5, 0, 1, 7}

// Finding the Next Permutation of an array by generating all permutations
func nextPermutationAll(arr []int) {
	allPermutation := [][]int{}
	findPermutation(&allPermutation, arr, 0)

  // Custom sort function
	compare := func(a, b []int) bool {
		for i := range a {
			if a[i] != b[i] {
				return a[i] < b[i]
			}
		}
		return false
	}

	sort.Slice(allPermutation, func(i, j int) bool {
		return compare(allPermutation[i], allPermutation[j])
	})
	for i := range allPermutation {
		flag := true

    // searching for input array
		for j := range allPermutation[i] {
			if allPermutation[i][j] != arr[j] {
				flag = false
				break
			}
		}
    
		if !flag {
			continue
		}
    
    // update the next permutation after the permutation found
		if i < len(allPermutation)-1 {
			res := make([]int, len(allPermutation[i+1]))
			copy(res, allPermutation[i+1])
			for j := 0; j < len(res); j++ {
				arr[j] = res[j]
			}
		}

    // if last in the array , then update arr as first permutation
		if i == len(allPermutation)-1 {
			res := make([]int, len(allPermutation[0]))
			copy(res, allPermutation[0])
			for j := 0; j < len(res); j++ {
				arr[j] = res[j]
			}
		}
		break
	}
}

// Generating all Permutation for an array
func findPermutation(allPermutation *[][]int, arr []int, pos int) {
	if pos == len(arr)-1 {
		temp := make([]int, len(arr))
		copy(temp, arr)
		*allPermutation = append(*allPermutation, temp)
		return
	}
	for i := pos; i < len(arr); i++ {
		swap(arr, pos, i)
		findPermutation(allPermutation, arr, pos+1)
		swap(arr, pos, i)
	}
}

// swap two numbers
func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
