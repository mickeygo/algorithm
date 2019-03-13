package graph

// DepthFirstSearch 深度优先搜索
type DepthFirstSearch struct {
	// 标记已搜索的顶点，其中顶点加入的顺序可作为搜索的顺序
	marked []int
	count  int
}

// NewDFS new an dfs
func NewDFS() *DepthFirstSearch {
	return &DepthFirstSearch{}
}

// DFS 深度优先搜索
// `v` 搜索的起点
func (s *DepthFirstSearch) DFS(g *Graph, v int) {
	if v < 0 || v > g.V() {
		panic("v must betwwen 0 and g.V")
	}

	s.marked = append(s.marked, v)
	s.count++

	for _, adj := range g.Adj(v) {
		if indexOf(s.marked, adj) != -1 {
			s.DFS(g, adj)
		}
	}
}

func indexOf(s []int, value int) int {
	for i := 0; i < len(s); i++ {
		if s[i] == value {
			return i
		}
	}

	return -1
}
