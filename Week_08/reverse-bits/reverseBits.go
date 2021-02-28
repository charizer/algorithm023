package reverse_bits

func ReverseBits(num uint32) uint32 {
	ans := uint32(0)
	// 依次获取最低位然后左移
	for i:=0; i < 32; i++ {
		ans = ans << 1 + num & 1
		num = num >> 1
	}
	return ans
}
