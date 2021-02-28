package n_queens

import "math/bits"

func SolveNQueens(n int) [][]string {
	ans := [][]string{}
	size := (1 << n) - 1
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	dfs(queens,0, 0, 0, 0, n, size, &ans)
	return ans
}

func dfs(queens []int, row, cols, ld, rd, n, size int, ans *[][]string){
	if row == n {
		board := generateBoard(queens, n)
		*ans = append(*ans, board)
		return
	}
	//取得现在可以填充的位置，1表示可以填充的，0表示已经填充过了
	pos := size &(^(cols | ld | rd))
	for pos != 0 {
		//取得pos的最后一位1，也就是这轮需要填充的位置
		p := pos & -pos
		//将pos的最后一位1置为0，表明我们将这个位置填充了
		pos &= pos - 1
		//二进制中一的位数
		column := bits.OnesCount(uint(p - 1))
		queens[row] = column
		//下探到下一层
		dfs(queens, row+1, cols | p, (ld | p) << 1, (rd | p) >> 1, n, size, ans)
		queens[row] = -1
 	}
}

func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}
