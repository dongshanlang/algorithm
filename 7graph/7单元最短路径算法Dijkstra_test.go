/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2022/1/5 10:52 上午
 * @Desc:
 */

package _graph

import (
	"fmt"
	"math"
	"testing"
)

//Dijkstra算法要求边的权值不能为负
//从一个顶点到其余各顶点的最短路径算法，解决的是有权图中最短路径问题
//迪杰斯特拉算法主要特点是从起始点开始，采用贪心算法的策略，每次遍历到始点距离最近且未访问过的顶点的邻接节点，直到扩展到终点为止
type Dijkstra struct {
	DistanceMap      map[*node]int
	selectedNodesMap map[*node]bool
}

func NewDijkstra() *Dijkstra {
	return &Dijkstra{
		DistanceMap:      make(map[*node]int),
		selectedNodesMap: make(map[*node]bool),
	}
}
func (d *Dijkstra) BuildDistanceMap(from *node) {
	d.DistanceMap[from] = 0
	minNode := d.getMinimumDistanceAndUnselectedNode()
	for minNode != nil {
		for item := minNode.Edges.PopFront(); item != nil; item = minNode.Edges.PopFront() {
			edge := item.(*edge)
			if _, in := d.DistanceMap[edge.To]; !in {
				d.DistanceMap[edge.To] = edge.Value + d.DistanceMap[edge.From]
			} else {
				d.DistanceMap[edge.To] = int(math.Min(float64(d.DistanceMap[edge.To]), float64(d.DistanceMap[edge.From]+edge.Value)))
			}
		}
		d.selectedNodesMap[minNode] = true
		minNode = d.getMinimumDistanceAndUnselectedNode()
	}
}

func (d *Dijkstra) getMinimumDistanceAndUnselectedNode() *node {
	var nodeMin *node
	minDistance := math.MaxInt64
	for node, distance := range d.DistanceMap {
		if distance < minDistance && !d.selectedNodesMap[node] {
			minDistance = distance
			nodeMin = node
		}
	}
	return nodeMin
}
func TestDijkstra(t *testing.T) {
	g := initGraph()
	var d = NewDijkstra()
	d.BuildDistanceMap(g.Nodes[1])
	fmt.Println(d)
}
