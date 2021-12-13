/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/13 6:28 下午
 * @Desc:
 */

package mstack

import (
	"fmt"
	"testing"
)

type TsNode struct {
	Value int
}

func TestMyStack(t *testing.T) {
	s := New()
	s.Push(&TsNode{Value: 1})
	s.Push(&TsNode{Value: 2})
	s.Push(&TsNode{Value: 3})
	s.Push(&TsNode{Value: 4})
	s.Push(&TsNode{Value: 5})

	for item := s.Pop(); item != nil; item = s.Pop() {
		fmt.Println(item.(*TsNode).Value)
	}
}
