package scoretools

import (
	"fmt"
	info "package/src/info"
	"package/src/table"
)

///計算scatter分數／／／

type ScatterResult struct {
	ScatterAmount int
	Scatterpay    int
	Fgsession     int
}

func (result *Way_Game_Combo) WayGameScore() {
	for index := 0; index < len(result.WayGameComboResult); index++ {
		k := result.WayGameComboResult[index]
		score := table.Game.PayTable[k.Symbol][k.Combo] * k.Linequantity * info.PlayBetLevel
		result.WayGameComboResult[index].Score = score
	}

}

func (result *Line_Game_Combo) LineGameScore() {
	for index := 0; index < len(result.LineGameComboResult); index++ {
		k := result.LineGameComboResult[index]
		score := table.Game.PayTable[k.Symbol][k.Combo] * info.PlayBetLevel
		result.LineGameComboResult[index].Score = score
	}

}

func (scatterinfo *ScatterResult) ScatterResult(gameStatus string) {
	if gameStatus == info.GameStatus.MainGame {
		for i := 0; i < len(table.Game.NGScatterInfo); i++ {
			if scatterinfo.ScatterAmount == table.Game.Scatter.NGScatterInfo[i].ScatterAmount {
				scatterinfo.Scatterpay = table.Game.Scatter.NGScatterInfo[i].PayMutiple * info.PlayerBet
				scatterinfo.Fgsession = table.Game.Scatter.NGScatterInfo[i].FGSession

			}

		}
	} else if gameStatus == info.GameStatus.FreeGame {
		for i := 0; i < len(table.Game.FGScatterInfo); i++ {
			if scatterinfo.ScatterAmount == table.Game.Scatter.FGScatterInfo[i].ScatterAmount {
				scatterinfo.Scatterpay = table.Game.Scatter.FGScatterInfo[i].PayMutiple * info.PlayerBet
				scatterinfo.Fgsession = table.Game.Scatter.FGScatterInfo[i].FGSession

			}

		}
	} else {
		fmt.Println("Enter Wrong GameStatus")
	}

}

type Bonus struct {
	Amount int
	Score  int
	Combo  int
}
