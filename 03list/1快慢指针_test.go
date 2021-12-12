/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/10 3:27 下午
 * @Desc:
 */

package _3list

import (
	"testing"
)

/**
 * 快慢指针
 * 1。输入链表头节点，奇数长度返回中点，偶数长度返回上中点
 * 2。输入链表头节点，奇数长度返回中点，偶数长度返回下中点1	A
 * 3。输入链表头节点，奇数长度返回中点前一个，偶数长度返回上中点前一个
 * 4。输入链表头节点，奇数长度返回中点前一个，偶数长度返回下中点前一个
 */
func TestMidOrUpMid(t *testing.T) {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n4 := &Node{Value: 4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	getMidOrUpMid(n1)
}
func getMidOrUpMid(node *Node) {

}
