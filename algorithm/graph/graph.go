package graph

import (
	"enmotech_utils/algorithm/queue"
	"fmt"
)

type Graph struct {
	steps  map[string]struct{}
	points []string
	matrix [][]int
	edegs  int

	queue *queue.ArrayQueue
}

func NewGraph(data []string) *Graph {
	n := len(data)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	//dataMap := make(map[string]int)
	//for i,v := range data{
	//	dataMap[v] = i
	//}
	return &Graph{
		points: data,
		matrix: matrix,
		steps:  make(map[string]struct{}),
		queue:  queue.NewArrayQueue(),
	}
}

func (g *Graph) AddPoint(point string) {
	g.points = append(g.points, point)
	n := len(g.points)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == n-1 || j == n-1 {
				matrix[i][j] = 0
			} else {
				matrix[i][j] = g.matrix[i][j]
			}

		}
	}
	g.matrix = matrix
}

func (g *Graph) InsertEdge(i, j, weight int) {
	g.matrix[i][j] = weight
	g.matrix[j][i] = weight
	g.edegs++
}

func (g *Graph) InsertEdgeByPoint(m, n string, weight int) {
	i, j := g.GetIndex(m), g.GetIndex(n)
	g.matrix[i][j] = weight
	g.matrix[j][i] = weight
	g.edegs++
}

func (g *Graph) GetEdgeNums() int {
	return g.edegs
}

func (g *Graph) GetValue(index int) string {
	return g.points[index]
}

func (g *Graph) CheckEdge(x, y string) bool {

	i, j := g.GetIndex(x), g.GetIndex(y)

	if g.matrix[i][j] > 0 {
		return true
	}
	return false
}

func (g *Graph) GetIndex(value string) int {
	for i, v := range g.points {
		if v == value {
			return i
		}
	}
	return -1
}

func (g *Graph) GetWeight(v1, v2 int) int {
	return g.matrix[v1][v2]
}

func (g *Graph) GetPointNums() int {
	return len(g.points)
}

func (g *Graph) ShowGraph() {
	for i := 0; i < len(g.matrix); i++ {
		for j := 0; j < len(g.matrix); j++ {
			fmt.Print(g.matrix[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func (g *Graph) DFS() {
	if len(g.points) == 0 {
		return
	}
	g.dfs(g.points[0])
}

func (g *Graph) BFS() {
	if len(g.points) == 0 {
		return
	}
	g.queue.Add(g.points[0])
	for !g.queue.IsEmpty() {
		m := g.queue.Get()
		fmt.Println(m)
		g.steps[m.(string)] = struct{}{}
		g.bfs(m.(string))
	}
}

func (g *Graph) dfs(p string) {
	if _, ok := g.steps[p]; !ok {
		fmt.Println(p)
		g.steps[p] = struct{}{}
	}

	for i := 0; i < len(g.points); i++ {
		if n := g.getNeighbor(p, i); n != "" && g.CheckEdge(p, n) {
			//fmt.Println(g.points[i])
			g.dfs(n)
		}
	}

	//g.dfs(p)
	//g.points[]
	//if

}

func (g *Graph) bfs(p string) {

	if _, ok := g.steps[p]; !ok {
		g.steps[p] = struct{}{}
	}

	for i := 0; i < len(g.points); i++ {
		if n := g.getNeighbor(p, i); n != "" && g.CheckEdge(p, n) {
			g.steps[g.points[i]] = struct{}{}
			//fmt.Println()
			g.queue.Add(g.points[i])
			//tmp = append(tmp,i)
		}
	}

	//fmt.Println(g.queue.Get())
	//
	//tmp := []int{}

	//
	//for i:=0;i<len(tmp);i++{
	//		g.bfs(g.points[tmp[i]])
	//}
	////g.dfs(p)
	////g.points[]
	////if

}

func (g *Graph) getNeighbor(p string, i int) string {
	index := g.GetIndex(p)
	if _, ok := g.steps[g.points[i]]; g.matrix[index][i] > 0 && !ok {
		return g.points[i]
	}
	return ""
}
