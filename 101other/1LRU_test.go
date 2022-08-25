/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2022/2/9 4:47 下午
 * @Desc:
 */

package other

import (
	"fmt"
	"testing"
)

func TestCal(t *testing.T) {
	capital := 800000.00
	for ; capital >= 0; capital -= 50000 {
		interestYearRate := 0.052
		interestMonthRate := interestYearRate / 12
		restMonth := 285.0
		interest := capital*interestMonthRate + capital/restMonth
		monthRepaymentCapital := capital / restMonth
		monthRepaymentInterest := capital * interestMonthRate
		fmt.Println("capital: ", capital, "| monthRepaymentCapital: ", monthRepaymentCapital, "| monthRepaymentInterest: ", monthRepaymentInterest, "| all: ", interest, "allWithCommon: ", interest+1600)
	}
}
