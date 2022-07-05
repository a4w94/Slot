package Statistic

import (
	"package/src/info"
	scoretools "package/src/scoretool"
	"package/src/table"
)

type StatisticTable struct {
	///NG combo統計表初始///
	NGHit   [info.Symbolamount][info.Comboresultnum]int
	NGScore [info.Symbolamount][info.Comboresultnum]int
	///FG combo統計表初始///
	FGHit   [info.Symbolamount][info.Comboresultnum]int
	FGScore [info.Symbolamount][info.Comboresultnum]int

	Scattercombo [info.Comboresultnum]int
}

type StatisticTable_Rate struct {
	///NG combo統計表初始///
	NGHitRate [info.Symbolamount][info.Comboresultnum]float64
	NGRTP     [info.Symbolamount][info.Comboresultnum]float64
	///FG combo統計表初始///
	FGHitRate [info.Symbolamount][info.Comboresultnum]float64
	FGRTP     [info.Symbolamount][info.Comboresultnum]float64

	Scattercombo [info.Comboresultnum]float64
}

func (result *StatisticTable) SymbolHit_WayGame(game_status string, input scoretools.Way_Game_Combo, inputScatter scoretools.ScatterResult) {
	if game_status == info.GameStatus.MainGame {
		for _, tmp := range input.WayGameComboResult {

			if tmp.Symbol != info.Scatter {
				result.NGHit[tmp.Symbol][tmp.Combo] += tmp.Linequantity
			}
		}
		//scatter 計次
		result.NGHit[info.Scatter][inputScatter.ScatterAmount]++
	} else if game_status == info.GameStatus.FreeGame {
		for _, tmp := range input.WayGameComboResult {
			result.FGHit[tmp.Symbol][tmp.Combo] += tmp.Linequantity
		}
		//scatter 計次
		result.FGHit[info.Scatter][inputScatter.ScatterAmount]++
	}

}

func (result *StatisticTable) SymbolHit_LineGame(game_status string, input scoretools.Line_Game_Combo, inputScatter scoretools.ScatterResult) {
	if game_status == info.GameStatus.MainGame {
		for _, tmp := range input.LineGameComboResult {

			if tmp.Symbol != info.Scatter {
				result.NGHit[tmp.Symbol][tmp.Combo]++
				result.NGScore[tmp.Symbol][tmp.Combo] += tmp.Score

			}
		}
		//scatter 計次
		result.NGHit[info.Scatter][inputScatter.ScatterAmount]++
	} else if game_status == info.GameStatus.FreeGame {
		for _, tmp := range input.LineGameComboResult {
			result.FGHit[tmp.Symbol][tmp.Combo]++
			result.FGScore[tmp.Symbol][tmp.Combo] += tmp.Score

		}
		//scatter 計次
		result.FGHit[info.Scatter][inputScatter.ScatterAmount]++
	}

}

func (result *StatisticTable) SymbolScore_WayGame(game_status string, input scoretools.Way_Game_Combo, inputScatter scoretools.ScatterResult) {
	if game_status == info.GameStatus.MainGame {
		for _, tmp := range input.WayGameComboResult {

			if tmp.Symbol != info.Scatter {
				result.NGScore[tmp.Symbol][tmp.Combo] += tmp.Linequantity * table.Game.PayTable[tmp.Symbol][tmp.Combo] * info.PlayBetLevel
			}
		}
		//scatter 計次
		result.NGScore[info.Scatter][inputScatter.ScatterAmount] += table.Game.NGScatterInfo[inputScatter.ScatterAmount].PayMutiple * info.PlayerBet

	} else if game_status == info.GameStatus.FreeGame {
		for _, tmp := range input.WayGameComboResult {
			result.FGHit[tmp.Symbol][tmp.Combo] += tmp.Linequantity * table.Game.PayTable[tmp.Symbol][tmp.Combo] * info.PlayBetLevel
		}
		//scatter 計次
		result.FGScore[info.Scatter][inputScatter.ScatterAmount] += table.Game.NGScatterInfo[inputScatter.ScatterAmount].PayMutiple * info.PlayerBet
	}

}
