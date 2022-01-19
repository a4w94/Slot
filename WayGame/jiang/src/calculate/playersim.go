package calculate

import (
	"fmt"
	"math/rand"
	"package/src/info"
	"package/src/public"
	"package/src/table"
	"time"
)

var (
	PlayerAmount int
)

type Player struct {
	Bet            int //玩家下注金額
	BetLevel       int //玩家下注等級
	AccBet         int //玩家總累積下注
	AccScore       int //玩家總累積得分
	InitMoney      int
	RemainingMoney int
}

type PlayerSimResult struct {
	Number_of_People_ScoreDouble int
	Number_of_People_ScoreZero   int
}

type PlayerSimResultRate struct {
	Number_of_People_ScoreDouble_Rate float64
	Number_of_People_ScoreZero_Rate   float64
}

func PlayerSim(playerAmount int, rtp int) PlayerSimResultRate {
	//初始化資訊
	table.Init()
	//修改global
	public.ChangeGlobalRtpModule(rtp)

	PlayerAmount = playerAmount

	rand.Seed(int64(time.Now().UnixNano()))
	table.Init()

	var Result PlayerSimResult
	var ResultRate PlayerSimResultRate

	for i := 0; i < PlayerAmount; i++ {
		var player Player
		player.InitPlayer()

		for {
			if player.RemainingMoney >= player.Bet {
				var each_Round_Result EachRoundResult
				each_Round_Result.EachRound()
				player.CalResult(each_Round_Result)
				//fmt.Println(player)

				if player.RemainingMoney >= player.InitMoney*2 {
					Result.Number_of_People_ScoreDouble++
					break
				}

			} else {
				//fmt.Println(player)
				Result.Number_of_People_ScoreZero++
				break
			}
		}

	}

	ResultRate.CalPlayerSimRate(Result)

	return ResultRate
}

func (result *PlayerSimResultRate) CalPlayerSimRate(input PlayerSimResult) {
	div := func(a, b int) float64 {
		tmp := float64(a) / float64(b)
		return tmp
	}
	totalAmount := input.Number_of_People_ScoreDouble + input.Number_of_People_ScoreZero

	result.Number_of_People_ScoreDouble_Rate = div(input.Number_of_People_ScoreDouble, totalAmount)
	result.Number_of_People_ScoreZero_Rate = div(input.Number_of_People_ScoreZero, totalAmount)

	fmt.Println(result, input)

}

func (result *Player) InitPlayer() {
	result.BetLevel = 1
	result.InitMoney = 2000

	result.Bet = result.BetLevel * info.Bet
	result.AccBet = 0
	result.AccScore = 0
	result.RemainingMoney = result.InitMoney
}

func (result *Player) CalResult(input EachRoundResult) {
	result.RemainingMoney -= result.Bet

	result.AccBet += result.Bet
	result.AccScore += input.MainGame.TotalScore + input.FreeGame.TotalScore
	result.RemainingMoney += input.MainGame.TotalScore + input.FreeGame.TotalScore
}
