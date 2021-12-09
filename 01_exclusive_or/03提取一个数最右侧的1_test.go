/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/3 6:32 下午
 * @Desc:
 */

package _1_exclusive_or

import "testing"

//提取一个数二进制表示中最右侧的1
func TestGetOne(t *testing.T)  {
	t.Log(getOne(36))
}

//分析
//给出一个数36，二进制表示：0010 0100
//目标是求出最右侧的1，也就是第三位上面的1
//步骤：将其取反得到     1101 1011
//    再将其+1         1101 1100   可以看出最右侧的1是相同的，相同1的左侧因为取反变得不同，相同1的右侧，因为+1全部变成了0
//此时 将原始数和取反+1的数   &运算，可以得到原始数最右侧的1
func getOne(number int)int{
	return number & (^number +1)
}