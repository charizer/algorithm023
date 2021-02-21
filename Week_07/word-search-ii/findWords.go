package word_search_ii

func FindWords(board [][]byte, words []string) []string {
	ans := []string{}
	if len(board) <= 0 || len(board[0]) <= 0 || len(words) <= 0 {
		return ans
	}
	// 向四个方向扩展
	directs := [][]int{{0,-1},{1,0},{0,1},{-1,0}}
	// 构建trie树
	trie := Constructor()
	for _, word := range words {
		trie.Insert(word)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(trie.root,i, j, directs, board, &ans)
		}
	}
	return ans
}

func dfs(node *TrieNode, row, col int, directs [][]int, board [][]byte, ans *[]string) {
	// 越界
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
		return
	}
	c := board[row][col]
	// 已经遍历过或者trie树中不存在此字符
	if c == '#' || node.Children[c-'a'] == nil {
		return
	}
	node = node.Children[c-'a']
	// 已经匹配到一个单词的结尾
	if node.Word != "" {
		*ans = append(*ans, node.Word)
		// 防重
		node.Word = ""
	}
	// 标记为已遍历
	board[row][col] = '#'
	// 向四个方向扩展下一个字符
	for _, dr := range directs {
		dfs(node,row + dr[0], col + dr[1], directs, board, ans)
	}
	// 标记清除
	board[row][col] = c
}

type TrieNode struct {
	// 孩子
	Children []*TrieNode
	// 单词
	Word  string
}

type Trie struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			Children: make([]*TrieNode, 26),
		},
	}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	p := this.root
	for _, ch := range word {
		// 没有孩子则创建
		if p.Children[ch-'a'] == nil {
			p.Children[ch-'a'] = &TrieNode{Children: make([]*TrieNode, 26)}
		}
		p = p.Children[ch-'a']
	}
	// 插入单词后标记单词
	p.Word = word
}
