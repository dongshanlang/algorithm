/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/9 5:42 下午
 * @Desc:
 */

package tsort

import (
	"fmt"
	"math"
	"testing"
)

func TestBucketSort(t *testing.T) {
	arr := []int{9287, 1, 4, 5, 6, 8, 9, 7, 3, 2, 0, 34, 65, 12, 876, 123, 4567}
	fmt.Println(arr)
	BucketSort(arr)
	fmt.Println(arr)
}

//桶排序
func BucketSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	//1.找到最大值
	maxValue := 0
	for _, v := range arr {
		if v > maxValue {
			maxValue = v
		}
	}
	//2.找到最大值的位数
	var digits = 1
	for maxValue/10 > 0 {
		digits++
		maxValue = maxValue / 10
	}
	var helpArr = make([]int, len(arr))
	for digit := 0; digit < digits; digit++ {
		//3.遍历，进桶
		bucket1 := make([]int, 10)
		for _, v := range arr {
			pos := v / int(math.Pow10(digit)) % 10
			bucket1[pos]++
		}
		//4.生成累加计数桶
		bucket2 := bucket1[:]
		for i := 1; i < len(bucket2); i++ {
			bucket2[i] = bucket2[i] + bucket2[i-1]
		}
		//5.增加临时桶

		for i := len(arr) - 1; i >= 0; i-- {
			pos := arr[i] / int(math.Pow10(digit)) % 10
			helpArr[bucket2[pos]-1] = arr[i]
			bucket2[pos]--
		}
		//6.将helpArr的排序结果拷贝给原始数组，完成一次排序
		for i := 0; i < len(helpArr); i++ {
			arr[i] = helpArr[i]
		}
	}
}
