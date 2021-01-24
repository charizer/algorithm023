package lemonade_change

func LemonadeChange(bills []int) bool {
	// leetcode上bills是空数组时返回true, 不加判断也一致
	// if len(bills) <= 0 {
	//	return true
	// }
	// 贪心算法
	bill5, bill10 := 0, 0
	for _, bill := range bills {
		// 5元，则不用找，5元余额+1
		if bill == 5 {
			bill5++
		// 10元，需要找5元，直接减掉5元，然后判断5元是否小于0，小于0则不满足条件
		}else if bill == 10 {
			bill10++
			bill5--
		}else{
			// 否则是20元， 则可以先找10元再找5元，没有10元则需要找3个5元，减了后如果5元小于则不满足条件
			if bill10 > 0 {
				bill10--
				bill5--
			}else{
				bill5 -=3
			}
		}
		// 前边直接扣的，如果扣了后5元小于0，说明不满足条件
		if bill5 < 0 {
			return false
		}
	}
	return true
}
