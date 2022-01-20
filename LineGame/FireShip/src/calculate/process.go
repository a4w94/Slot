package calculate

import (
	"fmt"
	"math/rand"
	"package/src/info"
	tools "package/src/rngtools"
	scoretools "package/src/scoretool"
	"package/src/table"
)

var RTP int

type MainGameEachRoundResult struct {
	GameStatus          string
	Panel               [info.Col][info.Reelamount]int
	TotalScore          int
	ScoreWithoutScatter int

	scoretools.Way_Game_Combo

	scoretools.Line_Game_Combo

	scoretools.ScatterResult
	FreeTriggerStatus bool

	//擴充
	BonusTriggerStatus bool
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

	scoretools.Line_Game_Combo

	scoretools.ScatterResult
	ReTriggerStatus bool

	//擴充

}

type Bonus struct {
	Panel [][]int
}

func (result *MainGameEachRoundResult) MainGame() {
	result.GameStatus = info.GameStatus.MainGame

	//生成盤面
	if AllComboControl == false {
		result.Panel = tools.GameRng(result.GameStatus)
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
		result.Line_Game_Combo.CombojudgeLineGame(result.Panel)

	} else {
		fmt.Println("其他模式")
	}

	//特殊流程

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

	} else if info.GameMode == info.GameStatus.LineGame {
		result.Line_Game_Combo.CombojudgeLineGame(result.Panel)

	} else {
		fmt.Println("其他模式")
	}

	//特殊流程

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
			result.ScoreWithoutScatter += result.Way_Game_Combo.WayGameComboResult[i].Score
		}

	} else if info.GameMode == info.GameStatus.LineGame {
		for i := 0; i < len(result.Line_Game_Combo.LineGameComboResult); i++ {
			result.ScoreWithoutScatter += result.Line_Game_Combo.LineGameComboResult[i].Score
		}

	} else {
		fmt.Println("其他模式")
	}

}

func (result *MainGameEachRoundResult) Chech_Enter_Bonus() {

	if result.Panel[0][0] == 1 && result.Panel[1][0] == 1 && result.Panel[2][0] == 1 && result.Panel[3][0] == 1 {
		result.BonusTriggerStatus = true
	}

}

func (result *MainGameEachRoundResult) BonusGame() {

	var panel [][]int
	for _, m := range result.Panel {
		var arr []int
		for _, k := range m {
			if k == 1 || k == info.Wild {
				arr = append(arr, 1)
			} else if k == info.Scatter {
				arr = append(arr, info.Scatter)
			} else {
				arr = append(arr, info.Space)
			}

		}
		panel = append(panel, arr)
	}
	fmt.Println("Enter Panel")
	for _, m := range panel {
		fmt.Println(m)
	}
	fmt.Println()
	//for round := 0; round < 3; round++ {

	for i, m := range panel {
		for j, k := range m {
			if k == info.Scatter {
				panel = append(panel, []int{})
				panel[i][j] = info.Space
			}
		}
	}

	fmt.Println("AfterGrowPanel")
	for _, m := range panel {
		fmt.Println(m)
	}
	// }
}

func GenerateBonu(gamestatus string, resultlen int) {
	randpanel := func(table [info.Reelamount][]int) [][]int {
		var panel [][]int
		var index []int
		for i := 0; i < info.Reelamount; i++ {
			randnumber := rand.Intn(len(table[i]))
			index = append(index, randnumber)
		}

		for i := 0; i < resultlen; i++ {
			var arr []int
			for k := 0; k < len(index); k++ {
				symbolindex := (index[k] + i) % len(table[k])
				arr = append(arr, table[k][symbolindex])
			}
			panel = append(panel, arr)
		}
		return panel
	}
	switch gamestatus {
	case info.GameStatus.MainGame:
		result := randpanel(table.Game.NGBonusTable965)
		for _, m := range result {
			fmt.Println(m)
		}

	}
}
