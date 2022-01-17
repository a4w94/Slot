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

type Way_Game_Score struct {
	ScoreWithoutScatter int
}

type Line_Game_Score struct {
	LinePay             [info.Linenum]int
	ScoreWithoutScatter int
}

func (result *Way_Game_Score) WayGameScore(input Way_Game_Combo) {
	for _, k := range input.WayGameComboResult {
		score := table.Game.PayTable[k.Symbol][k.Combo] * k.Linequantity * info.PlayBetLevel
		result.ScoreWithoutScatter += score
	}
}

func (result *Line_Game_Score) LineGameScore(input Line_Game_Combo) {
	for index, k := range input.LineGameComboResult {
		score := table.Game.PayTable[k.Symbol][k.Combo]

		result.LinePay[index] += score
		result.ScoreWithoutScatter += score
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
