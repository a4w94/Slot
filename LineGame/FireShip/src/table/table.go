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
	WeightRTP_Module
	Scatter
	BonusTable
}

type GameStriTable struct {
	NGStriTablertp95  [info.Reelamount][]int
	NGStriTablertp965 [info.Reelamount][]int
	NGStriTablertp99  [info.Reelamount][]int

	FGStriTablertp95  [info.Reelamount][]int
	FGStriTablertp965 [info.Reelamount][]int
	FGStriTablertp99  [info.Reelamount][]int

	NGStriTable2 [info.Reelamount][]int
	FGStriTable2 [info.Reelamount][]int
}

type BonusTable struct {
	NGBonusTable95  [info.Reelamount][]int
	NGBonusTable965 [info.Reelamount][]int
	NGBonusTable99  [info.Reelamount][]int

	FGBonusTable95  [info.Reelamount][]int
	FGBonusTable965 [info.Reelamount][]int
	FGBonusTable99  [info.Reelamount][]int
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
	getweight(xlsxng, xlsxfg)

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

	printWeight := func(game_status string, stritable_name string, input Weight) {
		fmt.Println(game_status, stritable_name)
		fmt.Println("Table Weight")
		fmt.Println(input.Table)
		fmt.Println("Bonus Scatter Weight")
		fmt.Println(input.RespinScatter)
		fmt.Println()

	}
	printTable("NG", "965", Game.NGStriTablertp965)
	printTable("NG", "95", Game.NGStriTablertp95)
	printTable("NG", "99", Game.NGStriTablertp99)
	printTable("FG", "965", Game.FGStriTablertp965)
	printTable("FG", "95", Game.FGStriTablertp95)
	printTable("FG", "99", Game.FGStriTablertp99)
	printTable("NG2", "965", Game.NGStriTable2)
	printTable("FG2", "965", Game.FGStriTable2)

	//Bonus Table
	printTable("NGBG", "95", Game.NGBonusTable95)
	printTable("NGBG", "965", Game.NGBonusTable965)
	printTable("NGBG", "99", Game.NGBonusTable99)
	printTable("FGBG", "95", Game.FGBonusTable95)
	printTable("FGBG", "965", Game.FGBonusTable965)
	printTable("FGBG", "99", Game.FGBonusTable99)

	//Weight
	printWeight("NGWeight", "95", Game.NGWeight95)
	printWeight("NGWeight", "965", Game.NGWeight965)
	printWeight("NGWeight", "99", Game.NGWeight99)
	printWeight("FGWeight", "95", Game.FGWeight95)
	printWeight("FGWeight", "965", Game.FGWeight965)
	printWeight("FGWeight", "99", Game.FGWeight99)

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
			//fmt.Println(rowng[i])

			for k := 0; k < info.Reelamount; k++ {
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

	rowng := xlsxng.GetRows("NGTable2")

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

	temp := &Game.NGStriTable2
	*temp = stritable

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

	rowfg := xlsxfg.GetRows("FGTable2")

	stritablef := [info.Reelamount][]int{}

	for i := 0; i < len(rowfg); i++ {
		for k := 0; k < len(rowfg[i]); k++ {
			if rowfg[i][k] == "" {
				continue
			} else {
				element, err := strconv.Atoi(rowfg[i][k])
				if err != nil {
					panic(err)
				}
				stritablef[k] = append(stritablef[k], element)
			}

		}
	}

	temp2 := &Game.FGStriTable2
	*temp2 = stritablef

	//NG Bonus
	for i := 0; i < len(rtproutie); i++ {
		rowng := xlsxng.GetRows("Bonus" + rtproutie[i])

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
			temp := &Game.BonusTable.NGBonusTable95
			*temp = stritable
		case "965":
			temp := &Game.BonusTable.NGBonusTable965
			*temp = stritable
		case "99":
			temp := &Game.BonusTable.NGBonusTable99
			*temp = stritable
		}

	}

	//FG Bonus
	for i := 0; i < len(rtproutie); i++ {
		rowng := xlsxfg.GetRows("Bonus" + rtproutie[i])

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
			temp := &Game.BonusTable.FGBonusTable95
			*temp = stritable
		case "965":
			temp := &Game.BonusTable.FGBonusTable965
			*temp = stritable
		case "99":
			temp := &Game.BonusTable.FGBonusTable99
			*temp = stritable
		}

	}

}
