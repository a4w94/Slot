package calculate

import (
	"fmt"
	"package/src/info"
	tools "package/src/rngtools"
	scoretools "package/src/scoretool"
)

var RTP int

type MainGameEachRoundResult struct {
	GameStatus          string
	Panel               [][info.Reelamount]int
	TotalScore          int
	ScoreWithoutScatter int

	scoretools.Way_Game_Combo

	scoretools.Line_Game_Combo

	scoretools.ScatterResult
	FreeTriggerStatus bool

	//擴充

	MainGameBonus scoretools.Bonus
	MainGameGrow  tools.GrowPanel
}

type FreeGameTotalResult struct {
	TotalSession int
	TotalScore   int

	TotalScoreWithoutScatter int
	ScatterScore             int

	TotalScoreHitTimes int

	TotalRetriggerTimes int
	TotalFreeGameBonus  scoretools.Bonus
	TotalBonusHit       int
}

type FreeGameEachRoundResult struct {
	GameStatus          string
	Panel               [][info.Reelamount]int
	ScoreWithoutScatter int

	scoretools.Way_Game_Combo

	scoretools.Line_Game_Combo

	scoretools.ScatterResult
	ReTriggerStatus bool

	//擴充
	FreeGameBonus scoretools.Bonus
	FreeGameGrow  tools.GrowPanel
}

func (result *MainGameEachRoundResult) MainGame() {
	result.GameStatus = info.GameStatus.MainGame

	//生成盤面
	if !AllComboControl {
		result.MainGameGrow = tools.GameRng_Jiang(result.GameStatus)
		result.Panel = result.MainGameGrow.AfterGrowPanel
	}

	//scatter 相關
	result.ScatterResult.ScatterAmount = tools.CountPanelScatterAmount(result.Panel)
	result.ScatterResult.ScatterResult(result.GameStatus)
	if result.ScatterResult.ScatterAmount >= 3 {
		result.FreeTriggerStatus = true
	}

	//計算combo
	if info.GameMode == info.GameStatus.WayGame {
		result.Way_Game_Combo.CombojudgeWayGame(result.Panel)

	} else if info.GameMode == info.GameStatus.LineGame {
		//result.Line_Game_Combo.CombojudgeLineGame(result.Panel)

	} else {
		fmt.Println("其他模式")
	}

	//特殊流程
	result.MainGameSpecila()

	//計算分數
	if info.GameMode == info.GameStatus.WayGame {
		result.Way_Game_Combo.WayGameScore()

	} else if info.GameMode == info.GameStatus.LineGame {
		result.Line_Game_Combo.LineGameScore()

	} else {
		fmt.Println("其他模式")
	}
	//計算main game 該次總分

	if info.GameMode == info.GameStatus.WayGame {
		for i := 0; i < len(result.Way_Game_Combo.WayGameComboResult); i++ {
			score := result.Way_Game_Combo.WayGameComboResult[i].Score
			result.TotalScore += score
			result.ScoreWithoutScatter += score
		}

	} else if info.GameMode == info.GameStatus.LineGame {
		for i := 0; i < len(result.Line_Game_Combo.LineGameComboResult); i++ {
			score := result.Line_Game_Combo.LineGameComboResult[i].Score
			result.TotalScore += score
			result.ScoreWithoutScatter += score
		}

	} else {
		fmt.Println("其他模式")
	}

	result.TotalScore += result.ScatterResult.Scatterpay
	//fmt.Println("bonus score", result.MainGameBonus.Score)
	result.TotalScore += result.MainGameBonus.Score

	// if result.TotalScore > 0 {
	// 	fmt.Println(result.WayGameComboResult)
	// }

}

func (result *MainGameEachRoundResult) MainGameSpecila() {
	result.MainGameBonus = scoretools.BounScore(result.GameStatus, result.Panel)

}

//FreeGame 流程
func (totalresult *FreeGameTotalResult) FreeGame() {

	for s := 0; s < totalresult.TotalSession; s++ {
		var freeEachRoundResult FreeGameEachRoundResult
		freeEachRoundResult.EachRoundFreeGame()

		//free game retrigger
		if freeEachRoundResult.ReTriggerStatus {
			//加局
			totalresult.TotalSession += freeEachRoundResult.Fgsession

			//retrigger times
			totalresult.TotalRetriggerTimes++

		}

		//分數累加
		//no scatter score
		totalresult.TotalScoreWithoutScatter += freeEachRoundResult.ScoreWithoutScatter
		// scatter score
		totalresult.ScatterScore += freeEachRoundResult.Scatterpay

		//擴充
		totalresult.TotalFreeGameBonus.Score += freeEachRoundResult.FreeGameBonus.Score
		if freeEachRoundResult.FreeGameBonus.Score > 0 {
			totalresult.TotalBonusHit++
		}

		//得分次數
		eachTotalScore := freeEachRoundResult.ScoreWithoutScatter + freeEachRoundResult.Scatterpay + freeEachRoundResult.FreeGameBonus.Score
		if eachTotalScore > 0 {
			totalresult.TotalScoreHitTimes++
		}

	}

	//Free Game Total
	totalresult.TotalScore = totalresult.TotalScoreWithoutScatter + totalresult.ScatterScore
	totalresult.TotalScore += totalresult.TotalFreeGameBonus.Score
}

//每局Free Game
func (result *FreeGameEachRoundResult) EachRoundFreeGame() {
	result.GameStatus = info.GameStatus.FreeGame

	result.FreeGameGrow = tools.GameRng_Jiang(result.GameStatus)

	//生成盤面
	result.Panel = result.FreeGameGrow.AfterGrowPanel
	//scatter 相關
	result.ScatterResult.ScatterAmount = tools.CountPanelScatterAmount(result.Panel)
	result.ScatterResult.ScatterResult(result.GameStatus)
	if result.ScatterResult.ScatterAmount >= 3 {
		result.ReTriggerStatus = true
	}

	//計算combo
	if info.GameMode == info.GameStatus.WayGame {
		result.Way_Game_Combo.CombojudgeWayGame(result.Panel)

	} else if info.GameMode == info.GameStatus.LineGame {
		//result.Line_Game_Combo.CombojudgeLineGame(result.Panel)

	} else {
		fmt.Println("其他模式")
	}

	//計算分數
	if info.GameMode == info.GameStatus.WayGame {
		result.Way_Game_Combo.WayGameScore()

	} else if info.GameMode == info.GameStatus.LineGame {
		result.Line_Game_Combo.LineGameScore()

	} else {
		fmt.Println("其他模式")
	}
	//計算free game 該次總分

	if info.GameMode == info.GameStatus.WayGame {
		for i := 0; i < len(result.Way_Game_Combo.WayGameComboResult); i++ {
			result.ScoreWithoutScatter += result.Way_Game_Combo.WayGameComboResult[i].Score
		}

	} else if info.GameMode == info.GameStatus.LineGame {
		for i := 0; i < len(result.Line_Game_Combo.LineGameComboResult); i++ {
			result.ScoreWithoutScatter += result.Line_Game_Combo.LineGameComboResult[i].Score
		}

	} else {
		fmt.Println("其他模式")
	}

	//特殊流程
	result.FreeGameSpecila()
}

func (result *FreeGameEachRoundResult) FreeGameSpecila() {
	result.FreeGameBonus = scoretools.BounScore(result.GameStatus, result.Panel)

	multiple := scoretools.FreeGameMultipe().ReturnMultiple
	result.ScoreWithoutScatter *= int(multiple)
	result.Scatterpay *= int(multiple)
	result.FreeGameBonus.Score *= int(multiple)

}
