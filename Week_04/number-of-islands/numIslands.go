package number_of_islands

func NumIslands(grid [][]byte) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 是陆地
			if grid[i][j] == '1' {
				ans++
				dfs(i, j, grid)
			}
		}
	}
	return ans
}

func dfs(row, col int, grid [][]byte) {
	// 越界或者不是陆地
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] != '1' {
		return
	}
	grid[row][col] = '0'
	// 四个方向继续处理
	dfs(row+1, col, grid)
	dfs(row-1, col, grid)
	dfs(row, col+1, grid)
	dfs(row, col-1, grid)
}
