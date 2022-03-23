/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2022/3/22 4:14 下午
 * @Desc:
 */

package _0binary_search

import (
	"fmt"
	"testing"
)

//描述
//在一个二维数组array中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
//[
//[1,2,8,9],
//[2,4,9,12],
//[4,7,10,13],
//[6,8,11,15]
//]
//给定 target = 7，返回 true。
//
//给定 target = 3，返回 false。
//
//数据范围：矩阵的长宽满足 0≤n,m≤500 ， 矩阵中的值满足 0≤val≤10^9
//
//进阶：空间复杂度 O(1)O(1) ，时间复杂度 O(n+m)
//示例1
//输入：
//7,[[1,2,8,9],[2,4,9,12],[4,7,10,13],[6,8,11,15]]

//返回值：
//true

//说明：
//存在7，返回true
//示例2
//输入：
//1,[[2]]

//返回值：
//false

//示例3
//输入：
//3,[[1,2,8,9],[2,4,9,12],[4,7,10,13],[6,8,11,15]]

//返回值：
//false

//说明：
//不存在3，返回false

func Find(target int, array [][]int) bool {
	if len(array) == 0 {
		return false
	}
	col := 0
	for col < len(array[0]) {
		if array[0][col] == target {
			return true
		} else if array[0][col] < target {
			if col == len(array[0])-1 {
				break
			}
			col++
		} else {
			break
		}
	}
	row := 1
	for row < len(array) {
		if array[row][col] == target {
			return true
		} else if array[row][col] < target {
			row++
		} else {
			col--
		}
		if col < 0 || row >= len(array) {
			return false
		}
	}
	return false
}
func TestFOr(t *testing.T) {
	arr := [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}}
	fmt.Println(Find(15, arr))
}
