// Package unionfind union-find 算法
package unionfind

type QuickFind struct {
	id    []int // 分量 id
	count int   // 分量数量
}

// Compute
func (uf *QuickFind) Compute(N int, pairs [][2]int) {
	uf.count = N

	uf.id = make([]int, N)
	for i := 0; i < N; i++ {
		uf.id[i] = i
	}

	for _, pair := range pairs {
		if !uf.Connected(pair[0], pair[1]) {
			uf.union(pair[0], pair[1])
		}
	}
}

// Connected 判断 p、q 之间是否在同一份量中
func (uf *QuickFind) Connected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

// Find p (0 到 N-1) 所在的分量的标识符, O(1)
func (uf *QuickFind) find(p int) int {
	return uf.id[p]
}

// Union 将 p 与 q 之间添加一条连接, O(N)
func (uf *QuickFind) union(p, q int) {
	pID := uf.find(p)
	qID := uf.find(q)

	if pID == qID {
		return
	}

	for i := 0; i < len(uf.id); i++ {
		if uf.id[i] == pID {
			uf.id[i] = qID
		}
	}

	uf.count--
}

// Count 连通分量的数量
func (uf *QuickFind) Count() int {
	return uf.count
}
