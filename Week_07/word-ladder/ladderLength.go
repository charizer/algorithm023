package word_ladder

func LadderLength(beginWord string, endWord string, wordList []string) int {
	ans := 0
	wordMap := make(map[string]bool)
	for _, word := range wordList {
		wordMap[word] = true
	}
	// 字符串中不包括最后的单词，则肯定达不到目标，提前结束
	if !wordMap[endWord] {
		return ans
	}
	beginMap, endMap, visitedMap := make(map[string]bool), make(map[string]bool), make(map[string]bool)
	beginMap[beginWord] = true
	endMap[endWord] = true
	for len(beginMap) > 0 && len(endMap) > 0 {
		ans++
		// 元素少的一端开始处理
		if len(beginMap) > len(endMap) {
			beginMap, endMap = endMap, beginMap
		}
		// 记录下一层bfs元素
		tmp := make(map[string]bool)
		for word, _ := range beginMap {
			bw := []byte(word)
			// 依次变单词的每一位
			for i := 0; i < len(bw); i++ {
				old := bw[i]
				// 字符替换，从 a ~ z
				for j := 0; j < 26; j++ {
					bw[i] = byte('a' + j)
					sw := string(bw)
					// 两端碰头，则满足条件
					if endMap[sw] {
						return ans + 1
					}
					// 如果这个单词还没处理，并且变到的当前单词是字典里的词
					if !visitedMap[sw] && wordMap[sw] {
						tmp[sw] = true
						// 标记已访问
						visitedMap[sw] = true
					}
				}
				// 恢复被替换的字符
				bw[i] = old
			}
			// 下一层bfs处理元素
			beginMap = tmp
		}
	}
	// 不能变换
	return 0
}
