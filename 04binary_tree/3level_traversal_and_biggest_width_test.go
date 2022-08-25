/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/14 11:03 上午
 * @Desc:
 */

package ttree

import (
	"algorithm/utils/mqueue"
	"fmt"
	"testing"
)

//获取树的最大宽度
//实质是宽度优先遍历，借助队列，可以通过设置flag变量的方式，记录某层的结束
func getBiggestWidthOfTree(head *Node) int {
	if head == nil {
		return 0
	}
	var currentNode *Node
	var l = mqueue.New()
	l.PushBack(head)
	var currentLevelEnd = head
	var nextLevelEnd *Node = nil
	var max int
	var currentLevelNodeCount = 0
	for !l.IsEmpty() {
		currentNode = l.PopFront().(*Node)
		currentLevelNodeCount++
		if currentNode.Left != nil {
			nextLevelEnd = currentNode.Left
			l.PushBack(currentNode.Left)
		}
		if currentNode.Right != nil {
			nextLevelEnd = currentNode.Right
			l.PushBack(currentNode.Right)
		}
		if currentNode == currentLevelEnd {
			max = getMax(max, currentLevelNodeCount)
			currentLevelNodeCount = 0
			currentLevelEnd = nextLevelEnd
		}
	}
	return max
}
func getMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func TestGetMostWidthOfTree(t *testing.T) {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n4 := &Node{Value: 4}
	n5 := &Node{Value: 5}
	n6 := &Node{Value: 6}
	n7 := &Node{Value: 7}
	n8 := &Node{Value: 8}
	n9 := &Node{Value: 9}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	n5.Left = n8
	n5.Right = n9
	max := getBiggestWidthOfTree(n1)
	fmt.Println(max)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestLevelOrder(t *testing.T) {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}
	n6 := &TreeNode{Val: 6}
	n7 := &TreeNode{Val: 7}
	n8 := &TreeNode{Val: 8}
	n9 := &TreeNode{Val: 9}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	n5.Left = n8
	n5.Right = n9
	ret := levelOrder(n1)
	fmt.Println(ret)
}

func levelOrder(root *TreeNode) [][]int {
	// write code here
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	result = append(result, make([]int, 0))
	currentLevelEnd := root
	var nextLevelEnd *TreeNode
	currentLevel := 0
	nodeList := make([]*TreeNode, 0)
	nodeList = append(nodeList, root)

	for i := 0; i < len(nodeList); i++ {
		result[currentLevel] = append(result[currentLevel], nodeList[i].Val)
		if nodeList[i].Left != nil {
			nodeList = append(nodeList, nodeList[i].Left)
			nextLevelEnd = nodeList[i].Left
		}
		if nodeList[i].Right != nil {
			nodeList = append(nodeList, nodeList[i].Right)
			nextLevelEnd = nodeList[i].Right
		}
		if nodeList[i] == currentLevelEnd {
			currentLevel++
			currentLevelEnd = nextLevelEnd
			nextLevelEnd = nil
			result = append(result, make([]int, 0))
		}
	}
	fmt.Println(currentLevel)
	return result[0 : len(result)-1]
}

func TestLevelTranverse1(t *testing.T) {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}
	n6 := &TreeNode{Val: 6}
	n7 := &TreeNode{Val: 7}
	n8 := &TreeNode{Val: 8}
	n9 := &TreeNode{Val: 9}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	n5.Left = n8
	n5.Right = n9
	var orderNodes []*TreeNode
	cur := n1
	orderNodes = append(orderNodes, cur)
	var step int
	for cur != nil {
		if cur.Left != nil {
			orderNodes = append(orderNodes, cur.Left)
		}
		if cur.Right != nil {
			orderNodes = append(orderNodes, cur.Right)
		}
		step++
		if step >= len(orderNodes) {
			break
		}
		cur = orderNodes[step]
	}
	for _, v := range orderNodes {
		fmt.Println(v.Val)
	}
}
