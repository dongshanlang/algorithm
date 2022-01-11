/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/16 9:04 上午
 * @Desc:
 */

package ttree

import (
	"fmt"
	"math"
	"testing"
)

//给定一棵二叉树的头节点head，任何两个节点之间都存在距离，返回整棵树的最大距离
type DistanceInfo struct {
	BiggestDistance int
	Height          int
}

func GetDistanceInfo(head *Node) DistanceInfo {
	if head == nil {
		return DistanceInfo{
			BiggestDistance: 0,
			Height:          0,
		}
	}
	leftDistanceInfo := GetDistanceInfo(head.Left)
	rightDistanceInfo := GetDistanceInfo(head.Right)

	height := math.Max(float64(leftDistanceInfo.Height), float64(rightDistanceInfo.Height)) + 1
	distance := math.Max(math.Max(float64(leftDistanceInfo.BiggestDistance), float64(rightDistanceInfo.BiggestDistance)),
		float64(leftDistanceInfo.Height+1+rightDistanceInfo.Height))
	return DistanceInfo{
		BiggestDistance: int(distance),
		Height:          int(height),
	}
}

func TestGetDistanceInfo(t *testing.T) {
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
	fmt.Println(GetDistanceInfo(n1).BiggestDistance)
}
