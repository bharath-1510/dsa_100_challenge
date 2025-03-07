package day07

func MaximumSumSubArrayWithIndex(arr []int) []int {
	res := make([]int,0)
	maxSum := arr[0]
	currSum := arr[0]
	start ,end := 0,0
	s :=0
	for i, _ := range arr {
		if i == 0 {
			continue
		}
		if arr[i] >  currSum+arr[i] {
			s=i
		}
		currSum = max(arr[i], currSum+arr[i])
		if currSum > maxSum {
			start = s
			end = i
		}
		maxSum = max(maxSum, currSum)
	}
	res = append(res, maxSum)
	res = append(res, start)
	res = append(res, end)
	return res
}

func MaximumSumSubArray(arr []int) int {
	maxSum := arr[0]
	currSum := arr[0]
	for i, _ := range arr {
		if i == 0 {
			continue
		}
		currSum = max(arr[i], currSum+arr[i])
		maxSum = max(maxSum, currSum)
	}
	return maxSum
}