package design_a_leaderboard

import "sort"

type Leaderboard struct {
	// 记录参赛者分数
	count []int
}


func Constructor() Leaderboard {
	return Leaderboard {
		count: make([]int, 10000),
	}
}

// 添加分数
func (this *Leaderboard) AddScore(playerId int, score int)  {
	this.count[playerId] += score
}

// 求最大K个数的和
func (this *Leaderboard) Top(K int) int {
	tmp := make([]int, len(this.count))
	copy(tmp, this.count)
	// 按分数从大到小排
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] > tmp[j]
	})
	sum := 0
	for i := 0; i < K; i++ {
		sum += tmp[i]
	}
	return sum
}

// 重置指定参赛者分数
func (this *Leaderboard) Reset(playerId int)  {
	this.count[playerId] = 0
}


/**
 * Your Leaderboard object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddScore(playerId,score);
 * param_2 := obj.Top(K);
 * obj.Reset(playerId);
 */
