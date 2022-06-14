package scoretools

import (
	info "package/src/info"
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

type Line_Game_Combo struct {
	LineGameComboResult [info.Linenum]LineGameEachWay
}
type LineGameEachWay struct {
	ResultLine [info.Reelamount]int
	Symbol     int
	Combo      int
	Score      int
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

func (result *Way_Game_Combo) CombojudgeWayGame(panel [][info.Reelamount]int) {

	var countsymbol [info.Symbolamount][info.Reelamount]int

	for i := 0; i < len(panel); i++ {

		for j := 0; j < info.Reelamount; j++ {
			countsymbol[panel[i][j]][j]++
		}

	}

	for i := 0; i < len(countsymbol); i++ {
		var tmp WayGameEachWay
		tmp.Symbol = i
		if i != info.Wild {
		count:
			for k := 0; k < len(countsymbol[i]); k++ {
				linequantity := countsymbol[i][k] + countsymbol[info.Wild][k]
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
		if tmp.Linequantity != 0 && tmp.Symbol != info.Scatter {
			result.WayGameComboResult = append(result.WayGameComboResult, tmp)
		}
		//fmt.Println("計算waygame", tmp)
	}

}
