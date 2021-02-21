package implement_trie_prefix_tree

type TrieNode struct {
	// 孩子
	children []*TrieNode
	// 是否是单词
	IsWord  bool
}

type Trie struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			children: make([]*TrieNode, 26),
		},
	}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	p := this.root
	for _, ch := range word {
		// 没有孩子则创建
		if p.children[ch-'a'] == nil {
			p.children[ch-'a'] = &TrieNode{children: make([]*TrieNode, 26)}
		}
		p = p.children[ch-'a']
	}
	// 插入单词后标记单词
	p.IsWord = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	p := this.root
	for _, ch := range word {
		if p.children[ch-'a'] == nil {
			return false
		}
		p = p.children[ch-'a']
	}
	// 最后是单词结尾则已经查到，否则只是一个前缀
	return p.IsWord
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	p := this.root
	for _, ch := range prefix {
		if p.children[ch-'a'] == nil {
			return false
		}
		p = p.children[ch-'a']
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
