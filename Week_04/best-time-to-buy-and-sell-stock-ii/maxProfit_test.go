package best_time_to_buy_and_sell_stock_ii

import "testing"

func TestMaxProfit(t *testing.T) {
	prices := []int{7,1,5,3,6,4}
	want := 7
	got := MaxProfit(prices)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	prices = []int{1,2,3,4,5}
	want = 4
	got = MaxProfit(prices)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	prices = []int{7,6,4,3,1}
	want = 0
	got = MaxProfit(prices)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
