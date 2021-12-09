/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/3 6:36 下午
 * @Desc:
 */

package _1_exclusive_or

import "testing"

//给定一个数N，请输出N的二进制表示中，出现了多少个1

func TestGetOneCountInNumber(t *testing.T)  {
	t.Log(getOneCountInNumber(15))
}

//分析，这个是查找一个数二进制表示中最右侧1的应用
func getOneCountInNumber(number int)int{
	count :=0
	for ;number!=0;{
		rightOne:=number&(^number+1)
		count++
		number^=rightOne
	}
	return count
}
