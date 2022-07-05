package table

import (
	"fmt"
	"os"
	info "package/src/info"
	"reflect"
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
	Weight_Module
}

type GameStriTable struct {
	NGStriTablertp95  MegaWay_Table
	NGStriTablertp965 MegaWay_Table
	NGStriTablertp99  MegaWay_Table
	NGStriTablertp92  MegaWay_Table
	NGStriTablertp90  MegaWay_Table

	FGStriTablertp95  MegaWay_Table
	FGStriTablertp965 MegaWay_Table
	FGStriTablertp99  MegaWay_Table
	FGStriTablertp92  MegaWay_Table
	FGStriTablertp90  MegaWay_Table
}

type MegaWay_Table struct {
	R1 [][]int
	R2 [][]int
	R3 [][]int
	R4 [][]int
	R5 [][]int
	R6 [][]int
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
	getWeightTable(xlsxng, xlsxfg)

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

	fmt.Println("Weight:")
	m := Game.Weight_Module.RTP965
	t := reflect.TypeOf(m)
	v := reflect.ValueOf(m)
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name, v.Field(i))
	}

}

func Error(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printTable(game_status string, stritable_name string, input [info.Reelamount][]int) {
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
	fmt.Println("========================================================================================================================================================================")
}
func getexcelparsheet(xlsxng, xlsxfg *excelize.File) {

	var rtproutie = []string{"965", "965", "99", "92", "90"}

	for i := 0; i < len(rtproutie); i++ {
		rowng := xlsxng.GetRows("rtp" + rtproutie[i])
		var stritable MegaWay_Table

		for k := 0; k < info.Reelamount; k++ {
			table_Len := len(rowng[10*k])
			for all_col := 0; all_col < table_Len; all_col++ {
				var tmp []int
				for m := 0; m < info.MegaWay_EachReel_Max; m++ {
					row_index := m
					col_index := all_col
					if rowng[row_index][col_index] != "" {

						sym, _ := strconv.Atoi(rowng[row_index][col_index])
						tmp = append(tmp, sym)
					}
				}
				fmt.Println(tmp)

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
		case "92":
			temp := &Game.NGStriTablertp92
			*temp = stritable
		case "90":
			temp := &Game.NGStriTablertp90
			*temp = stritable
		}

	}

	///FreeGame///

	for i := 0; i < len(rtproutie); i++ {
		rowng := xlsxfg.GetRows("rtp" + rtproutie[i])

		var stritable MegaWay_Table

		for i := 0; i < len(rowng); i++ {

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
		case "92":
			temp := &Game.FGStriTablertp92
			*temp = stritable
		case "90":
			temp := &Game.FGStriTablertp90
			*temp = stritable
		}

	}

}
