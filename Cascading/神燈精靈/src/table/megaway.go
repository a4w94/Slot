package table

import (
	"fmt"
	"package/command"
	"reflect"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type GameStriTable_MegaWay struct {
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

func Init_MegaWay() {
	xlsxng, err := excelize.OpenFile(Excelroutieng)

	Error(err)

	xlsxfg, err := excelize.OpenFile(Excelroutiefg)

	Error(err)

	getPayTable(xlsxng, xlsxfg)
	getScatterInfo()
}

func getMega_Way_parsheet(xlsxng, xlsxfg *excelize.File) {

	var rtproutie = []string{"965"}

	for i := 0; i < len(rtproutie); i++ {
		rowng := xlsxng.GetRows("rtp" + rtproutie[i])
		var stritable MegaWay_Table

		for k := 0; k < command.Panel_Col; k++ {
			excel_index := 10 * k
			table_Len := len(rowng[excel_index])
			fmt.Println("table_len", table_Len)
			for all_col := 0; all_col < table_Len; all_col++ {
				var tmp []int
				for m := 0; m < command.MegaWay_EachReel_Max; m++ {

					row_index := excel_index + m
					col_index := all_col
					if rowng[row_index][col_index] != "" {

						sym, _ := strconv.Atoi(rowng[row_index][col_index])
						tmp = append(tmp, sym)

						//fmt.Println(row_index, col_index, sym)

					}
				}

				if len(tmp) > 0 {
					switch k {
					case 0:
						stritable.R1 = append(stritable.R1, tmp)
					case 1:
						stritable.R2 = append(stritable.R2, tmp)
					case 2:
						stritable.R3 = append(stritable.R3, tmp)
					case 3:
						stritable.R4 = append(stritable.R4, tmp)
					case 4:
						stritable.R5 = append(stritable.R5, tmp)
					case 5:
						stritable.R6 = append(stritable.R6, tmp)
					}

				}
			}

		}
		fmt.Println(stritable)

		switch rtproutie[i] {
		case "95":
			temp := &MegaWayTable.NGStriTablertp95
			*temp = stritable
		case "965":
			temp := &MegaWayTable.NGStriTablertp965
			*temp = stritable
		case "99":
			temp := &MegaWayTable.NGStriTablertp99
			*temp = stritable
		case "92":
			temp := &MegaWayTable.NGStriTablertp92
			*temp = stritable
		case "90":
			temp := &MegaWayTable.NGStriTablertp90
			*temp = stritable
		}

	}

	///FreeGame///

	for i := 0; i < len(rtproutie); i++ {
		rowng := xlsxfg.GetRows("rtp" + rtproutie[i])

		var stritable MegaWay_Table

		for k := 0; k < command.Panel_Col; k++ {
			table_Len := len(rowng[10*k])
			for all_col := 0; all_col < table_Len; all_col++ {
				var tmp []int
				for m := 0; m < command.MegaWay_EachReel_Max; m++ {
					row_index := m
					col_index := all_col
					if rowng[row_index][col_index] != "" {

						sym, _ := strconv.Atoi(rowng[row_index][col_index])
						tmp = append(tmp, sym)
					}
				}

				if len(tmp) > 0 {
					switch k {
					case 0:
						stritable.R1 = append(stritable.R1, tmp)
					case 1:
						stritable.R2 = append(stritable.R2, tmp)
					case 2:
						stritable.R3 = append(stritable.R3, tmp)
					case 3:
						stritable.R4 = append(stritable.R4, tmp)
					case 4:
						stritable.R5 = append(stritable.R5, tmp)
					case 5:
						stritable.R6 = append(stritable.R6, tmp)
					}

				}
			}

		}
		switch rtproutie[i] {
		case "95":
			temp := &MegaWayTable.FGStriTablertp95
			*temp = stritable
		case "965":
			temp := &MegaWayTable.FGStriTablertp965
			*temp = stritable
		case "99":
			temp := &MegaWayTable.FGStriTablertp99
			*temp = stritable
		case "92":
			temp := &MegaWayTable.FGStriTablertp92
			*temp = stritable
		case "90":
			temp := &MegaWayTable.FGStriTablertp90
			*temp = stritable
		}

	}

}

func printMegaWay_Table(game_status string, stritable_name string, table MegaWay_Table) {
	fmt.Println(game_status, stritable_name)
	tp := reflect.TypeOf(table)

	for i := 0; i < tp.NumField(); i++ {
		var t [][]int
		switch i {
		case 0:
			t = table.R1
		case 1:
			t = table.R2
		case 2:
			t = table.R3
		case 3:
			t = table.R4
		case 4:
			t = table.R5
		case 5:
			t = table.R6
		}
		fmt.Println(tp.Field(i).Name)
		for _, j := range t {
			fmt.Println(j)
		}
		fmt.Println()
	}

}
