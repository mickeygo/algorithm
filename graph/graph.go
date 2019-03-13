package graph

import (
	"strconv"
	"strings"
)

// NEWLINE \r\n
const NEWLINE = "\r\n"

// Graph 无向图
type Graph struct {
	// 图中的节点个数
	vertex int
	// 图中的边的条数
	eage int
	// 与节点相邻的节点集合
	adjacencies map[int][]int
}

// NewGraph 初始化一个新的无向图
func NewGraph(in map[int][]int) *Graph {
	graph := Graph{
		vertex:      len(in),
		adjacencies: make(map[int][]int),
	}

	for k, v := range in {
		for _, w := range v {
			graph.AddEage(k, w)
		}
	}

	return &graph
}

// AddEage 在已有的节点间新增一条边 v-w
// {0 <= `v` < V} and {0 <= `w` < V}
func (g *Graph) AddEage(v, w int) {
	g.validateVertex(v)
	g.validateVertex(w)

	if v == w {
		return
	}

	g.eage++
	g.adjacencies[v] = append(g.adjacencies[v], w)
	g.adjacencies[w] = append(g.adjacencies[w], v)
}

// V 获取图中的节点个数
func (g *Graph) V() int {
	return g.vertex
}

// E 获取图中的边的条数
func (g *Graph) E() int {
	return g.eage
}

// Adj 获取与节点相邻的节点集合
func (g *Graph) Adj(v int) []int {
	return g.adjacencies[v]
}

func (g *Graph) String() string {
	str := strings.Builder{}
	str.WriteString(strconv.Itoa(g.vertex) + " vertices, " + strconv.Itoa(g.eage) + " edges " + NEWLINE)

	for k, vertexs := range g.adjacencies {
		str.WriteString(strconv.Itoa(k) + ": ")
		for _, v := range vertexs {
			str.WriteString(strconv.Itoa(v) + " ")
		}

		str.WriteString(NEWLINE)
	}

	return str.String()
}

func (g *Graph) validateVertex(v int) {
	if v < 0 || v > g.vertex {
		panic("v must betwwen 0 and g.V")
	}
}
