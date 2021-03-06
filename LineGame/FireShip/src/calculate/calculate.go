package calculate

import (
	"fmt"
	"math/rand"
	"package/src/Statistic"
	"package/src/calculate/allcombo"
	"package/src/info"
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
	MainGameRTP_with_scatter    float64
	MainGameRTP_without_scatter float64

	MainGame_ScatterRTP      float64
	MainGame_TriggeFree_Rate float64
	MainGame_ScoreRange_Rate [scoretools.MutipelRange]float64
	MainGame_ScoreRange_RTP  [scoretools.MutipelRange]float64

	//Free Game
	FreeGameRTP_with_scatter float64
	FreeGame_ScatterRTP      float64
	FreeGame_Retrigger_Rate  float64
	FreeGame_Avg_Round       float64
	FreeGame_ScoreRange_Rate [scoretools.MutipelRange]float64
	FreeGame_ScoreRange_RTP  [scoretools.MutipelRange]float64

	//Total
	TotalRTP     float64
	ScoreHitRate float64

	//擴充
	SymbolComboTotalHit Statistic.StatisticTable_Rate

	MainGame_Bonus TotalBonus_Rate
	FreeGame_Bonus TotalBonus_Rate
}

type TotalBonus_Rate struct {
	RTP        float64
	Enter      float64
	LockNumber float64
	Round      float64
}

type TotalRoundResult struct {
	//Total
	TotalBet         int
	TotalScore       int //MainGameScore+FreeGameScore
	EveryScoreRecord []int

	//Main Game
	MainGameTotalScore                        int                          //包含Scatter
	MainGameScore_no_Scatter                  int                          //不包含scatte 分數
	MainGameScatterScore                      int                          //scatter 分數
	MainGameTriggerFreeTotalTimes             int                          //進入free game 次數
	MainGameTotalScore_MultipleRange_Times    [scoretools.MutipelRange]int //倍率區間次數
	MainGameTotalScore_MultipleRange_ScoreAcc [scoretools.MutipelRange]int //倍率區間分數累加

	//Free Game
	FreeGameTotalScore                   int                          //包含Scatter
	FreeGameScore_no_Scatter             int                          //不包含scatte 分數
	FreeGameScatterScore                 int                          //scatter 分數
	FreeGameRetriggeTotalTimes           int                          //retrigger次數
	FreeGameTotal_MultipleRange_Times    [scoretools.MutipelRange]int //倍率區間次數
	FreeGameTotal_MultipleRange_ScoreAcc [scoretools.MutipelRange]int //倍率區間分數累加
	FreeGameTotal_Round                  int                          //Free Game總遊玩局數

	//Bonus Game
	MainGame_Bonus TotalBonus
	FreeGame_Bonus TotalBonus

	//擴充
	SymbolComboTotalHit Statistic.StatisticTable
	AllComboPanel       []allcombo.Panel
}

type TotalBonus struct {
	TotalScore int
	Enter      int
	LockNumber int
	Round      int
}

type EachRoundResult struct {
	Time time.Time
	//Main Game
	MainGame           MainGameEachRoundResult
	MainGameScoreRange int

	//Free Game
	FreeGame                FreeGameTotalResult
	FreeGameTotalScoreRange int
}

func Simulate(session, rtp int) TotalRoundResultRate {

	t := time.Now()
	//初始化資訊
	RTP = rtp
	table.Init()

	public.ChangeGlobalRtpModule(rtp)

	rand.Seed(int64(time.Now().UnixNano()))

	Session = session
	var TotalTimes TotalRoundResult
	var TotalRate TotalRoundResultRate
	TotalTimes.TotalRound()
	TotalRate.TotalRate(TotalTimes)
	TotalRate.PrintResult()
	fmt.Println(time.Since(t))
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
	//-->without scatter rtp
	result.MainGameRTP_without_scatter = div(input.MainGameScore_no_Scatter, input.TotalBet)
	//-->scatter rtp
	result.MainGame_ScatterRTP = div(input.MainGameScatterScore, input.TotalBet)
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
	//-->Retrigger rate
	result.FreeGame_Retrigger_Rate = div(input.FreeGameRetriggeTotalTimes, input.FreeGameTotal_Round)
	//-->分數倍率區間頻率與ＲＴＰ
	for i := 0; i < scoretools.MutipelRange; i++ {
		result.FreeGame_ScoreRange_Rate[i] = div(input.FreeGameTotal_MultipleRange_Times[i], input.MainGameTriggerFreeTotalTimes)
		result.FreeGame_ScoreRange_RTP[i] = div(input.FreeGameTotal_MultipleRange_ScoreAcc[i], input.TotalBet)
	}

	result.FreeGame_Avg_Round = div(input.FreeGameTotal_Round, input.MainGameTriggerFreeTotalTimes)

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

	result.MainGame_Bonus.LockNumber = div(input.MainGame_Bonus.LockNumber, input.MainGame_Bonus.Enter)
	result.MainGame_Bonus.Round = div(input.MainGame_Bonus.Round, input.MainGame_Bonus.Enter)
	result.MainGame_Bonus.RTP = div(input.MainGame_Bonus.TotalScore, input.TotalBet)
	result.MainGame_Bonus.Enter = div(input.MainGame_Bonus.Enter, Session)
}

func (result *TotalRoundResult) TotalRound() {
	alllen := 1
	for _, m := range public.Ngstritable {
		alllen *= len(m)
	}
	// bar := progressbar.Default(int64(alllen))
	if AllComboControl {
		result.AllComboPanel = allcombo.ProductAllPanel()
		Session = len(result.AllComboPanel)
	}

	for i := 0; i < Session; i++ {
		// bar.Add(1)
		var each_Round_Result EachRoundResult

		if AllComboControl {
			each_Round_Result.MainGame.Panel = result.AllComboPanel[i].P
		}
		each_Round_Result.EachRound()

		//累加
		result.TotalBet += info.PlayerBet

		//Main Game
		//-->no scatter score
		result.MainGameScore_no_Scatter += (each_Round_Result.MainGame.ScoreWithoutScatter + each_Round_Result.MainGame.BonusTotalScore)
		//--> scatter score
		result.MainGameScatterScore += each_Round_Result.MainGame.Scatterpay
		//-->total score
		result.MainGameTotalScore += each_Round_Result.MainGame.TotalScore + each_Round_Result.MainGame.BonusTotalScore
		//--> trigger free game times
		if each_Round_Result.MainGame.FreeTriggerStatus {
			result.MainGameTriggerFreeTotalTimes++
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
		//-->total score
		result.FreeGameTotalScore += each_Round_Result.FreeGame.TotalScore
		//--> free game retrigger times
		result.FreeGameRetriggeTotalTimes += each_Round_Result.FreeGame.TotalRetriggerTimes
		//-->score mutiple range 分數倍率區間
		if each_Round_Result.MainGame.FreeTriggerStatus {
			result.FreeGameTotal_MultipleRange_Times[each_Round_Result.FreeGameTotalScoreRange]++
		}
		//-->score range 分數倍率區間ＲＴＰ
		result.FreeGameTotal_MultipleRange_ScoreAcc[each_Round_Result.FreeGameTotalScoreRange] += each_Round_Result.FreeGame.TotalScore
		//-->fg 總遊玩局數加總
		result.FreeGameTotal_Round += each_Round_Result.FreeGame.TotalSession

		//Total
		result.TotalScore += each_Round_Result.MainGame.TotalScore + each_Round_Result.FreeGame.TotalScore + each_Round_Result.MainGame.BonusTotalScore
		result.EveryScoreRecord = append(result.EveryScoreRecord, each_Round_Result.MainGame.TotalScore+each_Round_Result.FreeGame.TotalScore)

		//擴充
		//-->main game 計算各symbol combo 次數
		//-->main game 計算各symbol combo 總分
		if info.GameMode == info.GameStatus.WayGame {
			result.SymbolComboTotalHit.SymbolHit_WayGame(info.GameStatus.MainGame, each_Round_Result.MainGame.Way_Game_Combo, each_Round_Result.MainGame.ScatterResult)
		} else if info.GameMode == info.GameStatus.LineGame {
			result.SymbolComboTotalHit.SymbolHit_LineGame(info.GameStatus.MainGame, each_Round_Result.MainGame.Line_Game_Combo, each_Round_Result.MainGame.ScatterResult)
		} else {
			fmt.Println("Wrong Mode")
		}

		//Bonus
		//NG
		result.MainGame_Bonus.TotalScore += each_Round_Result.MainGame.Bonus.BonusTotalScore
		result.MainGame_Bonus.LockNumber += each_Round_Result.MainGame.LockNum
		result.MainGame_Bonus.Round += each_Round_Result.MainGame.Bonus.TotalRound
		if each_Round_Result.MainGame.BonusTriggerStatus {
			result.MainGame_Bonus.Enter++
		}
	}
	for i, m := range result.SymbolComboTotalHit.NGHit {
		fmt.Println(table.Game.PayTableSymbol[i], m)
	}

	fmt.Println("BonusScore:", result.MainGame_Bonus.TotalScore)

}

func (result *EachRoundResult) EachRound() {

	result.MainGame.MainGame()

	if result.MainGame.BonusTriggerStatus {
		//fmt.Println("Enter Bonus")
		StartBonus(&result.MainGame).BonusGame()
		// fmt.Println("結束盤面")
		// fmt.Println(result.MainGame.Bonus.Panel)

	}

	if result.MainGame.FreeTriggerStatus {
		result.FreeGame.TotalSession = result.MainGame.Fgsession
		result.FreeGame.FreeGame()

	}

	//result.MainGame.PrintEachRoudResult()
	if result.MainGame.Line_Game_Combo.LineGameComboResult[0].Symbol == info.Wild && result.MainGame.Line_Game_Combo.LineGameComboResult[0].Combo == 3 {
		//result.MainGame.PrintEachRoudResult()
	}
	if result.MainGame.TotalScore > 0 {
		//result.MainGame.PrintEachRoudResult()
	}
	//計算
	//Main Game
	//-->總分倍率區間index
	result.MainGameScoreRange = scoretools.Multiplejudge(result.MainGame.TotalScore, info.PlayerBet)
	//Free Game
	//-->總分倍率區間index
	result.FreeGameTotalScoreRange = scoretools.Multiplejudge(result.FreeGame.TotalScore, info.PlayerBet)

	result.Time = time.Now()
}
