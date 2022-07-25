package scoretools

import (
	"package/command"
	"package/src/table"
)

type Line_Game_Combo struct {
	LineGameComboResult [command.Linenum]LineGameEachWay
}
type LineGameEachWay struct {
	ResultLine [command.Panel_Col]int
	Symbol     int
	Combo      int
	Score      int
}

func (result *Line_Game_Combo) CombojudgeLineGame(panel [command.Panel_Row][command.Panel_Col]int) {

	for index, arr := range table.LineTable {
		var resultline [command.Panel_Col]int

		for reel, lineindex := range arr {

			resultline[reel] = panel[lineindex][reel]
		}
		result.LineGameComboResult[index].ResultLine = resultline

		comboquantity := 1

		var symbokind int

		if resultline[0] == command.Wild {
			comboquantity = 0

			for i := 0; i < command.Panel_Col; i++ {
				if resultline[i] == command.Wild {
					symbokind = command.Wild
					continue
				} else if resultline[i] != command.Wild {
					symbokind = resultline[i]

					break
				}

			}
			for i := 0; i < command.Panel_Col; i++ {
				if symbokind == resultline[i] || resultline[i] == command.Wild {
					comboquantity = comboquantity + 1
				} else {
					break
				}
			}
		} else {
			symbokind = resultline[0]

			for i := 1; i < command.Panel_Col; i++ {
				if resultline[0] == resultline[i] || resultline[i] == command.Wild {
					comboquantity = comboquantity + 1

				} else {
					break
				}
			}

		}

		var payresult int

		payresult = table.PayTable[symbokind][comboquantity]

		wildquantity := 1
		var wildpayresult int

		if resultline[0] == command.Wild {
			for i := 1; i < command.Panel_Col; i++ {
				if resultline[0] == resultline[i] {
					wildquantity = wildquantity + 1

				} else {
					break
				}
			}
			wildpayresult = table.PayTable[command.Wild][wildquantity]

		}

		if wildpayresult >= payresult {

			result.LineGameComboResult[index].Symbol = command.Wild
			result.LineGameComboResult[index].Combo = wildquantity

		} else {

			result.LineGameComboResult[index].Symbol = symbokind
			result.LineGameComboResult[index].Combo = comboquantity

		}
	}

}

func (result *Line_Game_Combo) LineGameScore() {
	for index := 0; index < len(result.LineGameComboResult); index++ {
		k := result.LineGameComboResult[index]
		score := table.PayTable[k.Symbol][k.Combo] * command.PlayBetLevel
		result.LineGameComboResult[index].Score = score
	}

}
