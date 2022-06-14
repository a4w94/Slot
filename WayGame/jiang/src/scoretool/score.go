package scoretools

import (
	"fmt"
	info "package/src/info"
	"package/src/table"
)

///計算scatter分數／／／

type ScatterResult struct {
	ScatterAmount int
	Scatterpay    int
	Fgsession     int
}

func (result *Way_Game_Combo) WayGameScore() {
	for index := 0; index < len(result.WayGameComboResult); index++ {
		k := result.WayGameComboResult[index]
		score := table.Game.PayTable[k.Symbol][k.Combo] * k.Linequantity * info.PlayBetLevel
		result.WayGameComboResult[index].Score = score
	}

}

func (result *Line_Game_Combo) LineGameScore() {
	for index := 0; index < len(result.LineGameComboResult); index++ {
		k := result.LineGameComboResult[index]
		score := table.Game.PayTable[k.Symbol][k.Combo] * info.PlayBetLevel
		result.LineGameComboResult[index].Score = score
	}

}

func (scatterinfo *ScatterResult) ScatterResult(gameStatus string) {
	if gameStatus == info.GameStatus.MainGame {
		for i := 0; i < len(table.Game.NGScatterInfo); i++ {
			if scatterinfo.ScatterAmount == table.Game.Scatter.NGScatterInfo[i].ScatterAmount {
				scatterinfo.Scatterpay = table.Game.Scatter.NGScatterInfo[i].PayMutiple * info.PlayerBet
				scatterinfo.Fgsession = table.Game.Scatter.NGScatterInfo[i].FGSession

			}

		}
	} else if gameStatus == info.GameStatus.FreeGame {
		for i := 0; i < len(table.Game.FGScatterInfo); i++ {
			if scatterinfo.ScatterAmount == table.Game.Scatter.FGScatterInfo[i].ScatterAmount {
				scatterinfo.Scatterpay = table.Game.Scatter.FGScatterInfo[i].PayMutiple * info.PlayerBet
				scatterinfo.Fgsession = table.Game.Scatter.FGScatterInfo[i].FGSession

			}

		}
	} else {
		fmt.Println("Enter Wrong GameStatus")
	}

}

type Bonus struct {
	Amount int
	Score  int
	Combo  int
}

func BounScore(gameStatus string, panel [][info.Reelamount]int) Bonus {
	var result Bonus

	var count2wild int

	for i := 0; i < 2; i++ {
		for j := 0; j < len(panel); j++ {
			if panel[j][i] == info.Wild {
				count2wild++
			}
		}
	}

	if count2wild == 2 {

		//fmt.Println("2wild")
		var eachReelBonusAmount = [info.Reelamount]int{1, 1, 0, 0, 0}

	countbonus:
		for i := 2; i < info.Reelamount; i++ {
			//fmt.Println("Reel", i)
			if eachReelBonusAmount[i-1] == 0 {
				break countbonus
			}
			var BonusAmount int
			for j := 0; j < len(panel); j++ {
				if panel[j][i] == info.Bonus {
					BonusAmount++
					var random table.RandomResult

					//計算bonus 個數
					result.Amount++

					//計算分數
					if gameStatus == info.GameStatus.MainGame {

						random.RandResult(table.Game.RTP965.MainGame_Bonus)
					} else if gameStatus == info.GameStatus.FreeGame {
						random.RandResult(table.Game.RTP965.FreeGame_Bonus)

					}
					score := random.ReturnMultiple * float64(info.PlayerBet)
					result.Score += int(score)
					// fmt.Println(score, result.Score)
					// fmt.Println("bonus random", random)

				}
			}
			eachReelBonusAmount[i] = BonusAmount
			if BonusAmount > 0 {
				result.Combo = i + 1
			}
			//fmt.Println(BonusAmount)
		}

	}

	//fmt.Println("bonus result", result)
	return result
}

func FreeGameMultipe() table.RandomResult {
	var random table.RandomResult

	random.RandResult(table.Game.RTP965.FreeGameMultiple_1)
	//fmt.Println(table.Game.RTP965.FreeGameMultiple_1)
	if random.ReturnMultiple == -1 {
		//fmt.Println(table.Game.RTP965.FreeGameMultiple_2)
		random.RandResult(table.Game.RTP965.FreeGameMultiple_2)
	}

	return random
}
