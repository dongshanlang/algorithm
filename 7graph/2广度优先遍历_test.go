/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/31 2:30 下午
 * @Desc:
 */

package _graph

import (
	"algorithm/utils/mqueue"
	"fmt"
	"testing"
)

var (
	n1 = NewNode(1)
	n2 = NewNode(2)
	n3 = NewNode(3)
	n5 = NewNode(5)
	n6 = NewNode(6)
	g  = NewGraph()
)

func TestBreadthFirstSearch(t *testing.T) {
	InitGraph()
	q := mqueue.New()
	m := make(map[*node]bool)
	m[n1] = true
	q.PushBack(n1)
	for !q.IsEmpty() {
		curNode := q.PopFront().(*node)
		fmt.Println(curNode.Value)
		for next := curNode.NextNodes.PopFront(); next != nil; next = curNode.NextNodes.PopFront() {
			nextNode := next.(*node)
			if !m[nextNode] { //确保不会重复进队列
				q.PushBack(next)
				m[nextNode] = true
			}
		}
	}
}
func InitGraph() {

	n1.NextNodes.PushBack(n2)
	n1.Out++
	n2.In++
	n1.Edges.PushBack(NewEdge(1, n1, n2))
	n1.NextNodes.PushBack(n3)
	n1.Out++
	n3.In++
	n1.Edges.PushBack(NewEdge(3, n1, n3))

	n2.NextNodes.PushBack(n5)
	n2.Edges.PushBack(NewEdge(1, n2, n5))
	n2.Out++
	n5.In++

	n3.NextNodes.PushBack(n6)
	n3.Edges.PushBack(NewEdge(3, n3, n6))
	n3.Out++
	n6.In++

	n5.NextNodes.PushBack(n3)
	n5.Edges.PushBack(NewEdge(2, n5, n3))
	n5.Out++
	n3.In++

	n5.NextNodes.PushBack(n6)
	n5.Edges.PushBack(NewEdge(2, n5, n6))
	n5.Out++
	n6.In++

	n3.NextNodes.PushBack(n1)
	n3.Edges.PushBack(NewEdge(1, n3, n1))
	n3.Out++
	//n1.In++

	g.Nodes[1] = n1
	g.Nodes[2] = n2
	g.Nodes[3] = n3
	//g.Nodes[4] = n4
	g.Nodes[5] = n5
	g.Nodes[6] = n6

}
