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
	Panel               [info.Col][info.Reelamount]int
	TotalScore          int
	ScoreWithoutScatter int

	scoretools.Way_Game_Combo
	scoretools.Way_Game_Score
	scoretools.Line_Game_Combo
	scoretools.Line_Game_Score

	scoretools.ScatterResult
	FreeTriggerStatus bool

	//擴充

}

type FreeGameTotalResult struct {
	TotalSession int
	TotalScore   int

	TotalScoreWithoutScatter int
	ScatterScore             int

	TotalRetriggerTimes int
}

type FreeGameEachRoundResult struct {
	GameStatus          string
	Panel               [info.Col][info.Reelamount]int
	ScoreWithoutScatter int

	scoretools.Way_Game_Combo
	scoretools.Way_Game_Score
	scoretools.Line_Game_Combo
	scoretools.Line_Game_Score

	scoretools.ScatterResult
	ReTriggerStatus bool

	//擴充

}

func (result *MainGameEachRoundResult) MainGame() {
	result.GameStatus = info.GameStatus.MainGame

	//生成盤面
	result.Panel = tools.GameRng(result.GameStatus)

	//scatter 相關
	result.ScatterResult.ScatterAmount = tools.CountPanelScatterAmount(result.Panel)
	result.ScatterResult.ScatterResult(result.GameStatus)
	if result.ScatterResult.ScatterAmount >= 3 {
		result.FreeTriggerStatus = true
	}

	//計算combo
	if info.GameMode == info.GameStatus.WayGame {
		result.Way_Game_Combo.CombojudgeWayGame(result.Panel)
		result.Way_Game_Score.WayGameScore(result.Way_Game_Combo)

	} else if info.GameMode == info.GameStatus.LineGame {
		result.Line_Game_Combo.CombojudgeLineGame(result.Panel)
		result.Line_Game_Score.LineGameScore(result.Line_Game_Combo)

	} else {
		fmt.Println("其他模式")
	}

	//特殊流程

	//計算main game 該次總分
	result.TotalScore = result.Way_Game_Score.ScoreWithoutScatter + result.ScatterResult.Scatterpay

}

func (result *MainGameEachRoundResult) MainGameSpecila() {

}

//FreeGame 流程
func (totalresult *FreeGameTotalResult) FreeGame() {

	for s := 0; s < totalresult.TotalSession; s++ {
		var freeEachRoundResult FreeGameEachRoundResult
		freeEachRoundResult.EachRoundFreeGame()

		//free game retrigger
		if freeEachRoundResult.ReTriggerStatus == true {
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
	}

	//Free Game Total
	totalresult.TotalScore = totalresult.TotalScoreWithoutScatter + totalresult.ScatterScore
}

//每局Free Game
func (result *FreeGameEachRoundResult) EachRoundFreeGame() {
	result.GameStatus = info.GameStatus.FreeGame

	//生成盤面
	result.Panel = tools.GameRng(result.GameStatus)

	//scatter 相關
	result.ScatterResult.ScatterAmount = tools.CountPanelScatterAmount(result.Panel)
	result.ScatterResult.ScatterResult(result.GameStatus)
	if result.ScatterResult.ScatterAmount >= 3 {
		result.ReTriggerStatus = true
	}

	//計算combo
	if info.GameMode == info.GameStatus.WayGame {
		result.Way_Game_Combo.CombojudgeWayGame(result.Panel)
		result.Way_Game_Score.WayGameScore(result.Way_Game_Combo)

	} else if info.GameMode == info.GameStatus.LineGame {
		result.Line_Game_Combo.CombojudgeLineGame(result.Panel)
		result.Line_Game_Score.LineGameScore(result.Line_Game_Combo)

	} else {
		fmt.Println("其他模式")
	}
}

func BonusGame() {

}
