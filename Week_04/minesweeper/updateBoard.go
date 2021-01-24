package minesweeper

func UpdateBoard(board [][]byte, click []int) [][]byte {
	m, n := len(board), len(board[0])
	row, col := click[0], click[1]
	// 挖到雷
	if board[row][col] == 'M' {
		board[row][col] = 'X'
	} else {
		count := 0
		for i := -1; i < 2; i++ {
			for  j := -1; j < 2; j++ {
				if i == 0 && j == 0 {
					continue
				}
				r, c := row + i, col + j
				if r < 0 || r >= m || c < 0 || c >= n {
					continue
				}
				if board[r][c] == 'M' || board[r][c] == 'X' {
					count++
				}
			}
		}
		if count > 0 {
			board[row][col] = byte(count + int('0'))
		} else {
			board[row][col] = 'B'
			for i := -1; i < 2; i++ {
				for j := -1; j < 2; j++ {
					if i == 0 && j == 0 {
						continue
					}
					r, c := row + i, col + j
					if r < 0 || r >= m || c < 0 || c >= n {
						continue
					}
					if board[r][c] == 'E' {
						UpdateBoard(board, []int{r, c})
					}
				}
			}
		}
	}
	return board
}
