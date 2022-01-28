package table

import (
	info "package/src/info"
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

func getPayLineTable(xlsxng, xlsxfg *excelize.File) {
	rowlinetable := xlsxng.GetRows("LineTable")

	temp := &Game.LineTable
	for i := 0; i < len(rowlinetable); i++ {
		for k := 1; k < info.Comboresultnum; k++ {
			if rowlinetable[i][k] == "" {
				continue
			} else {
				ele, _ := strconv.Atoi(rowlinetable[i][k])
				temp[i][k-1] = ele
			}

		}
	}

	rowpaytable := xlsxng.GetRows("PayTable")

	temp1 := &Game.PayTable
	for i := 0; i < len(rowpaytable); i++ {
		for k := 1; k < info.Comboresultnum+1; k++ {

			if rowpaytable[i][k] == "" {
				continue
			} else {
				ele, _ := strconv.Atoi(rowpaytable[i][k])
				temp1[i][k-1] = ele
			}

		}
	}

	temp2 := &Game.PayTableSymbol
	for i := 0; i < len(rowpaytable); i++ {
		temp2[i] = rowpaytable[i][0]
	}

	temp3 := &Game.BonusLineTable
	rowbglinetable := xlsxng.GetRows("BonusLineTable")
	for i := 0; i < 80; i++ {
		var arr []int
		for k := 1; k < 6; k++ {
			ele, _ := strconv.Atoi(rowbglinetable[i][k])
			arr = append(arr, ele)
		}
		*temp3 = append(*temp3, arr)
	}

}

func getScatterInfo() {
	var GameScatterInfo = Scatter{
		NGScatterInfo: []ScatterInfo{
			{ScatterAmount: 0, PayMutiple: 0, FGSession: 0},
			{ScatterAmount: 1, PayMutiple: 0, FGSession: 0},
			{ScatterAmount: 2, PayMutiple: 0, FGSession: 0},
			{ScatterAmount: 3, PayMutiple: 1, FGSession: 0},
			{ScatterAmount: 4, PayMutiple: 10, FGSession: 0},
			{ScatterAmount: 5, PayMutiple: 100, FGSession: 0},
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
	tmp := &Game.Scatter
	*tmp = GameScatterInfo
}
