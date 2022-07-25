package calculate

import (
	"fmt"
	"math/rand"
	"package/command"
	"package/src/Statistic"
	"package/src/calculate/allcombo"
	"package/src/public"
	scoretools "package/src/scoretool"
	"package/src/table"
	"time"
)

var (
	Session         int
	AllComboControl bool = false
)

type TotalRoundResultRate struct {

	//Main Game
	MainGameRTP_with_scatter float64
	MainGame_ScatterRTP      float64
	MainGame_TriggeFree_Rate float64
	MainGame_ScoreRange_Rate [scoretools.MutipelRange]float64
	MainGame_ScoreRange_RTP  [scoretools.MutipelRange]float64

	//Free Game
	FreeGameRTP_with_scatter float64
	FreeGame_ScatterRTP      float64
	FreeGame_Retrigger_Rate  float64
	FreeGame_Score_Hit_Rate  float64
	FreeGame_ScoreRange_Rate [scoretools.MutipelRange]float64
	FreeGame_ScoreRange_RTP  [scoretools.MutipelRange]float64

	//Total
	TotalRTP     float64
	ScoreHitRate float64

	//擴充
	SymbolComboTotalHit Statistic.StatisticTable_Rate

	//數據
	Math_Data
}

type TotalRoundResult struct {
	//Total
	TotalBet         int
	TotalScore       int //MainGameScore+FreeGameScore
	EveryScoreRecord []int

	//Main Game
	MainGameTotalScore            int //包含Scatter
	MainGameScore_no_Scatter      int //不包含scatte 分數
	MainGameScatterScore          int //scatter 分數
	MainGameTriggerFreeTotalTimes int //進入free game 次數
	MainGameTriggerFreeTotalRound int //進入free game 總局數

	MainGameTotalScore_MultipleRange_Times    [scoretools.MutipelRange]int //倍率區間次數
	MainGameTotalScore_MultipleRange_ScoreAcc [scoretools.MutipelRange]int //倍率區間分數累加

	//Free Game
	FreeGameTotalScore                   int                          //包含Scatter
	FreeGameScore_no_Scatter             int                          //不包含scatte 分數
	FreeGameScatterScore                 int                          //scatter 分數
	FreeGameRetriggeTotalTimes           int                          //retrigger次數
	FreeGameScoreHitTimes                int                          //free game 得分次數
	FreeGameTotal_MultipleRange_Times    [scoretools.MutipelRange]int //倍率區間次數
	FreeGameTotal_MultipleRange_ScoreAcc [scoretools.MutipelRange]int //倍率區間分數累加

	//擴充
	SymbolComboTotalHit Statistic.StatisticTable
	AllComboPanel       []command.GamePanel

	//數據
	Math_Data
}

type EachRoundResult struct {
	//Main Game
	MainGame           MainGameEachRoundResult
	MainGameScoreRange int

	//Free Game
	FreeGame                FreeGameTotalResult
	FreeGameTotalScoreRange int
}

func Simulate(session, rtp int) TotalRoundResultRate {
	//初始化資訊
	table.InitTable()

	public.Init(rtp)

	rand.Seed(int64(time.Now().UnixNano()))

	Session = session
	var TotalTimes TotalRoundResult
	var TotalRate TotalRoundResultRate
	TotalTimes.TotalRound()
	TotalRate.TotalRate(TotalTimes)
	TotalRate.PrintResult()
	return TotalRate
}

//計算各數據％
func (result *TotalRoundResultRate) TotalRate(input TotalRoundResult) {
	div := func(a, b int) float64 {
		tmp := float64(a) / float64(b)
		return tmp
	}

	//Main Game
	//-->total rtp
	result.MainGameRTP_with_scatter = div(input.MainGameTotalScore, input.TotalBet)
	//-->scatter rtp
	result.MainGame_ScatterRTP = div(input.MainGameScatterScore, input.TotalBet)
	//--> bonus rtp
	//-->free game hit rate
	result.MainGame_TriggeFree_Rate = div(input.MainGameTriggerFreeTotalTimes, Session)
	//-->分數倍率區間頻率與ＲＴＰ
	for i := 0; i < scoretools.MutipelRange; i++ {
		result.MainGame_ScoreRange_Rate[i] = div(input.MainGameTotalScore_MultipleRange_Times[i], Session)
		result.MainGame_ScoreRange_RTP[i] = div(input.MainGameTotalScore_MultipleRange_ScoreAcc[i], input.TotalBet)
	}

	//Free Game
	//-->total rtp
	result.FreeGameRTP_with_scatter = div(input.FreeGameTotalScore, input.TotalBet)
	//-->scatter rtp
	result.FreeGame_ScatterRTP = div(input.FreeGameScatterScore, input.TotalBet)
	//--> bonus rtp
	//-->Retrigger rate
	result.FreeGame_Retrigger_Rate = div(input.FreeGameRetriggeTotalTimes, input.MainGameTriggerFreeTotalRound)
	//-->score hit rate
	result.FreeGame_Score_Hit_Rate = div(input.FreeGameScoreHitTimes, input.MainGameTriggerFreeTotalRound)
	//-->分數倍率區間頻率與ＲＴＰ
	for i := 0; i < scoretools.MutipelRange; i++ {
		result.FreeGame_ScoreRange_Rate[i] = div(input.FreeGameTotal_MultipleRange_Times[i], input.MainGameTriggerFreeTotalTimes)
		result.FreeGame_ScoreRange_RTP[i] = div(input.FreeGameTotal_MultipleRange_ScoreAcc[i], input.TotalBet)
	}

	//Total
	result.TotalRTP = div(input.TotalScore, input.TotalBet)

	//擴充
	//-->計算Main Game各個symbol combol hit rate
	for i := 0; i < len(result.SymbolComboTotalHit.NGHitRate); i++ {
		for j := 0; j < len(result.SymbolComboTotalHit.NGHitRate[i]); j++ {
			result.SymbolComboTotalHit.NGHitRate[i][j] = div(input.SymbolComboTotalHit.NGHit[i][j], Session)
		}
	}

	//-->計算Main Game各個symbol combol RTP
	for i := 0; i < len(result.SymbolComboTotalHit.NGRTP); i++ {
		for j := 0; j < len(result.SymbolComboTotalHit.NGRTP[i]); j++ {
			result.SymbolComboTotalHit.NGRTP[i][j] = div(input.SymbolComboTotalHit.NGScore[i][j], input.TotalBet)
		}
	}

	//數據計算
	result.Math_Data = input.Math_Data
	result.Math_Data.Calculate_Data()
}

func (result *TotalRoundResult) TotalRound() {
	alllen := 1
	for _, m := range public.Ngstritable {
		alllen *= len(m)
	}
	//bar := progressbar.Default(int64(alllen))
	if AllComboControl {
		result.AllComboPanel = allcombo.ProductAllPanel()
		Session = len(result.AllComboPanel)
	}

	for i := 0; i < Session; i++ {
		//bar.Add(1)
		var each_Round_Result EachRoundResult

		if AllComboControl {
			each_Round_Result.MainGame.Rng.Panel = result.AllComboPanel[i]
		}
		each_Round_Result.EachRound()

		//累加
		result.TotalBet += command.PlayerBet

		//Main Game
		//-->no scatter score
		result.MainGameScore_no_Scatter += each_Round_Result.MainGame.ScoreWithoutScatter
		//--> scatter score
		result.MainGameScatterScore += each_Round_Result.MainGame.Scatterpay
		//--> bonus score
		//-->total score
		result.MainGameTotalScore += each_Round_Result.MainGame.TotalScore
		//--> trigger free game times
		if each_Round_Result.MainGame.FreeTriggerStatus {
			result.MainGameTriggerFreeTotalTimes++
			result.MainGameTriggerFreeTotalRound += each_Round_Result.FreeGame.TotalSession
		}
		//-->score mutiple range 分數倍率區間
		result.MainGameTotalScore_MultipleRange_Times[each_Round_Result.MainGameScoreRange]++
		//-->score range 分數倍率區間ＲＴＰ
		result.MainGameTotalScore_MultipleRange_ScoreAcc[each_Round_Result.MainGameScoreRange] += each_Round_Result.MainGame.TotalScore

		//Free Game
		//-->no scatter score
		result.FreeGameScore_no_Scatter += each_Round_Result.FreeGame.TotalScoreWithoutScatter
		//--> scatter score
		result.FreeGameScatterScore += each_Round_Result.FreeGame.ScatterScore
		//--> bonus score
		//-->total score
		result.FreeGameTotalScore += each_Round_Result.FreeGame.TotalScore
		//--> free game retrigger times
		result.FreeGameRetriggeTotalTimes += each_Round_Result.FreeGame.TotalRetriggerTimes
		//-->free game 得分次數
		result.FreeGameScoreHitTimes += each_Round_Result.FreeGame.TotalScoreHitTimes
		//-->score mutiple range 分數倍率區間
		if each_Round_Result.MainGame.FreeTriggerStatus {
			result.FreeGameTotal_MultipleRange_Times[each_Round_Result.FreeGameTotalScoreRange]++
		}
		//-->score range 分數倍率區間ＲＴＰ
		result.FreeGameTotal_MultipleRange_ScoreAcc[each_Round_Result.FreeGameTotalScoreRange] += each_Round_Result.FreeGame.TotalScore

		//Total
		result.TotalScore += each_Round_Result.MainGame.TotalScore + each_Round_Result.FreeGame.TotalScore
		result.EveryScoreRecord = append(result.EveryScoreRecord, each_Round_Result.MainGame.TotalScore+each_Round_Result.FreeGame.TotalScore)

		//擴充
		//-->main game 計算各symbol combo 次數
		//-->main game 計算各symbol combo 總分
		if command.GameMode == command.WayGame {
			result.SymbolComboTotalHit.SymbolHit_WayGame(command.MainGame, each_Round_Result.MainGame.Way_Game_Combo, each_Round_Result.MainGame.ScatterResult)
			result.SymbolComboTotalHit.SymbolScore_WayGame(command.MainGame, each_Round_Result.MainGame.Way_Game_Combo, each_Round_Result.MainGame.ScatterResult)
		} else if command.GameMode == command.LineGame {
			//result.SymbolComboTotalHit.SymbolHit_LineGame(command.MainGame, each_Round_Result.MainGame.Line_Game_Combo, each_Round_Result.MainGame.ScatterResult)
		} else {
			fmt.Println("Wrong Mode")
		}

		//數據累加
		result.Math_Data.Statistic_All_Score(each_Round_Result)
	}
	for i, m := range result.SymbolComboTotalHit.NGHit {
		fmt.Println(table.PayTableSymbol[i], m)
	}

	fmt.Println("平均free game 局數", float64(result.MainGameTriggerFreeTotalRound)/float64(result.MainGameTriggerFreeTotalTimes))

}

var tmp int

func (result *EachRoundResult) EachRound() {

	result.MainGame.MainGame()

	if result.MainGame.FreeTriggerStatus {
		//fmt.Println(result.MainGame.Fgsession)
		result.FreeGame.TotalSession = result.MainGame.Fgsession
		result.FreeGame.FreeGame()

	}
	fmt.Println(result.MainGame.Rng.Panel)
	// if result.FreeGame.TotalScore > 0 {
	// 	result.FreeGame.PrintEachRoudResult()
	// }

	//計算
	//Main Game
	//-->總分倍率區間index
	result.MainGameScoreRange = scoretools.Multiplejudge(result.MainGame.TotalScore, command.PlayerBet)
	//Free Game
	//-->總分倍率區間index
	result.FreeGameTotalScoreRange = scoretools.Multiplejudge(result.FreeGame.TotalScore, command.PlayerBet)

}
