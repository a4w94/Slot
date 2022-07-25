package scoretools

import (
	"fmt"
	"package/command"
	"package/src/table"
)

///計算scatter分數／／／

type ScatterResult struct {
	ScatterAmount int
	Scatterpay    int
	Fgsession     int
}

func (scatterinfo *ScatterResult) ScatterResult(gameStatus command.GameStatusName) {
	if gameStatus == command.MainGame {
		for i := 0; i < len(table.ScatterTable.NGScatterInfo); i++ {
			if scatterinfo.ScatterAmount == table.ScatterTable.NGScatterInfo[i].ScatterAmount {
				scatterinfo.Scatterpay = table.ScatterTable.NGScatterInfo[i].PayMutiple * command.PlayerBet
				scatterinfo.Fgsession = table.ScatterTable.NGScatterInfo[i].FGSession

			}

		}
	} else if gameStatus == command.FreeGame {
		for i := 0; i < len(table.ScatterTable.FGScatterInfo); i++ {
			if scatterinfo.ScatterAmount == table.ScatterTable.FGScatterInfo[i].ScatterAmount {
				scatterinfo.Scatterpay = table.ScatterTable.FGScatterInfo[i].PayMutiple * command.PlayerBet
				scatterinfo.Fgsession = table.ScatterTable.FGScatterInfo[i].FGSession

			}

		}
	} else {
		fmt.Println("Enter Wrong GameStatus")
	}

}
