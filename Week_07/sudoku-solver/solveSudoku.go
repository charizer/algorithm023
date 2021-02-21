package sudoku_solver

func SolveSudoku(board [][]byte)  {
	if len(board) != 9 || len(board[0]) != 9 {
		return
	}
	solve(board)
}

func solve(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				continue
			}
			// 用'1' ~ '9'试探
			for ch := '1'; ch <= '9'; ch++ {
				bc := byte(ch)
				// 判断字符有效性
				if isValid(board, i, j, bc) {
					// 替换
					board[i][j] = bc
					// 已经满足条件，提前返回
					if solve(board) {
						return true
					} else {
						// 恢复状态
						board[i][j] = '.'
					}
				}
			}
			return false
		}
	}
	return true
}

// 判断字符填入是否是有效
func isValid(board [][]byte, row, col int, ch byte) bool {
	for i := 0; i < 9; i++ {
		// 行或列已经存在次字符
		if board[i][col] == ch || board[row][i] == ch {
			return false
		}
		// 3 * 3 单元格内有次字符
		if board[3 * (row / 3) + i / 3][ 3 * (col / 3) + i % 3] == ch {
			return false
		}
	}
	return true
}
