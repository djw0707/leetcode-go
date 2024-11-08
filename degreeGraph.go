package main

import (
	"fmt"
)

type Graph struct {
	edge   [][]int
	degree []int
}

var visited []bool
var ans [][]int

func (g *Graph) isConnected(i int, j int) bool {
	if g.bfsFind(i, j) == true {
		return true
	}
	return false
}

func (g *Graph) dfsGraph(i int, arr *[]int) {
	if len(g.edge[i]) == 0 && len(*arr) > 0 {
		// 最后一个元素如何停止
		*arr = append(*arr, i)
		tmp := make([]int, len(*arr))
		copy(tmp, *arr)
		ans = append(ans, tmp)
		*arr = (*arr)[:len(*arr)-1]
		return
	}
	// 深度优先搜索加回溯
	visited[i] = true
	*arr = append(*arr, i)
	for j := 0; j < len(g.edge[i]); j++ {
		cur := g.edge[i][j]
		if visited[cur] == false {
			g.dfsGraph(cur, arr)
		}
	}
	visited[i] = false
	*arr = (*arr)[:len(*arr)-1]
}

func (g *Graph) bfsFind(start int, target int) bool {
	queue := make([]int, 0)
	queue = append(queue, start)
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			cur := queue[i]
			if cur == target {
				return true
			}
			arr := g.edge[cur]
			for j := 0; j < len(arr); j++ {
				queue = append(queue, arr[j])
			}
		}
		queue = queue[l:]
	}
	return false
}

func buildGraph(n int) Graph {
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}
	return Graph{
		edge:   graph,
		degree: make([]int, n),
	}
}

func (g *Graph) insertEdge(i int, j int) bool {
	if g.isConnected(j, i) {
		return false
	}
	g.edge[i] = append(g.edge[i], j)
	g.degree[j]++
	return true
}

func main() {
	g := buildGraph(10)
	g.insertEdge(0, 1)
	g.insertEdge(1, 2)
	g.insertEdge(2, 3)
	g.insertEdge(3, 1)
	g.insertEdge(1, 4)
	g.insertEdge(2, 5)
	visited = make([]bool, len(g.degree))
	ans = make([][]int, 0)
	arr := make([]int, 0)
	fmt.Println(g)
	for i := 0; i < len(g.degree); i++ {
		if g.degree[i] == 0 && len(g.edge[i]) > 0 {
			g.dfsGraph(i, &arr)
		}
	}
	fmt.Println(ans)
}
