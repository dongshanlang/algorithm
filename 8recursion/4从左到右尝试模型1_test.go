/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2022/1/12 1:56 下午
 * @Desc:
 */

package trecursion

import (
	"algorithm/utils/mqueue"
	"fmt"
	"testing"
)

//规定1和A对应，2和B对应，3和C对应。。。
//那么一个数字字符串比如111就可以转化为：AAA、KA、AK
//给定一个只有数字的字符组成的字符串str，返回有多少种转化结果
//返回所有转化结果
type SpecialNum2Letter struct {
}

func (s *SpecialNum2Letter) process(str, cur string, index int, queue *mqueue.MQueue) {
	if index == len(str) {
		queue.PushBack(cur)
		return
	}
	if str[index] == '0' {
		//这个分支走错了
		return
	} else if str[index] == '1' {
		//可以两位，亦可以一位
		cur1 := cur
		cur1 = fmt.Sprintf("%s%c", cur, 'A')
		s.process(str, cur1, index+1, queue)
		if index+1 < len(str) {
			cur2 := cur
			cur2 = fmt.Sprintf("%s%c", cur2, s.getByteByNumber(str[index:index+2]))
			s.process(str, cur2, index+2, queue)
		}
	} else if str[index] == '2' { //<=26,可以端度也可以结合下一位 ,>26只能当前
		//处理一位
		cur1 := cur
		cur1 = fmt.Sprintf("%s%c", cur, 'B')
		s.process(str, cur1, index+1, queue)

		//处理两位<=26
		if index+1 < len(str) {
			byteChar := s.getByteByNumber(str[index : index+1])
			if byteChar != 0 {
				cur2 := cur
				cur2 = fmt.Sprintf("%s%c", cur2, byteChar)
				s.process(str, cur2, index+2, queue)
			}
		}
	} else if str[index] > '2' { //只能当前位是单独的一个字母
		//处理一位
		cur = fmt.Sprintf("%s%c", cur, s.getByteByNumber(string(str[index])))
		s.process(str, cur, index+1, queue)
	}
}

func TestGetNumber2Letter(t *testing.T) {
	var nl SpecialNum2Letter
	queue := mqueue.New()
	nl.process("1110", "", 0, queue)
	for item := queue.PopFront(); item != nil; item = queue.PopFront() {
		fmt.Println(item.(string))
	}
}

func (s *SpecialNum2Letter) getNumberByByte(b byte) string {
	switch b {
	case 'A':
		return "1"
	case 'B':
		return "2"
	case 'C':
		return "3"
	case 'D':
		return "4"
	case 'E':
		return "5"
	case 'F':
		return "6"
	case 'G':
		return "7"
	case 'H':
		return "8"
	case 'I':
		return "9"
	case 'J':
		return "10"
	case 'K':
		return "11"
	case 'L':
		return "12"
	case 'M':
		return "13"
	case 'N':
		return "14"
	case 'O':
		return "15"
	case 'P':
		return "16"
	case 'Q':
		return "17"
	case 'R':
		return "18"
	case 'S':
		return "19"
	case 'T':
		return "20"
	case 'U':
		return "21"
	case 'V':
		return "22"
	case 'W':
		return "23"
	case 'X':
		return "24"
	case 'Y':
		return "25"
	case 'Z':
		return "26"
	default:
		return ""
	}

}
func (s *SpecialNum2Letter) getByteByNumber(str string) byte {
	switch str {
	case "1":
		return 'A'
	case string('2'):
		return 'B'
	case string('3'):
		return 'C'
	case string('4'):
		return 'D'
	case "5":
		return 'E'
	case "6":
		return 'F'
	case "7":
		return 'G'
	case "8":
		return 'H'
	case "9":
		return 'I'
	case "10":
		return 'J'
	case "11":
		return 'K'
	case "12":
		return 'L'
	case "13":
		return 'M'
	case "14":
		return 'N'
	case "15":
		return 'O'
	case "16":
		return 'P'
	case "17":
		return 'Q'
	case "18":
		return 'R'
	case "19":
		return 'S'
	case "20":
		return 'T'
	case "21":
		return 'U'
	case "22":
		return 'V'
	case "23":
		return 'W'
	case "24":
		return 'X'
	case "25":
		return 'Y'
	case "26":
		return 'Z'
	default:
		return 0
	}
}
