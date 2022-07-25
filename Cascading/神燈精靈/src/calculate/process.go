package calculate

import (
	"fmt"
	"package/command"
	tools "package/src/rngtools"
	scoretools "package/src/scoretool"
)

var RTP int

type MainGameEachRoundResult struct {
	GameStatus          command.GameStatusName
	Rng                 tools.Cascading_Rng
	TotalScore          int
	ScoreWithoutScatter int

	Score_and_Combo_Result interface{}

	scoretools.Way_Game_Combo

	scoretools.Line_Game_Combo

	scoretools.ScatterResult
	FreeTriggerStatus bool

	//擴充

}

type FreeGameTotalResult struct {
	TotalSession int
	TotalScore   int

	TotalScoreWithoutScatter int
	ScatterScore             int

	TotalScoreHitTimes int

	TotalRetriggerTimes int
	TotalBonusHit       int
}

type FreeGameEachRoundResult struct {
	GameStatus command.GameStatusName
	Rng        tools.Cascading_Rng

	ScoreWithoutScatter int

	scoretools.Way_Game_Combo

	scoretools.Line_Game_Combo

	scoretools.ScatterResult
	ReTriggerStatus bool

	//擴充

}

func (result *MainGameEachRoundResult) MainGame() {
	result.GameStatus = command.MainGame

	//生成盤面
	if !AllComboControl {
		result.Rng.Generate_Rng(result.GameStatus)

	}

	//scatter 相關
	result.ScatterResult.ScatterAmount = tools.CountPanelScatterAmount(result.Rng.Panel)
	result.ScatterResult.ScatterResult(result.GameStatus)
	if result.ScatterResult.ScatterAmount >= 3 {
		result.FreeTriggerStatus = true
	}

	//計算combo

	if command.GameMode == command.WayGame {
		result.Way_Game_Combo.CombojudgeWayGame(result.Rng.Panel)

	} else if command.GameMode == command.LineGame {
		//result.Line_Game_Combo.CombojudgeLineGame(result.Panel)

	} else {
		fmt.Println("其他模式")
	}

	//特殊流程

	//計算分
	switch command.GameMode {
	case command.WayGame:
		result.Score_and_Combo_Result = scoretools.Way_Game_Combo{}
	case command.LineGame:
		result.Score_and_Combo_Result = scoretools.Line_Game_Combo{}
	case command.Cascading:
		result.Score_and_Combo_Result = scoretools.Cascading_Combo{}

	case command.MegaWay:

	}

	switch v := result.Score_and_Combo_Result.(type) {
	case scoretools.Cascading_Combo:
		//v.ComboJudge()

	}
	if command.GameMode == command.WayGame {
		result.Way_Game_Combo.WayGameScore()

	} else if command.GameMode == command.LineGame {
		result.Line_Game_Combo.LineGameScore()

	} else {
		fmt.Println("其他模式")
	}
	//計算main game 該次總分

	if command.GameMode == command.WayGame {
		for i := 0; i < len(result.Way_Game_Combo.WayGameComboResult); i++ {
			score := result.Way_Game_Combo.WayGameComboResult[i].Score
			result.TotalScore += score
			result.ScoreWithoutScatter += score
		}

	} else if command.GameMode == command.LineGame {
		for i := 0; i < len(result.Line_Game_Combo.LineGameComboResult); i++ {
			score := result.Line_Game_Combo.LineGameComboResult[i].Score
			result.TotalScore += score
			result.ScoreWithoutScatter += score
		}

	} else {
		fmt.Println("其他模式")
	}

	result.TotalScore += result.ScatterResult.Scatterpay

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

		//得分次數
		eachTotalScore := freeEachRoundResult.ScoreWithoutScatter + freeEachRoundResult.Scatterpay
		if eachTotalScore > 0 {
			totalresult.TotalScoreHitTimes++
		}

	}

	//Free Game Total
	totalresult.TotalScore = totalresult.TotalScoreWithoutScatter + totalresult.ScatterScore
}

//每局Free Game
func (result *FreeGameEachRoundResult) EachRoundFreeGame() {
	result.GameStatus = command.FreeGame

	//生成盤面
	result.Rng.Generate_Rng(result.GameStatus)
	//scatter 相關
	result.ScatterResult.ScatterAmount = tools.CountPanelScatterAmount(result.Rng.Panel)
	result.ScatterResult.ScatterResult(result.GameStatus)
	if result.ScatterResult.ScatterAmount >= 3 {
		result.ReTriggerStatus = true
	}

	//計算combo
	if command.GameMode == command.WayGame {
		result.Way_Game_Combo.CombojudgeWayGame(result.Rng.Panel)

	} else if command.GameMode == command.LineGame {
		//result.Line_Game_Combo.CombojudgeLineGame(result.Panel)

	} else {
		fmt.Println("其他模式")
	}

	//計算分數
	if command.GameMode == command.WayGame {
		result.Way_Game_Combo.WayGameScore()

	} else if command.GameMode == command.LineGame {
		result.Line_Game_Combo.LineGameScore()

	} else {
		fmt.Println("其他模式")
	}
	//計算free game 該次總分

	if command.GameMode == command.WayGame {
		for i := 0; i < len(result.Way_Game_Combo.WayGameComboResult); i++ {
			result.ScoreWithoutScatter += result.Way_Game_Combo.WayGameComboResult[i].Score
		}

	} else if command.GameMode == command.LineGame {
		for i := 0; i < len(result.Line_Game_Combo.LineGameComboResult); i++ {
			result.ScoreWithoutScatter += result.Line_Game_Combo.LineGameComboResult[i].Score
		}

	} else {
		fmt.Println("其他模式")
	}

	//特殊流程
}
