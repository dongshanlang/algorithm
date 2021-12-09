/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/7 9:19 上午
 * @Desc:
 */

package _2sort

import (
	"fmt"
	"testing"
)

//归并排序从小到大
func TestMergeSort(t *testing.T) {
	arr := []int{1, 4, 9, 6, 2, 5, 7, 3, 6, 4, 2}
	//Process(arr, 0, 8)
	mergeSort(arr)
	fmt.Println(arr)
}

//非递归
func mergeSort(arr []int) {
	var (
		start = 0
		end   = 0
		step  = 1
	)
	for step = 1; step < len(arr); step = step << 1 {
		for start = 0; start < len(arr); start += step * 2 {
			if (start + step*2) > len(arr) {
				end = len(arr) - 1
			} else {
				end = start + step*2 - 1
			}
			mid := start + step - 1
			if mid > len(arr)-1 {
				break
			}
			merge(arr, start, start+step-1, end)
		}
	}
}

//递归
func Process(arr []int, left, right int) {
	if left == right {
		return
	}
	mid := (left + right) / 2
	Process(arr, left, mid)
	Process(arr, mid+1, right)
	merge(arr, left, mid, right)
	//fmt.Println(arr)
}
func merge(arr []int, left, mid, right int) {
	if right == 8 {
		fmt.Println("stop")
	}
	var helpArr []int
	leftPointer := left
	rightPointer := mid + 1
	for leftPointer <= mid && rightPointer <= right {
		if arr[leftPointer] < arr[rightPointer] { //排序的顺序，正序倒叙
			helpArr = append(helpArr, arr[leftPointer])
			leftPointer++
			if leftPointer > mid {
				break
			}
		} else { //相等先取右侧的数
			helpArr = append(helpArr, arr[rightPointer])
			rightPointer++
			if rightPointer > right {
				break
			}
		}
	}
	for leftPointer <= mid {
		helpArr = append(helpArr, arr[leftPointer])
		leftPointer++
	}
	for rightPointer <= right {
		helpArr = append(helpArr, arr[rightPointer])
		rightPointer++
	}
	leftPointer = left
	for _, v := range helpArr {
		arr[leftPointer] = v
		leftPointer++
	}
}
