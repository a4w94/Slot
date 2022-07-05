package calculate

import (
	"fmt"
	"math/rand"
	"package/src/info"
	"package/src/public"
	tools "package/src/rngtools"
	scoretools "package/src/scoretool"
	"package/src/table"
)

var RTP int

type StartBonus interface {
	BonusGame()
}

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
	Bonus
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
	BonusTriggerStatus bool
	Bonus
}

type Bonus struct {
	BonusTotalScore int
	Panel           [][]int
	PayResult       []scoretools.LineGameEachWay
	TotalRound      int
	LockNum         int
}

func (result *MainGameEachRoundResult) MainGame() {
	result.GameStatus = info.GameStatus.MainGame
	//生成盤面
	if !AllComboControl {
		public.ChangeTable(result.GameStatus, RTP)
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

	result.Check_Enter_Bonus()
	//特殊流程

	//計算分數
	if info.GameMode == info.GameStatus.WayGame {
		result.Way_Game_Combo.WayGameScore()

	} else if info.GameMode == info.GameStatus.LineGame {
		result.Line_Game_Combo.LineGameScore()

	} else {
		fmt.Println("其他模式")
	}

	result.LetM1Zero()
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
		if freeEachRoundResult.BonusTriggerStatus {
			StartBonus(&freeEachRoundResult).BonusGame()

			totalresult.TotalScoreWithoutScatter += freeEachRoundResult.BonusTotalScore
		}
	}

	//Free Game Total
	totalresult.TotalScore = totalresult.TotalScoreWithoutScatter + totalresult.ScatterScore
}

//每局Free Game
func (result *FreeGameEachRoundResult) EachRoundFreeGame() {
	result.GameStatus = info.GameStatus.FreeGame
	public.ChangeTable(result.GameStatus, RTP)

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
	result.Check_Enter_Bonus()

	//計算分數
	if info.GameMode == info.GameStatus.WayGame {
		result.Way_Game_Combo.WayGameScore()

	} else if info.GameMode == info.GameStatus.LineGame {
		result.Line_Game_Combo.LineGameScore()

	} else {
		fmt.Println("其他模式")
	}

	result.LetM1Zero()

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

}

func (result *MainGameEachRoundResult) Check_Enter_Bonus() {

	if result.Panel[0][0] == 1 && result.Panel[1][0] == 1 && result.Panel[2][0] == 1 && result.Panel[3][0] == 1 {
		result.BonusTriggerStatus = true
	}

}

func (result *FreeGameEachRoundResult) Check_Enter_Bonus() {

	if result.Panel[0][0] == 1 && result.Panel[1][0] == 1 && result.Panel[2][0] == 1 && result.Panel[3][0] == 1 {
		result.BonusTriggerStatus = true
	}

}

func (result *MainGameEachRoundResult) BonusGame() {
	//fmt.Println("Start Bonus Game")
	//新增的m1 增加一局  新增的scatter 拓展一行

	//觸發後創建盤面
	var panel [][]int
	var growcol_by_enter_scatter int
	for _, m := range result.Panel {
		var arr []int
		for _, k := range m {
			if k == 1 || k == info.Wild {
				arr = append(arr, 1)
			} else if k == info.Scatter {
				arr = append(arr, info.Scatter)
				growcol_by_enter_scatter++
			} else {
				arr = append(arr, info.Space)
			}

		}
		panel = append(panel, arr)
	}
	for i := 0; i < growcol_by_enter_scatter; i++ {
		panel = append(panel, []int{1, info.Space, info.Space, info.Space, info.Space})
	}

	// fmt.Println("Enter Panel")
	// for _, m := range panel {
	// 	fmt.Println(m)
	// }
	// fmt.Println()

	round := 3
	for r := 0; r < round; r++ {
		//fmt.Println("round", r+1, "totalround:", round)
		newpanel := GenerateBonus(result.GameStatus, len(panel))

		var addround bool
		for i, m := range newpanel {
			for j, k := range m {
				if k == 1 && panel[i][j] == info.Space {
					addround = true
				}
				if panel[i][j] == 1 {
					newpanel[i][j] = 1

				}

			}
		}

		var addcol public.RandWeight
		switch RTP {
		case 95:
			addcol.Rand(table.Game.NGWeight95.RespinScatter)
		case 965:
			addcol.Rand(table.Game.NGWeight965.RespinScatter)
		case 99:
			addcol.Rand(table.Game.NGWeight99.RespinScatter)
		}

		if len(newpanel)+addcol.RandIndex <= 8 {
			for i := 0; i < addcol.RandIndex; i++ {
				newpanel = append(newpanel, []int{1, info.Space, info.Space, info.Space, info.Space})

			}
		} else {
			for i := 0; i < 8-len(newpanel); i++ {
				newpanel = append(newpanel, []int{1, info.Space, info.Space, info.Space, info.Space})

			}
		}
		if addround {
			round++
		}
		panel = newpanel

		// fmt.Println("AfterGrowPanel")
		// for _, m := range panel {
		// 	fmt.Println(m)
		// }
		// fmt.Println()
		// fmt.Println()
	}
	result.Bonus.Panel = panel

	result.Bonus.PayResult = scoretools.CounBonusScore(panel)

	for i := 0; i < len(result.Bonus.PayResult); i++ {
		result.Bonus.BonusTotalScore += result.Bonus.PayResult[i].Score
	}

	result.Bonus.TotalRound = round

	for _, m := range result.Bonus.Panel {
		for _, k := range m {
			if k == 1 {
				result.Bonus.LockNum++
			}
		}
	}

}

func (result *FreeGameEachRoundResult) BonusGame() {
	//fmt.Println("Start Bonus Game")
	//新增的m1 增加一局  新增的scatter 拓展一行

	//觸發後創建盤面
	var panel [][]int
	var growcol_by_enter_scatter int
	for _, m := range result.Panel {
		var arr []int
		for _, k := range m {
			if k == 1 || k == info.Wild {
				arr = append(arr, 1)
			} else if k == info.Scatter {
				arr = append(arr, info.Scatter)
				growcol_by_enter_scatter++
			} else {
				arr = append(arr, info.Space)
			}

		}
		panel = append(panel, arr)
	}
	for i := 0; i < growcol_by_enter_scatter; i++ {
		panel = append(panel, []int{1, info.Space, info.Space, info.Space, info.Space})
	}

	// fmt.Println("Enter Panel")
	// for _, m := range panel {
	// 	fmt.Println(m)
	// }
	// fmt.Println()

	round := 3
	for r := 0; r < round; r++ {
		//fmt.Println("round", r+1, "totalround:", round)
		newpanel := GenerateBonus(result.GameStatus, len(panel))

		var addround bool
		for i, m := range newpanel {
			for j, k := range m {
				if k == 1 && panel[i][j] == info.Space {
					addround = true
				}
				if panel[i][j] == 1 {
					newpanel[i][j] = 1

				}

			}
		}

		var addcol public.RandWeight
		switch RTP {
		case 95:
			addcol.Rand(table.Game.FGWeight95.RespinScatter)
		case 965:
			addcol.Rand(table.Game.FGWeight965.RespinScatter)
		case 99:
			addcol.Rand(table.Game.FGWeight99.RespinScatter)
		}

		if len(newpanel)+addcol.RandIndex <= 8 {
			for i := 0; i < addcol.RandIndex; i++ {
				newpanel = append(newpanel, []int{1, info.Space, info.Space, info.Space, info.Space})

			}
		} else {
			for i := 0; i < 8-len(newpanel); i++ {
				newpanel = append(newpanel, []int{1, info.Space, info.Space, info.Space, info.Space})

			}
		}
		if addround {
			round++
		}
		panel = newpanel

		// fmt.Println("AfterGrowPanel")
		// for _, m := range panel {
		// 	fmt.Println(m)
		// }
		// fmt.Println()
		// fmt.Println()
	}
	result.Bonus.Panel = panel

	result.Bonus.PayResult = scoretools.CounBonusScore(panel)

	for i := 0; i < len(result.Bonus.PayResult); i++ {
		result.Bonus.BonusTotalScore += result.Bonus.PayResult[i].Score
	}

	result.Bonus.TotalRound = round

	for _, m := range result.Bonus.Panel {
		for _, k := range m {
			if k == 1 {
				result.Bonus.LockNum++
			}
		}
	}

}

func (result *MainGameEachRoundResult) LetM1Zero() {
	if result.BonusTriggerStatus {
		for i, m := range result.LineGameComboResult {
			if m.Symbol == 1 {
				result.LineGameComboResult[i].Score = 0
			}
		}
	}
}

func (result *FreeGameEachRoundResult) LetM1Zero() {
	if result.BonusTriggerStatus {
		for i, m := range result.LineGameComboResult {
			if m.Symbol == 1 {
				result.LineGameComboResult[i].Score = 0
			}
		}
	}
}

func GenerateBonus(gamestatus string, resultlen int) [][]int {
	//根據盤面長度產盤
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
	var panel [][]int
	switch gamestatus {
	case info.GameStatus.MainGame:
		panel = randpanel(public.NGBonusTable)
		// fmt.Println("產生新盤面")
		// for _, m := range panel {
		// 	fmt.Println(m)
		// }
	case info.GameStatus.FreeGame:
		panel = randpanel(public.FGBonusTable)
		// fmt.Println("產生新盤面")

		// for _, m := range panel {
		// 	fmt.Println(m)
		// }
	}
	return panel
}
