package unionfind

type QuickUnion struct {
	id    []int // 分量 id
	count int   // 分量数量
}

// find 时间复杂度有 tree 的 height 决定
func (uf *QuickUnion) find(p int) int {
	// 迭代算法，寻找 p 所在组的根节点 root
	for p != uf.id[p] {
		p = uf.id[p]
	}

	return p
}

// union 时间复杂度有 tree 的 height 决定
func (uf *QuickUnion) union(p, q int) {
	pRoot := uf.find(p)
	qRoot := uf.find(q)
	if pRoot == qRoot {
		return
	}
	uf.id[pRoot] = qRoot // 将 p 组的节点归并到 q 节点下
	uf.count--
}
