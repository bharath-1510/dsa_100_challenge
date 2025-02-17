// Approach 2 :  Using Insertion Sort

// Original array : {1, 2, -3, 4, -5}

// Iterate through the numbers 

// if number is negative , perform insertion sort (rotate the the array by 1 anticlockwise.
// else continue

// Result Array : {-3 , -5 ,1, 2, 4}

func rearrangeArray_using_insertion_sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		if arr[i] < 0 {
			temp := arr[i]
			right := i - 1
			for ; right >= 0 && arr[right] > 0; right-- {
				arr[right+1] = arr[right]
			}
			arr[right+1] = temp
		}
	}
}
