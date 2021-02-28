package n_queens_ii

func TotalNQueens(n int) int {
	ans := 0
	size := (1 << n) - 1
	dfs(0, 0, 0, size, &ans)
	return ans
}

func dfs(row, ld, rd, size int, ans *int){
	if row == size {
		*ans++
		return
	}
	//取得现在可以填充的位置，1表示可以填充的，0表示已经填充过了
	pos := size &^ (row | ld | rd)
	for pos != 0 {
		//取得pos的最后一位1，也就是这轮需要填充的位置
		p := pos & -pos
		//将pos的最后一位1置为0，表明我们将这个位置填充了
		pos &= pos - 1
		//下探到下一层
		dfs(row | p, (ld | p) << 1, (rd | p) >> 1, size, ans)
	}
}
