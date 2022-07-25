package table

import (
	"fmt"
	"package/command"
	"package/weight"
	"reflect"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var CascadingTable GameStriTable_Cascading

type GameStriTable_Cascading struct {
	NGStriTablertp95  Cascading_Weight
	NGStriTablertp965 Cascading_Weight
	NGStriTablertp99  Cascading_Weight
	NGStriTablertp92  Cascading_Weight
	NGStriTablertp90  Cascading_Weight

	FGStriTablertp95  Cascading_Weight
	FGStriTablertp965 Cascading_Weight
	FGStriTablertp99  Cascading_Weight
	FGStriTablertp92  Cascading_Weight
	FGStriTablertp90  Cascading_Weight
}

type Cascading_Weight struct {
	Status                      command.GameStatusName `tag:"狀態"`
	Panel_Weight                weight.Weight          `tag:"盤面權重"`
	Symbol_Weight_for_EachPanel map[int]weight.Weight  `tag:"盤面對應獎圖權重"`
	Scatter                     Scatter_Weight         `tag:"Scatter 權重"`
}
type Scatter_Weight struct {
	//起始盤面scatter權重
	InitWeight weight.Weight
	//消除落下scatter權重
	FallWeight weight.Weight
}

func Init_Cascading() {
	xlsxng, err := excelize.OpenFile(Excelroutieng)

	Error(err)

	xlsxfg, err := excelize.OpenFile(Excelroutiefg)

	Error(err)

	getPayTable(xlsxng, xlsxfg)
	getScatterInfo()

	CascadingTable.NGStriTablertp965.get_Cascading_NG_parsheet("965", xlsxng)

	//get_Cascading_FG_parsheet(xlsxfg)
	//print
	printPayTable()
	printCascadingTable("NG", "965", CascadingTable.NGStriTablertp965)
}

func (c *Cascading_Weight) get_Cascading_NG_parsheet(rtp string, xlsxng *excelize.File) {

	//盤面使用權重表數量
	panel_amount := 4
	symbol_amount := 9
	scatter_Max_amount := 2

	rowng := xlsxng.GetRows("rtp" + rtp)

	c.Status = command.MainGame

	c.Symbol_Weight_for_EachPanel = make(map[int]weight.Weight)

	for panel := 0; panel < panel_amount; panel++ {

		var symbol_weight weight.Weight

		//盤面權重
		index, _ := strconv.Atoi(rowng[1][panel])
		weight, _ := strconv.Atoi(rowng[2][panel])
		acc_weight, _ := strconv.Atoi(rowng[3][panel])

		c.Panel_Weight.Index = append(c.Panel_Weight.Index, index)
		c.Panel_Weight.InitWeight = append(c.Panel_Weight.InitWeight, weight)
		c.Panel_Weight.AccWeight = append(c.Panel_Weight.AccWeight, acc_weight)

		//獎圖權重
		for sym := 0; sym < symbol_amount; sym++ {
			row := 5*(panel+1) + 1

			index, _ := strconv.Atoi(rowng[row][sym])
			weight, _ := strconv.Atoi(rowng[row+1][sym])
			acc_weight, _ := strconv.Atoi(rowng[row+2][sym])

			symbol_weight.Index = append(symbol_weight.Index, index)
			symbol_weight.InitWeight = append(symbol_weight.InitWeight, weight)
			symbol_weight.AccWeight = append(symbol_weight.AccWeight, acc_weight)

		}

		c.Symbol_Weight_for_EachPanel[index] = symbol_weight

	}

	//取得scatter 權重
	for i := 0; i <= scatter_Max_amount; i++ {
		//
		tmp := &c.Scatter.InitWeight

		index, _ := strconv.Atoi(rowng[25][i])
		weight, _ := strconv.Atoi(rowng[26][i])
		acc_weight, _ := strconv.Atoi(rowng[27][i])
		tmp.Index = append(tmp.Index, index)
		tmp.InitWeight = append(tmp.InitWeight, weight)
		tmp.AccWeight = append(tmp.AccWeight, acc_weight)

		//補牌scatter權重
		tmp_fall := &c.Scatter.FallWeight

		index_fall, _ := strconv.Atoi(rowng[29][i])
		weight_fall, _ := strconv.Atoi(rowng[30][i])
		acc_weight_fall, _ := strconv.Atoi(rowng[31][i])
		tmp_fall.Index = append(tmp_fall.Index, index_fall)
		tmp_fall.InitWeight = append(tmp_fall.InitWeight, weight_fall)
		tmp_fall.AccWeight = append(tmp_fall.AccWeight, acc_weight_fall)

	}

}

func printCascadingTable(game_status string, stritable_name string, input Cascading_Weight) {
	fmt.Println(game_status, stritable_name)

	t := reflect.TypeOf(input)
	v := reflect.ValueOf(input)

	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Tag)
		fmt.Printf("%v \n\n", v.Field(i))
	}

}
