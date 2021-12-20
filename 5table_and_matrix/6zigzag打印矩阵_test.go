/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/17 5:36 下午
 * @Desc:
 */

package _table_and_matrix

import (
	"fmt"
	"testing"
)

//请以“之"字型进行遍历输出,zigzag打印矩阵
//分析
//从坐上开始，右上，左下，右上的顺序打印，直到全部打印完成
//比如，4行4列的矩阵
//(0.0)->
//(0,1)->(1,0)->
//(2.0)->(1,1)->(0,2)->
//(0,3)->(1,2)->(2,1)->(3,0)

//就是找寻边界条件
//比如向下的边界条件有三个
//正常向下，col--，row++
//向下未到最底转向：row++
//向下到最低转方向：col++
func zigzagPrintMatrix(matrix *[][]int) {
	row := len(*matrix)
	col := len((*matrix)[0])
	if row < 1 || col < 1 {
		return
	}
	//边界值
	rowMax := row - 1
	colMax := col - 1
	rowFoot := 0
	colFoot := 0
	direction := true
	for rowFoot != rowMax || colFoot != colMax {
		fmt.Println((*matrix)[rowFoot][colFoot])
		if direction {
			if rowFoot > 0 && colFoot < colMax {
				rowFoot--
				colFoot++
			} else if rowFoot == 0 {
				colFoot++
				direction = !direction
			} else if rowFoot != 0 && colFoot == colMax {
				rowFoot++
				direction = !direction
			}

		} else { //向下
			if colFoot > 0 && rowFoot < rowMax { //正常向下
				colFoot--
				rowFoot++
			} else if rowFoot != rowMax && colFoot == 0 { //未到最低需转向
				rowFoot++
				direction = !direction
			} else if rowFoot == rowMax { //到底正常转向
				colFoot++
				direction = !direction
			}
		}
	}
	fmt.Println((*matrix)[rowFoot][colFoot])
}

func TestZigzagMatrix(t *testing.T) {
	var matrix = [][]int{
		{1, 2, 3, 4},
		{6, 7, 8, 9},
		{11, 12, 13, 14},
		{15, 16, 17, 18}}
	//var matrix = [][]int{
	//	{1, 2, 3, 4, 5},
	//	{6, 7, 8, 9, 10},
	//	{11, 12, 13, 14, 15}}
	zigzagPrintMatrix(&matrix)
}
