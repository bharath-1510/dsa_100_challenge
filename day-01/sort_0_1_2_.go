package day01

// Approach : Using Dutch National Flag

// Original Array : {0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1}

// Initialize start , mid as 0 and end as lenght of the array -1

// Iterate till mid <= end
// if number is 0 , then swap start and mid numbers and increment the start and mid by 1
// else if number is 1 , then increment the mid by 1
// else then swap end and mid numbers and decrement the end by 1

// Result Array : {0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2}

func sort_0_1_2_using_dnf(arr []int) {
	start := 0
	mid := 0
	end := len(arr) - 1
	for mid <= end {
		if arr[mid] == 0 {
			temp := arr[start]
			arr[start] = arr[mid]
			arr[mid] = temp
			start++
			mid++
		} else if arr[mid] == 1 {
			mid++
		} else {
			temp := arr[end]
			arr[end] = arr[mid]
			arr[mid] = temp
			end--
		}
	}
}
