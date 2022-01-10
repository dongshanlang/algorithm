/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2022/1/10 5:14 下午
 * @Desc:
 */

package _recursion

import (
	"fmt"
	"testing"
)

//这个问题是递归思想的典型代表。通常解决这样的问题我们可以有两个极端方向走：
//1）假设盘子个数为1,2等较小的数，
//2）假设盘子个数为n。当 n=1时，可以不借助辅助的柱子只一定一次就可以完成任务；当 n=2时，需要借助辅助的柱子移动三次完成任务。
//当 n较大时，我们可以设想如果 n个盘子的前 n−1个盘子都被移动到了辅助的柱子上，那么我们需要做的是把最大的那个盘子移动到B柱上，
//然后再把 n−1个盘子，按照之前的方法从辅助柱子移动到B柱上即可。

func MoveHanoi(n int, from, to, other string) {
	if n == 1 {
		fmt.Println(n, "move from ", from, " to ", to)
		return
	}
	MoveHanoi(n-1, from, other, to)
	fmt.Println(n, "move from ", from, " to ", to)
	MoveHanoi(n-1, other, to, from)
}
func TestMoveHanoi(t *testing.T) {
	MoveHanoi(7, "left", "right", "mid")
}
