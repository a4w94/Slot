package Statistic

import (
	"package/src/info"
	scoretools "package/src/scoretool"
	"package/src/table"
)

type StatisticTable struct {
	///NG combo統計表初始///
	NG [info.Symbolamount][info.Comboresultnum]int
	///FG combo統計表初始///
	FG [info.Symbolamount][info.Comboresultnum]int

	Scattercombo [info.Comboresultnum]int
}

type StatisticTable_Rate struct {
	///NG combo統計表初始///
	NG [info.Symbolamount][info.Comboresultnum]float64
	///FG combo統計表初始///
	FG [info.Symbolamount][info.Comboresultnum]float64

	Scattercombo [info.Comboresultnum]float64
}

func (result *StatisticTable) SymbolHitTimes_WayGame(game_status string, input scoretools.Way_Game_Combo, inputScatter scoretools.ScatterResult) {
	if game_status == info.GameStatus.MainGame {
		for _, tmp := range input.WayGameComboResult {

			if tmp.Symbol != info.Scatter {
				result.NG[tmp.Symbol][tmp.Combo] += tmp.Linequantity
			}
		}
		//scatter 計次
		result.NG[info.Scatter][inputScatter.ScatterAmount]++
	} else if game_status == info.GameStatus.FreeGame {
		for _, tmp := range input.WayGameComboResult {
			result.FG[tmp.Symbol][tmp.Combo] += tmp.Linequantity
		}
		//scatter 計次
		result.FG[info.Scatter][inputScatter.ScatterAmount]++
	}

}

func (result *StatisticTable) SymbolHitTimes_LineGame(game_status string, input scoretools.Line_Game_Combo, inputScatter scoretools.ScatterResult) {
	if game_status == info.GameStatus.MainGame {
		for _, tmp := range input.LineGameComboResult {

			if tmp.Symbol != info.Scatter {
				result.NG[tmp.Symbol][tmp.Combo]++
			}
		}
		//scatter 計次
		result.NG[info.Scatter][inputScatter.ScatterAmount]++
	} else if game_status == info.GameStatus.FreeGame {
		for _, tmp := range input.LineGameComboResult {
			result.FG[tmp.Symbol][tmp.Combo]++
		}
		//scatter 計次
		result.FG[info.Scatter][inputScatter.ScatterAmount]++
	}

}

func (result *StatisticTable) SymbolTotalScore(game_status string, hittimes StatisticTable) {
	if game_status == info.GameStatus.MainGame {

		for i := 1; i < len(result.NG); i++ {
			for j := 0; j < len(result.NG[i]); j++ {
				result.NG[i][j] = hittimes.NG[i][j] * table.Game.PayTable[i][j]
			}
		}

		for i := 1; i < len(result.NG[info.Scatter]); i++ {
			result.NG[info.Scatter][i] = hittimes.NG[info.Scatter][i] * table.Game.Scatter.NGScatterInfo[i].PayMutiple * info.PlayerBet
		}

	} else if game_status == info.GameStatus.FreeGame {

		for i := 1; i < len(result.FG); i++ {
			for j := 0; j < len(result.FG[i]); j++ {
				result.FG[i][j] = hittimes.FG[i][j] * table.Game.PayTable[i][j]
			}
		}
	}

}
