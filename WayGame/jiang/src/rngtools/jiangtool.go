package tools

import (
	"math/rand"
	info "package/src/info"
	"package/src/public"
	"package/src/table"
)

type GrowPanel struct {
	AfterGrowPanel [][info.Reelamount]int
	GrowIndex      int
	Score          int
}

///生成盤面
func GameRng_Jiang(gameStatus string) GrowPanel {
	var result GrowPanel
	var rngresult [info.Col][info.Reelamount]int
	var randNumArr []int

	if gameStatus == info.GameStatus.MainGame {
		for i := 0; i < info.Reelamount; i++ {

			randnumber := rand.Intn(len(public.Ngstritable[i]))
			//randnumber = 2
			randNumArr = append(randNumArr, randnumber)
			for j := 0; j < info.Col; j++ {

				rngresult[j][i] = public.Ngstritable[i][(randnumber+j)%(len(public.Ngstritable[i]))]
			}

		}
	} else if gameStatus == info.GameStatus.FreeGame {
		for i := 0; i < info.Reelamount; i++ {

			randnumber := rand.Intn(len(public.Fgstritable[i]))
			randNumArr = append(randNumArr, randnumber)

			for j := 0; j < info.Col; j++ {

				rngresult[j][i] = public.Fgstritable[i][(randnumber+j)%(len(public.Fgstritable[i]))]
			}

		}
	}

	for i := 0; i < len(rngresult); i++ {
		result.AfterGrowPanel = append(result.AfterGrowPanel, rngresult[i])
	}

	var count2wild int

	for i := 0; i < 2; i++ {

		for j := 0; j < len(rngresult); j++ {

			if rngresult[j][i] == info.Wild {
				count2wild++
			}
		}

	}

	if count2wild == 2 {
		//fmt.Println("2wild")
		var grow_panel table.RandomResult

		if gameStatus == info.GameStatus.MainGame {
			//fmt.Println(table.Game.RTP965.MainGame_Panel_Grow)
			grow_panel.RandResult(table.Game.RTP965.MainGame_Panel_Grow)
			//fmt.Println(grow_panel)
			result.GrowIndex = grow_panel.Index
			for i := 0; i < grow_panel.Index; i++ {
				arr := [5]int{0, 0}
				for j := 2; j < 5; j++ {
					arr[j] = public.Ngstritable[j][(randNumArr[j]+3+i)%(len(public.Ngstritable[j]))]
				}
				result.AfterGrowPanel = append(result.AfterGrowPanel, arr)
			}
		} else if gameStatus == info.GameStatus.FreeGame {
			grow_panel.RandResult(table.Game.RTP965.FreeGame_Panel_Grow)
			result.GrowIndex = grow_panel.Index

			for i := 0; i < grow_panel.Index; i++ {
				arr := [5]int{0, 0}
				for j := 2; j < 5; j++ {
					arr[j] = public.Fgstritable[j][(randNumArr[j]+3+i)%(len(public.Fgstritable[j]))]
				}
				result.AfterGrowPanel = append(result.AfterGrowPanel, arr)
			}
		}

	}

	// fmt.Println(randNumArr)
	// fmt.Println("rngresult")
	// for _, m := range rngresult {
	// 	fmt.Println(m)
	// }

	// fmt.Println("grow")
	// for _, m := range result.AfterGrowPanel {
	// 	fmt.Println(m)
	// }
	return result

}
