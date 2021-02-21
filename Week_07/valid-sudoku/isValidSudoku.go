package valid_sudoku

func IsValidSudoku(board [][]byte) bool {
	if len(board) != 9 || len(board[0]) != 9 {
		 return false
	}
	// 行、列、3*3格子
	rows, cols, boxs := make([]map[int]int, 9), make([]map[int]int, 9), make([]map[int]int, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make(map[int]int)
		cols[i] = make(map[int]int)
		boxs[i] = make(map[int]int)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				continue
			}
			// 格子里的数值
			num := int(board[i][j] - '0')
			// 3 * 3 单元格的位置
			bx := (i / 3 ) * 3 + j / 3
			// 更新行、列、3*3格子对应数值的个数
			rows[i][num]++
			cols[j][num]++
			boxs[bx][num]++
			if rows[i][num] > 1 || cols[j][num] > 1 || boxs[bx][num] > 1 {
				return false
			}
		}
	}
	return true
}
