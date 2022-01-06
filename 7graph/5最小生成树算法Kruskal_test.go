/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2022/1/5 10:51 上午
 * @Desc:
 */

package _graph

import (
	"fmt"
	"math"
	"testing"
)

//minimum spanning tree
//最小生成树算法之 Kruskal
//1.总是从权值最小的边开始考虑，依次考察权值依次变大的边
//2.当前的边要么进入最小生成树的集合，要么丢弃
//3.如果当前的边进入最小生成树的集合不会形成环，就要当前边
//4.如果当前的边进入最小生成树的集合会形成闭环，就不要当前边
//5.考察完所有边之后，最小生成树的集合也得到了
var graph *Graph

type Kruskal struct {
	g                   *Graph
	MinimumSpanningTree []*edge
}

func (k *Kruskal) Init(g *Graph) {
	if g != nil {
		k.g = g
	} else {
		fmt.Println("nil graph")
	}
}
func (k *Kruskal) BuildMinimumSpanningTree() {
	m := make(map[*node]bool)
	for _, e := range k.g.Edges {
		if !m[e.From] || !m[e.To] {
			k.MinimumSpanningTree = append(k.MinimumSpanningTree, e)
			m[e.From] = true
			m[e.To] = true
		}
	}
}
func (k *Kruskal) HeapSortEdges() {
	k.BuildSmallTopHeapByEdgeValue()
	edgeLen := len(k.g.Edges)
	for i := 0; i < edgeLen; i++ {
		k.g.Edges[0], k.g.Edges[edgeLen-1-i] = k.g.Edges[edgeLen-1-i], k.g.Edges[0]
		k.Heapify(0, edgeLen-1-i)
	}
}
func (k *Kruskal) BuildSmallTopHeapByEdgeValue() {
	if len(k.g.Edges) < 2 {
		return
	}
	index := len(k.g.Edges)/2 + 1
	for i := index; i >= 0; i-- {
		k.Heapify(i, len(k.g.Edges))
	}
}
func (k *Kruskal) Heapify(index, heapSize int) {
	if len(k.g.Edges) < 2 || index >= heapSize {
		return
	}
	left := 2*index + 1
	bigIndex := index
	for left < heapSize {
		if k.g.Edges[bigIndex].Value < k.g.Edges[left].Value {
			bigIndex = left
		}
		if left+1 < heapSize && k.g.Edges[bigIndex].Value < k.g.Edges[left+1].Value {
			bigIndex = left + 1
		}
		if bigIndex == index {
			break
		}
		k.g.Edges[index], k.g.Edges[bigIndex] = k.g.Edges[bigIndex], k.g.Edges[index]
		if left > math.MaxInt64/2 {
			break
		}
		index = bigIndex
		left = index*2 + 1
	}
}
func TestKruskal(t *testing.T) {
	var k Kruskal
	k.Init(initGraph())
	k.HeapSortEdges()
	k.BuildMinimumSpanningTree()
	fmt.Println(k.g.Edges)
}
func initGraph() *Graph {
	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)
	n5 := NewNode(5)
	n6 := NewNode(6)
	g := NewGraph()

	n1.NextNodes.PushBack(n2)
	n1.Out++
	n2.In++
	e1 := NewEdge(1, n1, n2)
	n1.Edges.PushBack(e1)
	n1.NextNodes.PushBack(n3)
	n1.Out++
	n3.In++
	e2 := NewEdge(3, n1, n3)
	n1.Edges.PushBack(e2)

	n2.NextNodes.PushBack(n5)
	e3 := NewEdge(1, n2, n5)
	n2.Edges.PushBack(e3)
	n2.Out++
	n5.In++

	n3.NextNodes.PushBack(n6)
	e4 := NewEdge(3, n3, n6)
	n3.Edges.PushBack(e4)
	n3.Out++
	n6.In++

	n5.NextNodes.PushBack(n3)
	e5 := NewEdge(2, n5, n3)
	n5.Edges.PushBack(e5)
	n5.Out++
	n3.In++

	n5.NextNodes.PushBack(n6)
	e6 := NewEdge(7, n5, n6)
	n5.Edges.PushBack(e6)
	n5.Out++
	n6.In++

	n3.NextNodes.PushBack(n1)
	e7 := NewEdge(1, n3, n1)
	n3.Edges.PushBack(e7)
	n3.Out++
	//n1.In++

	g.Nodes[1] = n1
	g.Nodes[2] = n2
	g.Nodes[3] = n3
	//g.Nodes[4] = n4
	g.Nodes[5] = n5
	g.Nodes[6] = n6
	g.Edges = append(g.Edges, e1)
	g.Edges = append(g.Edges, e2)
	g.Edges = append(g.Edges, e3)
	g.Edges = append(g.Edges, e4)
	g.Edges = append(g.Edges, e5)
	g.Edges = append(g.Edges, e6)
	g.Edges = append(g.Edges, e7)
	return g
}
