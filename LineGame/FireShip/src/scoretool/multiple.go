package scoretools

const MutipelRange = 9

var (
	Multiple = [MutipelRange]int{0, 0, 1, 5, 10, 30, 50, 100, 200}
)

///計算分數落在哪個區間///
func Multiplejudge(payoff int, playerbet int) int {

	div := payoff / playerbet
	var tmp int
	if payoff == 0 {
		tmp = 0
	} else {
		for i := 1; i < len(Multiple); i++ {
			if div >= Multiple[i-1] && div < Multiple[i] {
				tmp = i - 1
				break

			} else if div >= Multiple[len(Multiple)-1] {
				tmp = len(Multiple) - 1
				break
			}
		}
	}
	//fmt.Println("分數", payoff, "倍數", div, "傳出tmp", tmp)

	return tmp
}
