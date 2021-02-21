package number_of_islands

func NumIslands(grid [][]byte) int {
	ans := 0
	if len(grid) <= 0 || len(grid[0]) <= 0 {
		return ans
	}
	// 水格子数量
	spaces := 0
	directs := [][]int{{1,0},{0,1},{-1,0},{0,-1}}
	// 初始时，连通量 union.count = rows * cols, 处理过程中记下水格子数量，最后用连通量-水格子数量就是岛屿
	union := NewUnionFind(len(grid)*len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				// 向四个方向连通
				for _, dr := range directs {
					dx := i + dr[0]
					dy := j + dr[1]
					if dx >= 0 && dx < len(grid) && dy >= 0 && dy < len(grid[0]) && grid[dx][dy] == '1' {
						union.Union(i * len(grid[0]) + j, dx * len(grid[0]) + dy)
					}
				}
			}else{
				// 是水，数量+1
				spaces += 1
			}
		}
	}
	// 岛屿=连通量-水格子
	return union.count - spaces
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