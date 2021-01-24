package search_a_2d_matrix

func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) <= 0 || len(matrix[0]) <= 0 {
		return false
	}
	row, col := len(matrix), len(matrix[0])
	// 将矩阵转化为一维数组求解
	left, right := 0, row * col - 1
	for left <= right {
		mid := (left + right) / 2
		// 将mid位置转化为矩阵的行，列
		if matrix[mid/col][mid%col] == target {
			return true
		}else if matrix[mid/col][mid%col] > target {
			right = mid - 1
		} else{
			left = mid + 1
		}
	}
	return false
}
