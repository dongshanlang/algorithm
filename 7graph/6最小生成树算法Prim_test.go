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

//思路解析
//1.随机选取一个点，解锁该点的所有边，进入小顶堆
//2.从已经解锁的所有边中选出最小的边（选中边），
//3.选中边的to节点如果没有筛选过，将所有边进入小顶堆，回到2
//2和3运行完回到1

//最小生成树Prim算法
type Prim struct {
	g                   *Graph
	unlockedEdgeHeap    []*edge
	MinimumSpanningTree []*edge
}

func (p *Prim) InitPrim(g *Graph) {
	if g != nil {
		p.g = g
	}
}
func (p *Prim) AddEdge(e *edge) {
	p.unlockedEdgeHeap = append(p.unlockedEdgeHeap, e)
	p.HeapInsert(len(p.unlockedEdgeHeap) - 1)
}
func (p *Prim) GetMinimumEdge() *edge {
	if len(p.unlockedEdgeHeap) < 1 {
		return nil
	}
	minimumEdge := p.unlockedEdgeHeap[0]
	if len(p.unlockedEdgeHeap) == 1 {
		p.unlockedEdgeHeap = nil
	} else {
		p.unlockedEdgeHeap[0], p.unlockedEdgeHeap[len(p.unlockedEdgeHeap)-1] = p.unlockedEdgeHeap[len(p.unlockedEdgeHeap)-1], p.unlockedEdgeHeap[0]
		p.unlockedEdgeHeap = p.unlockedEdgeHeap[:len(p.unlockedEdgeHeap)-1]
		p.Heapify(0, len(p.unlockedEdgeHeap))
	}

	return minimumEdge
}
func (p *Prim) HeapInsert(index int) {
	if index >= len(p.unlockedEdgeHeap) || len(p.unlockedEdgeHeap) < 2 {
		return
	}
	parent := (index - 1) / 2
	for parent >= 0 {
		if p.unlockedEdgeHeap[parent].Value > p.unlockedEdgeHeap[index].Value {
			p.unlockedEdgeHeap[parent], p.unlockedEdgeHeap[index] = p.unlockedEdgeHeap[index], p.unlockedEdgeHeap[parent]
		} else {
			break
		}
		index = parent
		parent = (index - 1) / 2
	}
}
func (p *Prim) Heapify(index, heapSize int) {
	if len(p.unlockedEdgeHeap) < 2 {
		return
	}
	left := index*2 + 1
	smallIndex := index
	for left < heapSize {
		if p.unlockedEdgeHeap[smallIndex].Value > p.unlockedEdgeHeap[left].Value {
			smallIndex = left
		}
		if left+1 < heapSize && p.unlockedEdgeHeap[smallIndex].Value > p.unlockedEdgeHeap[left+1].Value {
			smallIndex = left + 1
		}
		if smallIndex == index {
			return
		}
		p.unlockedEdgeHeap[smallIndex], p.unlockedEdgeHeap[index] = p.unlockedEdgeHeap[index], p.unlockedEdgeHeap[smallIndex]
		index = smallIndex
		if index > math.MaxInt64/2 {
			break
		}
		left = index*2 + 1
	}
}
func (p *Prim) BuildMinimumSpanningTree() {
	if p.MinimumSpanningTree != nil {
		p.MinimumSpanningTree = nil
	}
	overNodesMap := make(map[*node]bool)

	for _, nodes := range p.g.Nodes {
		if !overNodesMap[nodes] {
			overNodesMap[nodes] = true
			for item := nodes.Edges.PopFront(); item != nil; item = nodes.Edges.PopFront() {
				e := item.(*edge)
				p.AddEdge(e)
			}
			for len(p.unlockedEdgeHeap) != 0 {
				minEdge := p.GetMinimumEdge()
				if minEdge != nil {
					if !overNodesMap[minEdge.To] {
						p.MinimumSpanningTree = append(p.MinimumSpanningTree, minEdge)
						overNodesMap[minEdge.To] = true
						for item := minEdge.To.Edges.PopFront(); item != nil; item = minEdge.To.Edges.PopFront() {
							e := item.(*edge)
							p.AddEdge(e)
						}
					}
				}
			}
		}
		//break
	}
}

func TestPrim(t *testing.T) {
	var p Prim
	p.InitPrim(initGraph())
	p.BuildMinimumSpanningTree()
	for _, e := range p.MinimumSpanningTree {
		fmt.Printf("%d ", e.Value)
	}
}
