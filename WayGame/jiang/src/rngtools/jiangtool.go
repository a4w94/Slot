package tools

import (
	"math/rand"
	info "package/src/info"
	"package/src/public"
)

///生成盤面
func GameRng_Jiang(gameStatus string) [info.Col][info.Reelamount]int {

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
