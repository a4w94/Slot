package Statistic

import (
	"package/command"
	scoretools "package/src/scoretool"
	"package/src/table"
)

type StatisticTable struct {
	///NG combo統計表初始///
	NGHit   [command.Symbolamount][command.Comboresultnum]int
	NGScore [command.Symbolamount][command.Comboresultnum]int
	///FG combo統計表初始///
	FGHit   [command.Symbolamount][command.Comboresultnum]int
	FGScore [command.Symbolamount][command.Comboresultnum]int

	Scattercombo [command.Comboresultnum]int
}

type StatisticTable_Rate struct {
	///NG combo統計表初始///
	NGHitRate [command.Symbolamount][command.Comboresultnum]float64
	NGRTP     [command.Symbolamount][command.Comboresultnum]float64
	///FG combo統計表初始///
	FGHitRate [command.Symbolamount][command.Comboresultnum]float64
	FGRTP     [command.Symbolamount][command.Comboresultnum]float64

	Scattercombo [command.Comboresultnum]float64
}

func (result *StatisticTable) SymbolHit_WayGame(game_status command.GameStatusName, input scoretools.Way_Game_Combo, inputScatter scoretools.ScatterResult) {
	if game_status == command.MainGame {
		for _, tmp := range input.WayGameComboResult {

			if tmp.Symbol != command.Scatter {
				result.NGHit[tmp.Symbol][tmp.Combo] += tmp.Linequantity
			}
		}
		//scatter 計次
		result.NGHit[command.Scatter][inputScatter.ScatterAmount]++
	} else if game_status == command.FreeGame {
		for _, tmp := range input.WayGameComboResult {
			result.FGHit[tmp.Symbol][tmp.Combo] += tmp.Linequantity
		}
		//scatter 計次
		result.FGHit[command.Scatter][inputScatter.ScatterAmount]++
	}

}

func (result *StatisticTable) SymbolHit_LineGame(game_status command.GameStatusName, input scoretools.Line_Game_Combo, inputScatter scoretools.ScatterResult) {
	if game_status == command.MainGame {
		for _, tmp := range input.LineGameComboResult {

			if tmp.Symbol != command.Scatter {
				result.NGHit[tmp.Symbol][tmp.Combo]++
				result.NGScore[tmp.Symbol][tmp.Combo] += tmp.Score

			}
		}
		//scatter 計次
		result.NGHit[command.Scatter][inputScatter.ScatterAmount]++
	} else if game_status == command.FreeGame {
		for _, tmp := range input.LineGameComboResult {
			result.FGHit[tmp.Symbol][tmp.Combo]++
			result.FGScore[tmp.Symbol][tmp.Combo] += tmp.Score

		}
		//scatter 計次
		result.FGHit[command.Scatter][inputScatter.ScatterAmount]++
	}

}

func (result *StatisticTable) SymbolScore_WayGame(game_status command.GameStatusName, input scoretools.Way_Game_Combo, inputScatter scoretools.ScatterResult) {
	if game_status == command.MainGame {
		for _, tmp := range input.WayGameComboResult {

			if tmp.Symbol != command.Scatter {
				result.NGScore[tmp.Symbol][tmp.Combo] += tmp.Linequantity * table.PayTable[tmp.Symbol][tmp.Combo] * command.PlayBetLevel
			}
		}
		//scatter 計次
		result.NGScore[command.Scatter][inputScatter.ScatterAmount] += table.ScatterTable.NGScatterInfo[inputScatter.ScatterAmount].PayMutiple * command.PlayerBet

	} else if game_status == command.FreeGame {
		for _, tmp := range input.WayGameComboResult {
			result.FGHit[tmp.Symbol][tmp.Combo] += tmp.Linequantity * table.PayTable[tmp.Symbol][tmp.Combo] * command.PlayBetLevel
		}
		//scatter 計次
		result.FGScore[command.Scatter][inputScatter.ScatterAmount] += table.ScatterTable.NGScatterInfo[inputScatter.ScatterAmount].PayMutiple * command.PlayerBet
	}

}
