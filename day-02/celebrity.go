package day02

// Bruteforce Approach
func FindCelebrityUsingBruteForce(matrix [][]int) int {
	n := len(matrix)
	for i := 0; i < n; i++ {
		flag := true
		for j := 0; j < n; j++ {
			if i != j && (matrix[i][j] == 1 || matrix[j][i] == 0) {
				flag = false
				break
			}
		}
		if flag {
			return i
		}
	}
	return -1
}

