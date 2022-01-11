/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/15 2:22 下午
 * @Desc:
 */

package ttree

import (
	"fmt"
	"testing"
)

//如何设计一个打印整棵树的打印函数
func printInOrder(head *Node, height int, to string, length int) {
	if head == nil {
		return
	}
	printInOrder(head.Right, height+1, "v", length)
	val := fmt.Sprintf("%s%d%s", to, head.Value, to)
	lenM := len(val)
	lenL := (length - lenM) / 2
	lenR := length - lenM - lenL
	val = fmt.Sprintf("%s%s%s", getSpace(lenL), val, getSpace(lenR))
	fmt.Println(fmt.Sprintf("%s%s", getSpace(height*length), val))
	printInOrder(head.Left, height+1, "^", length)
}
func getSpace(num int) string {
	var spaceStr []byte
	for num > 0 {
		spaceStr = append(spaceStr, ' ')
		num--
	}
	return string(spaceStr)
}
func TestPrintInOrder(t *testing.T) {
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
	printInOrder(n1, 1, "H", 5)
}
