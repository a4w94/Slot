package tools

import (
	"math/rand"
	"package/command"
	"package/src/public"
)

type LineGame_Rng struct {
	command.GameStatusName //NG or FG

	Panel command.GamePanel //產生盤面

}

func (w *LineGame_Rng) Generate_Rng(gameStatus command.GameStatusName) {
	w.GameStatusName = gameStatus

	var rngresult command.GamePanel

	if gameStatus == command.MainGame {
		for i := 0; i < command.Panel_Col; i++ {

			randnumber := rand.Intn(len(public.Ngstritable[i]))

			for j := 0; j < command.Panel_Row; j++ {

				rngresult[j][i] = public.Ngstritable[i][(randnumber+j)%(len(public.Ngstritable[i]))]
			}

		}
	} else if gameStatus == command.FreeGame {
		for i := 0; i < command.Panel_Col; i++ {

			randnumber := rand.Intn(len(public.Fgstritable[i]))

			for j := 0; j < command.Panel_Row; j++ {

				rngresult[j][i] = public.Fgstritable[i][(randnumber+j)%(len(public.Fgstritable[i]))]
			}

		}
	}

	w.Panel = rngresult

}
