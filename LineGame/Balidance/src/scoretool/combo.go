package scoretools

import (
	info "package/src/info"
	"package/src/table"
)

type Way_Game_Combo struct {
	WayGameComboResult [info.Col]WayGameEachWay
}
type WayGameEachWay struct {
	Symbol       int
	Combo        int
	Linequantity int
	Score        int
}

type Line_Game_Combo struct {
	LineGameComboResult [info.Linenum]LineGameEachWay
}
type LineGameEachWay struct {
	ResultLine [info.Reelamount]int
	Symbol     int
	Combo      int
	Score      int
	Multiple   int
}

func (result *Line_Game_Combo) CombojudgeLineGame(panel [info.Col][info.Reelamount]int) {

	for index, arr := range table.Game.LineTable {
		var resultline [info.Reelamount]int

		for reel, lineindex := range arr {

			resultline[reel] = panel[lineindex][reel]
		}
		result.LineGameComboResult[index].ResultLine = resultline

		comboquantity := 1

		var symbokind int

		if resultline[0] == info.Wild {
			comboquantity = 0

			for i := 0; i < info.Reelamount; i++ {
				if resultline[i] == info.Wild {
					symbokind = info.Wild
					continue
				} else if resultline[i] != info.Wild {
					symbokind = resultline[i]

					break
				}

			}
			for i := 0; i < info.Reelamount; i++ {
				if symbokind == resultline[i] || resultline[i] == info.Wild {
					comboquantity = comboquantity + 1
				} else {
					break
				}
			}
		} else {
			symbokind = resultline[0]

			for i := 1; i < info.Reelamount; i++ {
				if resultline[0] == resultline[i] || resultline[i] == info.Wild {
					comboquantity = comboquantity + 1

				} else {
					break
				}
			}

		}

		var payresult int

		payresult = table.Game.PayTable[symbokind][comboquantity]

		wildquantity := 1
		var wildpayresult int

		if resultline[0] == info.Wild {
			for i := 1; i < info.Reelamount; i++ {
				if resultline[0] == resultline[i] {
					wildquantity = wildquantity + 1

				} else {
					break
				}
			}
			wildpayresult = table.Game.PayTable[info.Wild][wildquantity]

		}

		if wildpayresult >= payresult {

			result.LineGameComboResult[index].Symbol = info.Wild
			result.LineGameComboResult[index].Combo = wildquantity

		} else {

			result.LineGameComboResult[index].Symbol = symbokind
			result.LineGameComboResult[index].Combo = comboquantity

		}
	}

}

func (result *Way_Game_Combo) CombojudgeWayGame(panel [info.Col][info.Reelamount]int) {

	var symbol int

	for i := 0; i < info.Col; i++ {
		var tmp WayGameEachWay

		symbol = panel[i][0]
		var combo int
		linequantity := 1
		for j := 1; j < info.Reelamount; j++ {
			eachquantity := 0
			for k := 0; k < info.Col; k++ {
				if panel[k][j] == symbol || panel[k][j] == info.Wild {
					eachquantity = eachquantity + 1
				}
			}

			if eachquantity == 0 {
				linequantity = linequantity * 1
			} else {
				linequantity = linequantity * eachquantity
			}

			if eachquantity == 0 {
				combo = j
				break
			} else {
				combo = 5
			}
		}

		tmp.Symbol = symbol
		tmp.Combo = combo
		tmp.Linequantity = linequantity

		result.WayGameComboResult[i] = tmp

	}

}
