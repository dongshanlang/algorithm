/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2022/1/4 9:49 上午
 * @Desc:
 */

package _graph

import (
	"algorithm/utils/mstack"
	"fmt"
	"testing"
)

func DepthFirstSearch(n *node) {
	if n == nil {
		return
	}
	s := mstack.New()
	m := make(map[*node]bool)
	m[n] = true
	fmt.Println(n.Value)
	s.Push(n)
	var currentNode *node
	for !s.IsEmpty() {
		currentNode = s.Pop().(*node)
		for next := currentNode.NextNodes.PopFront(); next != nil; next = currentNode.NextNodes.PopFront() {
			nextNode := next.(*node)
			if !m[nextNode] {
				fmt.Println(nextNode.Value)
				m[nextNode] = true
				s.Push(currentNode)
				s.Push(nextNode)
				break
			}
		}
	}
}
func TestDepthFirstSearch(t *testing.T) {
	InitGraph()
	DepthFirstSearch(n1)
}
