package string_to_integer_atoi

import "math"

func MyAtoi(s string) int {
	// 删除头部空格
	s = removePrefixSpace(s)
	if len(s) <= 0 {
		return 0
	}
	sign, num := 1, int64(0)
	// 判断第一位是否符号位，然后进行转化
	if s[0] == '+' {
		sign = 1
		num = str2int(s[1:], sign)
	}else if s[0] == '-'{
		sign = -1
		num = str2int(s[1:], sign)
	}else{
		num = str2int(s, sign)
	}
	return int(num)
}

// 删除开始的空格
func removePrefixSpace(s string) string {
	i := 0
	for i = 0; i < len(s); {
		if s[i] == ' ' {
			i++
		}else{
			break
		}
	}
	return s[i:]
}

// string转int，使用int64防止int32溢出
func str2int(s string, sign int) int64 {
	num := int64(0)
	for i := 0; i < len(s); i++ {
		if !isNumberChar(s[i]) {
			return num
		}
		num = num * 10 + int64(sign) * int64(s[i]-'0')
		// 正数溢出
		if sign == 1 && num > math.MaxInt32 {
			return math.MaxInt32
		}
		// 负数溢出
		if sign == -1 && num < math.MinInt32 {
			return math.MinInt32
		}
	}
	return num
}

// 字符是否是数字
func isNumberChar(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

// string转int，使用int32，溢出判断比直接使用int64麻烦
func str2int2(s string, sign int) int64 {
	num := int32(0)
	for i := 0; i < len(s); i++ {
		if !isNumberChar(s[i]) {
			return int64(num)
		}
		// 正数溢出
		if sign == 1 && num > math.MaxInt32/10 || (num == math.MaxInt32/10 && (int32(s[i]-'0') > math.MaxInt32%10))  {
			return math.MaxInt32
		}
		// 负数溢出
		if sign == -1 && num < math.MinInt32 /10 || (num == math.MinInt32/10 && ((-1 * int32(s[i]-'0')) < math.MinInt32%10)){
			return math.MinInt32
		}
		num = num * 10 + int32(sign) * int32(s[i]-'0')
	}
	return int64(num)
}

