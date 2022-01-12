/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2022/1/11 2:43 下午
 * @Desc:
 */

package trecursion

import (
	"algorithm/utils/mqueue"
	"fmt"
	"testing"
)

//字符串的全排列，比如[a,b,c]
//0位置可以有[a,]  [b,]   [c,]
//1位置对应可以有[a,b][a,c]   [b.a][b,c]  [c,a][c,b]
//2位置只剩下最后一个，填上即可
//每次选择当前位置的字符时，都需要从后边选出

func permutation(str string, index int, queue *mqueue.MQueue) {
	if index == len(str)-1 {
		queue.PushBack(string(str))
		return
	}
	byteStr := []byte(str)
	for i := index; i < len(str); i++ {
		byteStr[i], byteStr[index] = byteStr[index], str[i]
		permutation(string(byteStr), index+1, queue)
	}
}

func TestPermutation(t *testing.T) {
	str := "abc"
	queue := mqueue.New()
	permutation(str, 0, queue)
	for item := queue.PopFront(); item != nil; item = queue.PopFront() {
		fmt.Println(string(item.(string)))
	}
}
