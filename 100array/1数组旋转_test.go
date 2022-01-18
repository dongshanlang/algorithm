/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2022/1/18 9:55 上午
 * @Desc:
 */

package tarray

import (
	"fmt"
	"testing"
)

//搜索旋转数组。给定一个排序后的数组，包含n个整数，但这个数组已被旋转过很多次了，次数不详。请编写代码找出数组中的某个元素，假设数组元素原先是按升序排列的。若有多个相同元素，返回索引值最小的一个。
//
var (

	//示例1:
	//
	arr1    = []int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}
	target1 = 5
	//输出: 8（元素5在该数组中的索引）
	//示例2:
	//
	//输入：
	arr2    = []int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}
	target2 = 11

	//输出：-1 （没有找到）
	//输入：
	arr3    = []int{1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1}
	target3 = 2
)

func TestArray(t *testing.T) {
	fmt.Println(searchMinIndex(arr3))
}
func searchMinIndex(arr []int) int {
	n := len(arr)
	for n > 1 && arr[n-1] == arr[0] {
		n--
	}
	left := 0
	right := n - 1
	for left < right {
		mid := (left + right) / 2
		if arr[mid] <= arr[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
