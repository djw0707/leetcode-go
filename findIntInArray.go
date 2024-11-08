package findIntInArray

// 数组从左到右 从上到下都是有序的，可以从右上角开始二分查找 大于目标值向左查询 小于目标值向下查询
func dfs(i int, j int, matrix *[][]int, target int) bool {
	if i >= len(*matrix) || j < 0 {
		return false
	}
	if (*matrix)[i][j] > target {
		return dfs(i, j-1, matrix, target)
	}
	if (*matrix)[i][j] < target {
		return dfs(i+1, j, matrix, target)
	}
	return true
}

func searchMatrix(matrix [][]int, target int) bool {
	return dfs(0, len(matrix[0])-1, &matrix, target)
}
