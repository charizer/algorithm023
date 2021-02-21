package minimum_genetic_mutation

// 和单词接龙一样双端bfs
func MinMutation(start string, end string, bank []string) int {
	if len(start) != len(end) {
		return -1
	}
	bankMap := make(map[string]bool)
	for _, b := range bank {
		bankMap[b] = true
	}
	// 如果基因序列不存在最后的目标基因，则提前退出
	if !bankMap[end] {
		return -1
	}
	ans := 0
	// 双端bfs
	beginMap, endMap, visited := make(map[string]bool), make(map[string]bool), make(map[string]bool)
	beginMap[start] = true
	endMap[end] = true
	changeChs := []byte{'A','C','G','T'}
	for len(beginMap) > 0 && len(endMap) > 0 {
		if len(beginMap) > len(endMap) {
			beginMap, endMap = endMap, beginMap
		}
		tmpMap := make(map[string]bool)
		for word, _ := range beginMap {
			bw := []byte(word)
			// 一次替换序列的每一位
			for i := range bw {
				old := bw[i]
				for _, ch := range changeChs {
					bw[i] = ch
					sw := string(bw)
					// 提前找到结果
					if endMap[sw] {
						return ans + 1
					}
					if !visited[sw] && bankMap[sw] {
						visited[sw] = true
						tmpMap[sw] = true
					}
				}
				bw[i] = old
 			}
		}
		beginMap = tmpMap
		ans++
	}
	return -1
}
