package tools

import (
	"math/rand"
	info "package/src/info"
	"package/src/public"
	"package/src/table"
)

///生成盤面
func GameRng_Jiang(gameStatus string) [][info.Reelamount]int {

	var rngresult [info.Col][info.Reelamount]int
	var randNumArr []int

	if gameStatus == info.GameStatus.MainGame {
		for i := 0; i < info.Reelamount; i++ {

			randnumber := rand.Intn(len(public.Ngstritable[i]))
			randNumArr = append(randNumArr, randnumber)
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

	var rngAfterGrowResult [][info.Reelamount]int

	for i := 0; i < len(rngresult); i++ {
		rngAfterGrowResult = append(rngAfterGrowResult, rngresult[i])
	}

	var count2wild int

	for i := 0; i < 2; i++ {

		for _, n := range rngresult[i] {
			if n == info.Wild {
				count2wild++
			}
		}
	}

	if count2wild == 2 {
		var grow_panel table.RandomResult

		if gameStatus == info.GameStatus.MainGame {
			grow_panel.RandResult(table.Game.RTP965.MainGame_Panel_Grow)
			for i := 0; i < grow_panel.ReturnResult; i++ {
				arr := [5]int{0, 0}
				for j := 2; j < 5; j++ {
					arr[j] = public.Ngstritable[j][(randNumArr[j]+2+i)%(len(public.Ngstritable[j]))]
				}
				rngAfterGrowResult = append(rngAfterGrowResult, arr)
			}
		} else if gameStatus == info.GameStatus.FreeGame {
			grow_panel.RandResult(table.Game.RTP965.FreeGame_Panel_Grow)
			for i := 0; i < grow_panel.ReturnResult; i++ {
				arr := [5]int{0, 0}
				for j := 2; j < 5; j++ {
					arr[j] = public.Fgstritable[j][(randNumArr[j]+2+i)%(len(public.Fgstritable[j]))]
				}
				rngAfterGrowResult = append(rngAfterGrowResult, arr)
			}
		}

	}
	return rngAfterGrowResult

}
