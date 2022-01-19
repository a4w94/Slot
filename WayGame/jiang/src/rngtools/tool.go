package tools

///盤面工具///
import (
	"math/rand"
	info "package/src/info"
	"package/src/public"
)

///生成盤面
func GameRng(gameStatus string) [info.Col][info.Reelamount]int {

	var rngresult [info.Col][info.Reelamount]int

	if gameStatus == info.GameStatus.MainGame {
		for i := 0; i < info.Reelamount; i++ {

			randnumber := rand.Intn(len(public.Ngstritable[i]))

			for j := 0; j < info.Col; j++ {

				rngresult[j][i] = public.Ngstritable[i][(randnumber+j)%(len(public.Ngstritable[i]))]
			}

		}
	} else if gameStatus == info.GameStatus.FreeGame {
		for i := 0; i < info.Reelamount; i++ {

			randnumber := rand.Intn(len(public.Fgstritable[i]))

			for j := 0; j < info.Col; j++ {

				rngresult[j][i] = public.Fgstritable[i][(randnumber+j)%(len(public.Fgstritable[i]))]
			}

		}
	}

	return rngresult

}

///統計盤面scatter個數///
func CountPanelScatterAmount(panel [info.Col][info.Reelamount]int) int {

	var scatteramount int

	for i := 0; i < info.Col; i++ {

		for k := 0; k < info.Reelamount; k++ {
			if panel[i][k] == info.Scatter {
				scatteramount++

			}
		}

	}

	return scatteramount

}

///統計盤面wild個數///
func CountPanelWildAmount(panel [info.Col][info.Reelamount]int) int {

	var amount int

	for i := 0; i < info.Col; i++ {

		for k := 0; k < info.Reelamount; k++ {
			if panel[i][k] == info.Wild {
				amount++

			}
		}

	}

	return amount

}

//DoubleWay盤面//
func DoubleWayRng(panel [info.Col][info.Reelamount]int) [info.Col][info.Reelamount]int {

	var reversepanel [info.Col][info.Reelamount]int
	for k := 0; k < info.Col; k++ {
		for i := 0; i < info.Reelamount; i++ {
			reversepanel[k][info.Reelamount-1-i] = panel[k][i]
		}
	}

	return reversepanel
}

//中列拓展wild//
func MidWildExpand(panel [info.Col][info.Reelamount]int) [info.Col][info.Reelamount]int {

	for k := 0; k < info.Reelamount; k++ {
		if panel[1][k] == info.Wild {
			panel[0][k] = info.Wild
			panel[2][k] = info.Wild
		}
	}

	return panel

}

//隨機wild//
func RandomWild(panel [info.Col][info.Reelamount]int) [info.Col][info.Reelamount]int {
	//fmt.Println("input", panel)

	randomwildweight := [info.Reelamount][4]int{
		{0, 9, 11, 20},
		{0, 9, 11, 20},
		{0, 9, 11, 20},
		{0, 9, 11, 20},
		{0, 9, 11, 20},
	}

	var wildposition [info.Reelamount]int

	for i := 0; i < info.Reelamount; i++ {
		seed := rand.Intn(randomwildweight[i][3])
		//fmt.Println("seed", seed)

		for k := 1; k < 4; k++ {
			if seed >= randomwildweight[i][k-1] && seed < randomwildweight[i][k] {

				wildposition[i] = k - 1
			}
		}
	}
	//fmt.Println(wildposition)

	for i := 0; i < info.Reelamount; i++ {
		if wildposition[i] == 1 {
			panel[0][i] = info.Wild
			panel[1][i] = info.Wild
			panel[2][i] = info.Wild
		} else {
			panel[wildposition[i]][i] = info.Wild
		}

	}
	//fmt.Println("output", panel)

	return panel

}
