/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/16 10:14 上午
 * @Desc:
 */

package _4binary_tree

import (
	"fmt"
	"math"
	"testing"
)

//给定一棵树，找到最大的二叉搜索树

//分析：问题一分为二，包含当前节点，不包含当前节点
//包含当前节点的问题，
type SearchTreeInfo struct {
	Min, Max               int
	IsBinarySearchTree     bool
	MaxSubBinarySearchTree int
	BNode                  *Node
}

func getSearchTreeInfo(head *Node) *SearchTreeInfo {
	if head == nil {
		return nil
	}
	leftInfo := getSearchTreeInfo(head.Left)
	rightInfo := getSearchTreeInfo(head.Right)

	min := head.Value
	max := head.Value
	if leftInfo != nil {
		min = int(math.Min(float64(min), float64(leftInfo.Min)))
		max = int(math.Max(float64(max), float64(leftInfo.Max)))
	}
	if rightInfo != nil {
		min = int(math.Min(float64(min), float64(rightInfo.Min)))
		max = int(math.Max(float64(max), float64(rightInfo.Max)))
	}
	isBST := false
	maxSBST := 0
	var bNode *Node
	if leftInfo != nil {
		maxSBST = int(math.Max(float64(leftInfo.MaxSubBinarySearchTree), float64(maxSBST)))
		bNode = leftInfo.BNode
	}
	if rightInfo != nil {
		maxSBST = int(math.Max(float64(rightInfo.MaxSubBinarySearchTree), float64(maxSBST)))
		bNode = rightInfo.BNode
	}
	if rightInfo != nil && leftInfo != nil {
		if leftInfo.MaxSubBinarySearchTree > rightInfo.MaxSubBinarySearchTree {
			bNode = leftInfo.BNode
		} else {
			bNode = rightInfo.BNode
		}
	}

	if (leftInfo == nil || leftInfo.IsBinarySearchTree) &&
		(rightInfo == nil || rightInfo.IsBinarySearchTree) &&
		(rightInfo == nil || rightInfo.Min > head.Value) &&
		(leftInfo == nil || leftInfo.Max < head.Value) {
		isBST = true
		maxSBST = 0
		if leftInfo != nil {
			maxSBST += leftInfo.MaxSubBinarySearchTree
		}
		if rightInfo != nil {
			maxSBST += rightInfo.MaxSubBinarySearchTree
		}
		maxSBST += 1
		bNode = head
	}
	return &SearchTreeInfo{
		Min:                    min,
		Max:                    max,
		IsBinarySearchTree:     isBST,
		MaxSubBinarySearchTree: maxSBST,
		BNode:                  bNode,
	}
}
func TestGetBiggestSubSearchBinaryTree(t *testing.T) {
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
	info := getSearchTreeInfo(n1)
	fmt.Println(info)
	fmt.Println(info.BNode.Value)
}
