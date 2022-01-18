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
)

func TestArray(t *testing.T) {
	fmt.Println(search(arr1, target1))
	fmt.Println(search(arr2, target2))
}
func search(arr []int, target int) int {
	return recursiveSearch(arr, 0, len(arr)-1, target)
}
func recursiveSearch(arr []int, left, right, target int) int {
	if left == right {
		if arr[left] == target {
			return left
		}
		return -1
	}
	mid := left + (right-left)/2
	if arr[left] < arr[mid] && target > arr[left] && target <= arr[mid] {
		return recursiveSearch(arr, left, mid, target)
	}
	return recursiveSearch(arr, mid+1, right, target)
}
