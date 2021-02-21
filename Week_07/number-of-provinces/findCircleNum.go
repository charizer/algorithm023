package number_of_provinces

func FindCircleNum(isConnected [][]int) int {
	if len(isConnected) <= 0 {
		return 0
	}
	n := len(isConnected)
	// 初始化并查集
	union := NewUnionFind(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 {
				// 连通并查集
				union.Union(i,j)
			}
		}
	}
	return union.Count()
}

// 构造并查集结构
type UnionFind struct {
	// 连通分量
	count int
	// 节点的父节点
	parent []int
}

// 初始化
func NewUnionFind(n int) *UnionFind {
	u := &UnionFind{
		count: n,
		parent: make([]int, n),
	}
	for i := 0; i < n; i++ {
		u.parent[i] = i
	}
	return u
}

// 合并p，q
func (u *UnionFind) Union(p, q int){
	rootP := u.Find(p)
	rootQ := u.Find(q)
	if rootP == rootQ {
		return
	}
	u.parent[rootQ] = rootP
	u.count--
}

// 查找x的祖先
func (u *UnionFind) Find(x int) int {
	for u.parent[x] != x {
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}

// 返回连通分量个数
func (u *UnionFind) Count() int {
	return u.count
}


