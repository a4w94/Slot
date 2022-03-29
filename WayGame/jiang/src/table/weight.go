package table

import (
	"math/rand"
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
				arr_acc = append(arr_acc, arr[i-1]+arr_acc[i-2])
			}
		}
		temp.RTP965.MainGame_Panel_Grow.Multiple = []int{0, 1, 2}
		temp.RTP965.MainGame_Panel_Grow.InitWeight = arr
		temp.RTP965.MainGame_Panel_Grow.AccWeight = arr_acc

	}
	get_nggrow()

}

type RandomResult struct {
	RandSeed     int
	Index        int
	ReturnResult int
}

func (result *RandomResult) RandResult(input Weight) {
	result.RandSeed = rand.Intn(input.AccWeight[len(input.AccWeight)-1])
	if result.RandSeed < input.AccWeight[0] {
		result.Index = 0
		result.ReturnResult = input.Multiple[0]
	} else {
		for i := 0; i < len(input.AccWeight)-1; i++ {

			if input.AccWeight[i] <= result.RandSeed && result.RandSeed < input.AccWeight[i+1] {
				result.Index = i + 1
				result.ReturnResult = input.Multiple[i+1]
				break
			}
		}

	}
}
