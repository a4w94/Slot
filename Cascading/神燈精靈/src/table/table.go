package table

import (
	"fmt"
	"os"
	"package/command"
)

var (
	//cascading

	//megaway
	MegaWayTable GameStriTable_MegaWay
	//waygame
	WayGameStriTable GameStriTable
	//linegame
	LineGameStriTable GameStriTable
	LineTable         [command.Linenum][command.Panel_Col]int

	PayTableSymbol [command.Symbolamount]string
	PayTable       [command.Symbolamount][command.Comboresultnum]int
	ScatterTable   Scatter
)

type GameStriTable struct {
	NGStriTablertp95  [command.Panel_Col][]int
	NGStriTablertp965 [command.Panel_Col][]int
	NGStriTablertp99  [command.Panel_Col][]int
	NGStriTablertp92  [command.Panel_Col][]int
	NGStriTablertp90  [command.Panel_Col][]int

	FGStriTablertp95  [command.Panel_Col][]int
	FGStriTablertp965 [command.Panel_Col][]int
	FGStriTablertp99  [command.Panel_Col][]int
	FGStriTablertp92  [command.Panel_Col][]int
	FGStriTablertp90  [command.Panel_Col][]int
}

var Excelroutieng = "parsheet/ngparsheet.xlsx"
var Excelroutiefg = "parsheet/fgparsheet.xlsx"

func InitTable() {
	switch command.GameMode {
	case command.Cascading:
		Init_Cascading()
	case command.WayGame:
		Init_WayGame()
	case command.LineGame:
		Init_LineGame()
	case command.MegaWay:
		Init_MegaWay()
	}

}

func Error(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printTable(game_status string, stritable_name string, input [command.Panel_Col][]int) {
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
