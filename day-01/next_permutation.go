package day01
import (
	"sort"
)
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

// Approach 2 : Using Pivot value

// Original Array : {2, 4, 1, 7, 5, 0}

// Find the number where present value is greater than next value and store as pivot
// if not found , initialize as -1

// if pivot is -1 , then reverse the array ( This will give us the first permutation. Example : arr ={3,2,1} and Output will be {1,2,3} )

// find the number greater the pivot and swap the present number <-> pivot
// reverse the array from pivot and end of the array

// Result Array : {2, 4, 5, 0, 1, 7}

func nextPermutation(arr []int) {
	pivot := -1
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] < arr[i+1] {
			pivot = i
			break
		}
	}
	if pivot == -1 {
		reverse(arr, 0, len(arr)-1)
		return
	}
	for i := len(arr) - 1; i > pivot; i-- {
		if arr[i] > arr[pivot] {
			swap(arr, i, pivot)
			break
		}
	}
	reverse(arr, pivot+1, len(arr)-1)
}

func reverse(arr []int, start int, end int) {
	for start < end {
		swap(arr, start, end)
		start++
		end--
	}
}
