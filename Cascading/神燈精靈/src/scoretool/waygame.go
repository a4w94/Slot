package scoretools

import (
	"package/command"
	"package/src/table"
)

type Way_Game_Combo struct {
	WayGameComboResult []WayGameEachWay
}
type WayGameEachWay struct {
	Symbol       int
	Combo        int
	Linequantity int
	Score        int
}

func (result *Way_Game_Combo) CombojudgeWayGame(panel command.GamePanel) {

	var countsymbol [command.Symbolamount][command.Panel_Col]int

	for i := 0; i < len(panel); i++ {

		for j := 0; j < command.Panel_Col; j++ {
			countsymbol[panel[i][j]][j]++

		}

	}

	for i := 0; i < len(countsymbol); i++ {
		var tmp WayGameEachWay
		tmp.Symbol = i
		if i != command.Wild {
		count:
			for k := 0; k < len(countsymbol[i]); k++ {
				linequantity := countsymbol[i][k] + countsymbol[command.Wild][k]
				if k == 0 {
					if linequantity != 0 {
						tmp.Linequantity = linequantity
						tmp.Combo = k + 1
					}
				} else {
					if linequantity != 0 {
						tmp.Linequantity *= linequantity
						tmp.Combo = k + 1
					} else {
						break count
					}

				}

			}
		} else {
		countwild:
			for k := 0; k < len(countsymbol[i]); k++ {
				linequantity := countsymbol[i][k]
				if k == 0 {
					if linequantity != 0 {
						tmp.Linequantity = linequantity
						tmp.Combo = k + 1
					}
				} else {
					if linequantity != 0 {
						tmp.Linequantity *= linequantity
						tmp.Combo = k + 1
					} else {
						break countwild
					}

				}

			}

		}
		if tmp.Linequantity != 0 && tmp.Symbol != command.Scatter {
			result.WayGameComboResult = append(result.WayGameComboResult, tmp)
		}
		//fmt.Println("計算waygame", tmp)
	}

}

func (result *Way_Game_Combo) WayGameScore() {
	for index := 0; index < len(result.WayGameComboResult); index++ {
		k := result.WayGameComboResult[index]
		score := table.PayTable[k.Symbol][k.Combo] * k.Linequantity * command.PlayBetLevel
		result.WayGameComboResult[index].Score = score
	}

}
