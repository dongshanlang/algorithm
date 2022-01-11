/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/15 4:45 下午
 * @Desc:
 */

package ttree

import (
	"fmt"
	"testing"
)

//题干
//请把一段纸条竖着放在桌子上，然后从纸条的下边向上方对折1次，压出折痕后展开。
//此时折痕是凹下去的，即折痕凸起的方向指向纸条的背面。如果从纸条的下边向上方连续对折2次，
//压出折痕后展开，此时有三条折痕，从上到下依次是下折痕、下折痕、上折痕。
//给定一个输入参数N，代表纸条从下边向上方连续对折N次。请从上到下打印所有折痕的方向。
//例如： N=1时，打印：down；
//N=2时，打印： down down up

func printFoldPaper(currentLevel, allLevel int, down bool) {
	if currentLevel > allLevel {
		return
	}
	printFoldPaper(currentLevel+1, allLevel, true)
	if down {
		fmt.Printf("ao ")
	} else {
		fmt.Printf("tu ")
	}
	printFoldPaper(currentLevel+1, allLevel, false)
}
func TestPrintFoldPaper(t *testing.T) {
	printFoldPaper(1, 3, true)
}
