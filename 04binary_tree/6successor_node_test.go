/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/15 2:52 下午
 * @Desc:
 */

package ttree

import (
	"fmt"
	"testing"
)

//二叉树结构定义如下
type PNode struct {
	Value               int
	Left, Right, Parent *PNode
}

//给定二叉树中的某个节点，返回该节点的后继节点

//分析：
//后继节点，中序遍历顺序中一个节点的下一个节点就是后继节点
//1。暴力解决，找到父节点，中序遍历找到某节点后继节点
//2。通过观察：
//      1）如果给定节点右子树不为空，则后继节点为右子树最左侧节点
//      2）如果右子树为空，则由当前节点网上找node，终止条件：该node的父节点不为空，并且，node是父节点的左孩子，则node的父节点就是给定节点的后继节点
//前驱节点，中序遍历顺序中的一个节点的前一个节点就是前驱节点
func getSuccessorNode(node *PNode) *PNode {
	if node == nil {
		return node
	}
	var retNode *PNode
	if node.Right != nil {
		retNode = node.Right
		for retNode.Left != nil {
			retNode = retNode.Left
		}
		return retNode
	} else {
		parentNode := node.Parent
		curNode := node
		for parentNode != nil && curNode == parentNode.Right {
			curNode = parentNode
			parentNode = parentNode.Parent
		}
		return parentNode
	}
}
func TestGetSuccessorNode(t *testing.T) {
	n1 := &PNode{Value: 1}
	n2 := &PNode{Value: 2}
	n3 := &PNode{Value: 3}
	n4 := &PNode{Value: 4}
	n5 := &PNode{Value: 5}
	n6 := &PNode{Value: 6}
	n7 := &PNode{Value: 7}
	n8 := &PNode{Value: 8}
	n9 := &PNode{Value: 9}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5

	n3.Left = n6
	n3.Right = n7
	n3.Parent = n1
	n5.Left = n8
	n5.Right = n9

	n1.Parent = nil
	n2.Parent = n1
	n3.Parent = n1
	n4.Parent = n2
	n5.Parent = n2
	n6.Parent = n3
	n7.Parent = n3
	n8.Parent = n5
	n9.Parent = n5

	fmt.Println(getSuccessorNode(n7))
}
