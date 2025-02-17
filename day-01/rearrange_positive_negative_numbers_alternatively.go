// Approach 1 :  Using Two Arrays

// Original array : {1, 2, -3, 4, -5}

// Iterate through the numbers 

// if number is negative , add in negativeArray
// else add in positiveArray

// append both arrays

// Result Array : {-3 , -5 ,1, 2, 4}

func rearrangeArray_using_two_arrays(arr []int) {
	var negativeArray = make([]int, 0, 0)
	var positiveArray = make([]int, 0, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			negativeArray = append(negativeArray, arr[i])
		}

		if arr[i] > 0 {
			positiveArray = append(positiveArray, arr[i])
		}
	}
	j := 0
	for i := len(negativeArray); i < len(arr); i++ {
		negativeArray = append(negativeArray, positiveArray[j])
		j++
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = negativeArray[i]
	}

}


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
