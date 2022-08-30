/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/7 4:26 下午
 * @Desc:
 */

package tsort

import (
	"fmt"
	"testing"
)

//在一个数组中，一个数左边比它小的数的总和，叫数的小和，所有数的小和累加起来，叫数组小和。求数组小和。
//例子：[1,3,4,2,5]
//1左边比1小的数：没有
//3左边比3小的数：1
//4左边比4小的数：1，3
//2左边比1小的数：1
//5左边比5小的数：1，3，4，2
//所以数组的小和为：1+1+3+1+1+3+4+2=16
//要求时间复杂度O(NlogN)，空间复杂度O(N)
//分析merge时右侧大（左侧<右侧），找到一共有几个大，记录下载
func smallSum(arr []int, l, r int) int {
	if l == r {
		return 0
	}
	m := l + (r-l)>>1
	ls := smallSum(arr, l, m)
	rs := smallSum(arr, m+1, r)
	c := mergeSmallSum(arr, l, m, r)
	return ls + rs + c
}
func mergeSmallSum(arr []int, l, m, r int) int {
	if l >= r {
		return 0
	}
	var lStep, rStep, tmpIndex, cal int
	lStep = l
	rStep = m + 1
	length := r - l + 1
	tmpArr := make([]int, length)

	for lStep <= m && rStep <= r {
		if arr[lStep] < arr[rStep] {
			tmpArr[tmpIndex] = arr[lStep]
			cal += arr[lStep] * (r - rStep + 1)
			//fmt.Println("num: ", arr[lStep], ", times: ", r-rStep+1)
			tmpIndex++
			lStep++

		} else {
			tmpArr[tmpIndex] = arr[rStep]
			tmpIndex++
			rStep++
		}
	}
	for lStep <= m {
		tmpArr[tmpIndex] = arr[lStep]
		tmpIndex++
		lStep++
	}
	for rStep <= r {
		tmpArr[tmpIndex] = arr[rStep]
		tmpIndex++
		rStep++
	}
	for i := 0; i < len(tmpArr); i++ {
		arr[l+i] = tmpArr[i]
	}
	return cal
}
func TestSmallSum(t *testing.T) {
	arr := []int{1, 3, 4, 2, 5}
	fmt.Println(smallSum(arr, 0, len(arr)-1))
	fmt.Println(arr)
}
