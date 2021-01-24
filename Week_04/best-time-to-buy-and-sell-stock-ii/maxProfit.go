package best_time_to_buy_and_sell_stock_ii

func MaxProfit(prices []int) int {
	// 贪心算法
	ans := 0
	for i := range prices {
		// 如果当天比前一天大，则前一天买入，当天抛出
		if i > 0 && prices[i] - prices[i-1] > 0 {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}
