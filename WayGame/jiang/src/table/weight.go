package table

import (
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Weight_Module struct {
	RTP965 WeightInfo
}

type WeightInfo struct {
	MainGame_Panel_Grow Weight
	FreeGame_Panel_Grow Weight
	MainGame_Bonus      Weight
	FreeGame_Bonu       Weight
	FreeGameMultiple_1  Weight
	FreeGameMultiple_2  Weight
}

type Weight struct {
	Multiple   []int
	InitWeight []int
	AccWeight  []int
}

func getWeightTable(xlsxng, xlsxfg *excelize.File) {
	toint := func(input string) int {
		r, _ := strconv.Atoi(input)
		return r
	}
	row := xlsxng.GetRows("Weight")
	temp := &Game.Weight_Module

	//MainGame Panel Grow

	get_nggrow := func() {

		var arr []int
		var arr_acc []int

		for i := 1; i < 4; i++ {
			arr = append(arr, toint(row[1][i]))
			if i == 1 {
				arr_acc = append(arr_acc, arr[i-1])
			} else {
				arr_acc = append(arr_acc, arr[i-1]+arr[i-2])
			}
		}

		temp.RTP965.MainGame_Panel_Grow.InitWeight = arr
		temp.RTP965.MainGame_Panel_Grow.AccWeight = arr_acc

	}
	get_nggrow()

}
