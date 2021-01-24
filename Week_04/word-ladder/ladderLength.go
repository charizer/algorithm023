package word_ladder

// 在leetcode上耗时1000ms+
func LadderLength(beginWord string, endWord string, wordList []string) int {
	// 存放字典列表
	wm := make(map[string]bool)
	for _, word := range wordList {
		wm[word] = true
	}
	// bfs模板
	queue := []string{beginWord}
	ans := 0
	for len(queue) > 0 {
		ans++
		size := len(queue)
		for i := 0; i < size; i++ {
			// 出队列
			word := queue[0]
			queue = queue[1:]
			// 已经和最终单词相同，则返回值最小转换次数
			if word == endWord {
				return ans
			}
			// 已经变换过的从字典中删除
			delete(wm, word)
			// 单词的每位依次替换为：a~z
			for j := 0; j < len(word); j++ {
				bw := []byte(word)
				c := bw[j]
				for k := 0; k < 26; k++ {
					bw[j] = byte(int('a') + k)
					// 替换后的单词在字典列表中才入队
					if wm[string(bw)] {
						queue = append(queue, string(bw))
					}
				}
				bw[j] = c
			}
		}
	}
	return 0
}

// 优化：双向bfs, 优化后耗时几十ms
func LadderLength2(beginWord string, endWord string, wordList []string) int {
	ans := 0
	// 存放字典列表
	wm := make(map[string]bool)
	for _, word := range wordList {
		wm[word] = true
	}
	// 字典中不包括最后的单词，则肯定达不到目标，提前结束
	if !wm[endWord] {
		return ans
	}
	// 头部、尾部、已经访问的元素map
	beginMap, endMap, visited := make(map[string]bool), make(map[string]bool), make(map[string]bool)
	beginMap[beginWord] = true
	endMap[endWord] = true
	for len(beginMap) > 0 && len(endMap) > 0 {
		ans++
		// 元素少的一端开始处理
		if len(beginMap) > len(endMap) {
			beginMap, endMap = endMap, beginMap
		}
		tmp := make(map[string]bool)
		// 替换并判断
		for word, _ := range beginMap {
			bw := []byte(word)
			for i := 0; i < len(bw); i++ {
				old := bw[i]
				for j := 0; j < 26; j++ {
					bw[i] = byte(int('a') + j)
					sw := string(bw)
					if endMap[sw] {
						return ans + 1
					}
					if !visited[sw] && wm[sw] {
						tmp[sw] = true
						visited[sw] = true
					}
				}
				bw[i] = old
			}
		}
		beginMap = tmp
	}
	return 0
}