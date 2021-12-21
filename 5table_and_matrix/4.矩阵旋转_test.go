/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/17 5:33 下午
 * @Desc:
 */

package _table_and_matrix

import (
	"testing"
)

//定义一个矩阵
var matrix = [][]int{
	{1, 2, 3, 4},
	{5, 6, 7, 8},
	{9, 10, 11, 12},
	{13, 14, 15, 16},
}

//旋转后
//var matrix = [][]int{
//	{7, 4, 1},
//	{8, 5, 2},
//	{9, 6, 3},
//}
func rotateMatrix(matrix [][]int) [][]int {
	if len(matrix) <= 1 {
		return matrix
	}
	//check if it is n x n matrix
	//定义边界值
	var up, bottom, left, right int
	up = 0
	bottom = len(matrix) - 1
	left = 0
	right = len(matrix[0]) - 1
	//分组(0,0)->(0,max-1)
	//    (0,max)->(max-1, max)
	//    (max,max)->(max, 1)
	//    (max,0)->(1,0)
	//分析
	//up, bottom, left, right
	//
	//up,left   up,right-1
	//up,right      bottom-1, right
	//bottom, right   bottom,left+1
	//bottom,left     up+1,left
	//交换终止条件： left < right
	for left < right {
		for i := 0; left+i < right; i++ {
			tmp := matrix[up][left+i]
			matrix[up][left+i] = matrix[bottom-i][left]
			matrix[bottom-i][left] = matrix[bottom][right-i]
			matrix[bottom][right-i] = matrix[up+i][right]
			matrix[up+i][right] = tmp
		}
		left++
		right--
		bottom--
		up++
	}

	return matrix
}
func TestRotateMatrix(t *testing.T) {
	afterRotate := rotateMatrix(matrix)
	t.Log(afterRotate)
}
