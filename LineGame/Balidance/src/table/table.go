package table

import (
	"fmt"
	"os"
	info "package/src/info"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var Game = GameTabel{}

type GameTabel struct {
	GameStriTable
	PayTableSymbol [info.Symbolamount]string
	PayTable       [info.Symbolamount][info.Comboresultnum]int
	LineTable      [info.Linenum][info.Reelamount]int
	Scatter
}

type GameStriTable struct {
	NGStriTablertp95  [info.Reelamount][]int
	NGStriTablertp965 [info.Reelamount][]int
	NGStriTablertp99  [info.Reelamount][]int

	FGStriTablertp95  [info.Reelamount][]int
	FGStriTablertp965 [info.Reelamount][]int
	FGStriTablertp99  [info.Reelamount][]int
}

var Excelroutieng = "parsheet/ngparsheet.xlsx"
var Excelroutiefg = "parsheet/fgparsheet.xlsx"

func Init() {
	xlsxng, err := excelize.OpenFile(Excelroutieng)

	Error(err)

	xlsxfg, err := excelize.OpenFile(Excelroutiefg)

	Error(err)

	getexcelparsheet(xlsxng, xlsxfg)
	getPayLineTable(xlsxng, xlsxfg)
	getScatterInfo()

	fmt.Println("LineTable:")
	fmt.Println(Game.LineTable)
	fmt.Println()

	fmt.Println("PayTable:")
	for i, k := range Game.PayTable {
		if i >= 1 {
			fmt.Println(Game.PayTableSymbol[i], ":", k)
		}
	}

	fmt.Println()

	printTable := func(game_status string, stritable_name string, input [info.Reelamount][]int) {
		fmt.Println(game_status, stritable_name)
		for _, j := range input {
			for index, name := range j {
				if index == len(j)-1 {
					fmt.Print(name)
				} else {
					fmt.Print(name, ",")
				}
			}
			fmt.Println()
		}
	}
	printTable("NG", "965", Game.NGStriTablertp965)
	printTable("NG", "95", Game.NGStriTablertp95)
	printTable("NG", "99", Game.NGStriTablertp99)
	printTable("FG", "965", Game.FGStriTablertp965)
	printTable("FG", "95", Game.FGStriTablertp95)
	printTable("FG", "99", Game.FGStriTablertp99)

}

func Error(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getexcelparsheet(xlsxng, xlsxfg *excelize.File) {

	var rtproutie = []string{"95", "965", "99"}

	for i := 0; i < len(rtproutie); i++ {
		rowng := xlsxng.GetRows("rtp" + rtproutie[i])

		stritable := [info.Reelamount][]int{}

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
			temp := &Game.NGStriTablertp95
			*temp = stritable
		case "965":
			temp := &Game.NGStriTablertp965
			*temp = stritable
		case "99":
			temp := &Game.NGStriTablertp99
			*temp = stritable
		}

	}

	///FreeGame///

	for i := 0; i < len(rtproutie); i++ {
		rowng := xlsxfg.GetRows("rtp" + rtproutie[i])

		stritable := [info.Reelamount][]int{}

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
			temp := &Game.FGStriTablertp95
			*temp = stritable
		case "965":
			temp := &Game.FGStriTablertp965
			*temp = stritable
		case "99":
			temp := &Game.FGStriTablertp99
			*temp = stritable
		}

	}

}
