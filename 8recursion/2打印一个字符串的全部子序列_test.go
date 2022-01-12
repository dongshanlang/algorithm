/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2022/1/11 10:59 上午
 * @Desc:
 */

package trecursion

import (
	"algorithm/utils/mqueue"
	"fmt"
	"testing"
)

//1。打印一个字符串的全部子序列
func printSubSequenceString(str string) {
	q := mqueue.New()
	process(str, 0, q, "")
	for item := q.PopFront(); item != nil; item = q.PopFront() {
		fmt.Println(item.(string))
	}
}
func process(str string, index int, queue *mqueue.MQueue, path string) {
	if index == len(str) {
		queue.PushBack(path)
		return
	}
	no := path
	process(str, index+1, queue, no)
	yes := fmt.Sprintf("%s%c", path, str[index])
	process(str, index+1, queue, yes)
}
func TestPrintSequenceString(t *testing.T) {
	printSubSequenceString("hh")
}
