/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2022/3/25 10:50 上午
 * @Desc:
 */

package tarray

import (
	"fmt"
	"testing"
)

//找第K小数
//较难
//查找一个数组的第K小的数，注意同样大小算一样大。
//如  2 1 3 4 5 2 第三小数为3。
//输入描述
//输入有多组数据。
//每组输入n，然后输入n个整数(1<=n<=1000)，再输入k。
//输出描述
//输出第k小的整数。
//示例1输入：
//2 1 3 5 2 2
//3
//输出：
//3
var testArr1 []int = []int{2, 1, 3, 5, 2, 2}

func findKth(a []int, n, K int) int {
	if len(a) < K {
		return -1
	}
	deleteRepeatNum(&a)
	return recursiveSearch(a, 0, len(a), K)
}
func recursiveSearch(a []int, start, stop, K int) int {

}

func deleteRepeatNum(arr *[]int) {
	m := make(map[int]int, len(*arr))
	for _, v := range *arr {
		m[v] = 1
	}
	*arr = (*arr)[0:0]
	for k, _ := range m {
		*arr = append(*arr, k)
	}
	return
}
func TestDeleteRepeatNum(t *testing.T) {
	fmt.Println(testArr1)
	deleteRepeatNum(&testArr1)
	fmt.Println(testArr1)
}

func partition(arr []int, left, right int) (leftMid, rightMid int) {
	if right-left < 1 {
		return left, right
	}
	pivotValue := arr[right]
	leftP := left
	rightP := right - 1
	curP := left
	for curP <= rightP {
		if arr[curP] < pivotValue {
			arr[curP], arr[leftP] = arr[left], arr[leftP]
			leftP++
			curP++
		} else if arr[curP] > pivotValue {
			arr[curP], arr[rightP] = arr[rightP], arr[curP]
			rightP--
		} else {
			curP++
		}
	}
	arr[curP], arr[right] = arr[right], arr[curP]
	return leftP, rightP + 1
}

func QuickSort(arr []int, left, right int) {
	if right-left < 1 {
		return
	}
	leftMid, rightMid := partition(arr, left, right)
	QuickSort(arr, left, leftMid-1)
	QuickSort(arr, rightMid+1, right)
}
func TestQuickSort(t *testing.T) {
	arr := []int{1, 9, 5, 6, 4, 2, 8, 3, 4, 7}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)

}
