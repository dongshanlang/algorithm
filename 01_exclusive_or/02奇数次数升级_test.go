/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/3 6:05 下午
 * @Desc:
 */

package _1_exclusive_or

import (
	"testing"
)

//一个数组中有两种数出现了奇数次，其他数都出现了奇数次，怎么找到并打印这两种数
func TestGetTwoOddNumbers(t *testing.T)  {
	arr := []int{1,2,3,4,5,3,2,1, 4,5,123,234}
	t.Log(getTwoOddOccurNumbers(&arr))
}

//思路：
//1.全部异或一遍找到两个数eor1 = a^b
//2.找到eor1的最右侧的1：rightOne
//3.再次遍历数组，根据arr[i]&rightOne的结果分为两组
//4.两组分别再异或，得到两个数a、b

func getTwoOddOccurNumbers(arr  *[]int)(int,int){
	var (
		eor, num1, rightOne int
	)
	//1.
	for _,v:=range *arr{
		eor^=v
	}
	//2.找到一个数最右侧的1
	rightOne = eor &( ^eor +1)
	//fmt.Println(rightOne)
	//3.
	for _,v:=range *arr{
		//4.
		if v&rightOne != 0{
			num1 ^=v
		}
	}
	return num1, num1^eor
}