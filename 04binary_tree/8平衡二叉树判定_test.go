/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/15 6:20 下午
 * @Desc:
 */

package _4binary_tree

import (
	"fmt"
	"math"
	"testing"
)

//给定一棵树头节点，判定是否是平衡二叉树
//分析：
//平衡二叉树：左子树为平衡二叉树，右子树为平衡二叉树，左子树与右子树的高度差不大于1
type BalanceInfo struct {
	IsBalance bool
	Height    int
}

func judgeBalanceBinaryTree(head *Node) BalanceInfo {
	if head == nil {
		return BalanceInfo{IsBalance: true, Height: 0}
	}
	leftInfo := judgeBalanceBinaryTree(head.Left)
	rightInfo := judgeBalanceBinaryTree(head.Right)
	isBalance := true
	if !leftInfo.IsBalance || !rightInfo.IsBalance || math.Abs(float64(leftInfo.Height-rightInfo.Height)) > 1 {
		isBalance = false
	}
	height := math.Max(float64(leftInfo.Height), float64(rightInfo.Height)) + 1
	return BalanceInfo{
		IsBalance: isBalance,
		Height:    int(height),
	}
}
func TestJudgeBalanceBinaryTree(t *testing.T) {
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
	n8.Right = n9
	fmt.Println(judgeBalanceBinaryTree(n1))
}
