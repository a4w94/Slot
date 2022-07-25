package table

import (
	"fmt"
	"package/command"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func Init_LineGame() {
	xlsxng, err := excelize.OpenFile(Excelroutieng)

	Error(err)

	xlsxfg, err := excelize.OpenFile(Excelroutiefg)

	Error(err)

	getPayTable(xlsxng, xlsxfg)
	getScatterInfo()
	getLineTable(xlsxng, xlsxfg)
	get_LineGame_NG_parsheet(xlsxng, xlsxfg)
	get_LineGame_FG_parsheet(xlsxng, xlsxfg)

	printPayTable()
	printTable("NG", "99", LineGameStriTable.NGStriTablertp99)
	printTable("NG", "965", LineGameStriTable.NGStriTablertp965)
	printTable("NG", "95", LineGameStriTable.NGStriTablertp95)
	printTable("NG", "92", LineGameStriTable.NGStriTablertp92)
	printTable("NG", "90", LineGameStriTable.NGStriTablertp90)
	printTable("FG", "99", LineGameStriTable.FGStriTablertp99)
	printTable("FG", "965", LineGameStriTable.FGStriTablertp965)
	printTable("FG", "95", LineGameStriTable.FGStriTablertp95)
	printTable("FG", "92", LineGameStriTable.FGStriTablertp92)
	printTable("FG", "90", LineGameStriTable.FGStriTablertp90)

}

func getLineTable(xlsxng, xlsxfg *excelize.File) {
	rowlinetable := xlsxng.GetRows("LineTable")

	temp := LineTable
	for i := 0; i < len(rowlinetable); i++ {
		for k := 1; k < command.Comboresultnum; k++ {
			if rowlinetable[i][k] == "" {
				continue
			} else {
				ele, _ := strconv.Atoi(rowlinetable[i][k])
				temp[i][k-1] = ele
			}

		}
	}
}

func printLineTalbe() {
	fmt.Println("LineTable:")
	fmt.Println(LineTable)
	fmt.Println()
}

func get_LineGame_NG_parsheet(xlsxng, xlsxfg *excelize.File) {

	var rtproutie = []string{"95", "965", "99", "92", "90"}

	for i := 0; i < len(rtproutie); i++ {
		rowng := xlsxng.GetRows("rtp" + rtproutie[i])

		stritable := [command.Panel_Col][]int{}

		for i := 0; i < len(rowng); i++ {
			for k := 0; k < len(rowng[i]); k++ {
				if rowng[i][k] == "" {
					continue
				} else {
					element, err := strconv.Atoi(rowng[i][k])
					if err != nil {
						panic(err)
					}
					stritable[k] = append(stritable[k], element)
				}

			}
		}
		switch rtproutie[i] {
		case "95":
			temp := &LineGameStriTable.NGStriTablertp95
			*temp = stritable
		case "965":
			temp := &LineGameStriTable.NGStriTablertp965
			*temp = stritable
		case "99":
			temp := &LineGameStriTable.NGStriTablertp99
			*temp = stritable
		case "92":
			temp := &LineGameStriTable.NGStriTablertp92
			*temp = stritable
		case "90":
			temp := &LineGameStriTable.NGStriTablertp90
			*temp = stritable
		}

	}

}

func get_LineGame_FG_parsheet(xlsxng, xlsxfg *excelize.File) {

	var rtproutie = []string{"95", "965", "99", "92", "90"}

	///FreeGame///

	for i := 0; i < len(rtproutie); i++ {
		rowng := xlsxfg.GetRows("rtp" + rtproutie[i])

		stritable := [command.Panel_Col][]int{}

		for i := 0; i < len(rowng); i++ {
			for k := 0; k < len(rowng[i]); k++ {
				if rowng[i][k] == "" {
					continue
				} else {
					element, err := strconv.Atoi(rowng[i][k])
					if err != nil {
						panic(err)
					}
					stritable[k] = append(stritable[k], element)
				}

			}
		}
		switch rtproutie[i] {
		case "95":
			temp := &LineGameStriTable.FGStriTablertp95
			*temp = stritable
		case "965":
			temp := &LineGameStriTable.FGStriTablertp965
			*temp = stritable
		case "99":
			temp := &LineGameStriTable.FGStriTablertp99
			*temp = stritable
		case "92":
			temp := &LineGameStriTable.FGStriTablertp92
			*temp = stritable
		case "90":
			temp := &LineGameStriTable.FGStriTablertp90
			*temp = stritable
		}

	}

}
