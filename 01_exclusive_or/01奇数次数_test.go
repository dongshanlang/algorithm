/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/3 5:52 下午
 * @Desc:
 */

package _1_exclusive_or

import "testing"
//一组数中，某个数出现了奇数次数，其他的数出现的次数为偶数次，请查找出现奇数次数的数是多少？
func TestGetOddNumber(t *testing.T)  {
	arr:=[]int{1,2,3,4,3,2,1}
	number:=GetOddNumber(&arr)
	t.Log(number)
}
func GetOddNumber(arr *[]int)int  {
	number := 0
	for _,v:=range *arr{
		number=number^v
	}
	return number
}
