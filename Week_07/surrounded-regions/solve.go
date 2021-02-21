package surrounded_regions

func Solve(board [][]byte)  {
	m := len(board)
	if m <= 0 {
		return
	}
	n := len(board[0])
	// 处理边缘列上的'O'
	for i := 0; i < m; i++ {
		dfs(board, i, 0)
		dfs(board, i, n - 1)
	}
	// 处理边缘行上的'O'
	for j := 0; j < n; j++ {
		dfs(board, 0, j)
		dfs(board, m - 1, j)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 将改过的'A'变为'O'
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			}else if board[i][j] == 'O' {
				// 没变过的'O'，被'X'包围
				board[i][j] = 'X'
			}
		}
	}
}

// 将于边缘上的'O'相连的字符变为'A'
func dfs(board [][]byte, row, col int) {
	if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) || board[row][col] != 'O' {
		return
	}
	board[row][col] = 'A'
	dfs(board, row + 1, col)
	dfs(board, row - 1, col)
	dfs(board, row, col + 1)
	dfs(board, row, col - 1)
}
