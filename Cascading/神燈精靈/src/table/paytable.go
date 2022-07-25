package table

import (
	"fmt"
	"package/command"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Scatter struct {
	NGScatterInfo []ScatterInfo
	FGScatterInfo []ScatterInfo
}

type ScatterInfo struct {
	ScatterAmount int
	PayMutiple    int
	FGSession     int
}

func getPayTable(xlsxng, xlsxfg *excelize.File) {

	rowpaytable := xlsxng.GetRows("PayTable")

	for i := 0; i < command.Symbolamount; i++ {
		for k := 1; k < command.Comboresultnum+1; k++ {

			if rowpaytable[i][k] == "" {
				continue
			} else {
				ele, _ := strconv.Atoi(rowpaytable[i][k])
				PayTable[i][k-1] = ele
			}

		}
	}

	for i := 0; i < command.Symbolamount; i++ {
		PayTableSymbol[i] = rowpaytable[i][0]
	}

}

func getScatterInfo() {
	var GameScatterInfo = Scatter{
		NGScatterInfo: []ScatterInfo{
			{ScatterAmount: 0, PayMutiple: 0, FGSession: 0},
			{ScatterAmount: 1, PayMutiple: 0, FGSession: 0},
			{ScatterAmount: 2, PayMutiple: 2, FGSession: 0},
			{ScatterAmount: 3, PayMutiple: 0, FGSession: 10},
			{ScatterAmount: 4, PayMutiple: 0, FGSession: 10},
			{ScatterAmount: 5, PayMutiple: 0, FGSession: 10},
		},
		FGScatterInfo: []ScatterInfo{
			{ScatterAmount: 0, PayMutiple: 0, FGSession: 0},
			{ScatterAmount: 1, PayMutiple: 0, FGSession: 0},
			{ScatterAmount: 2, PayMutiple: 0, FGSession: 0},
			{ScatterAmount: 3, PayMutiple: 0, FGSession: 0},
			{ScatterAmount: 4, PayMutiple: 0, FGSession: 0},
			{ScatterAmount: 5, PayMutiple: 0, FGSession: 0},
		},
	}
	ScatterTable = GameScatterInfo
}

func printPayTable() {
	fmt.Println("PayTable:")
	for i, k := range PayTable {
		if i >= 1 {
			fmt.Println(PayTableSymbol[i], ":", k)
		}
	}
}
