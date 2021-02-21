package n_queens

import "strings"

func SolveNQueens(n int) [][]string {
	ans := [][]string{}
	board := make([]string, n)
	// 初始时棋牌全是'.'
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}
	backtrack(&ans, board, 0)
	return ans
}

func backtrack(ans *[][]string, board []string, row int) {
	// 已经放满，记录一个结果
	if row == len(board) {
		*ans = append(*ans, append([]string{}, board...))
		return
	}
	// 在当前行每一列放置一个皇后试探
	for col := 0; col < len(board); col ++ {
		if !isValid(board, row, col) {
			continue
		}
		// 放置皇后
		stringAt(&board[row], col, 'Q')
		// 继续放置下一行
		backtrack(ans,board, row + 1)
		// 恢复状态
		stringAt(&board[row], col, '.')
	}
}

func stringAt(s *string, col int, v rune) {
	rs := []rune(*s)
	rs[col] = v
	*s = string(rs)
}

func isValid(board []string, row, col int) bool {
	// 检查列是否有冲突
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// 左上对角线
	for i, j := row - 1, col - 1; i >= 0 && j >= 0 ; i, j = i - 1, j - 1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	// 右上对角线
	for i, j := row - 1, col + 1; i >= 0 && j < len(board); i, j = i - 1, j + 1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}
