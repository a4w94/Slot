package table

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type WeightRTP_Module struct {
	NGWeight95  Weight
	NGWeight965 Weight
	NGWeight99  Weight

	FGWeight95  Weight
	FGWeight965 Weight
	FGWeight99  Weight
}

type Weight struct {
	RespinScatter WeightArr
	Table         WeightArr
}

type WeightArr struct {
	Init_Arr []int
	Acc_Arr  []int
}

func getweight(xlsxng, xlsxfg *excelize.File) {
	var rtproutie = []string{"95", "965", "99"}

	//NG Weight
	for i := 0; i < len(rtproutie); i++ {
		rowng := xlsxng.GetRows("Weight" + rtproutie[i])
		fmt.Println(rowng)
		var tmpweight Weight
		get_table_weight := func() {

			t1, _ := strconv.Atoi(rowng[1][1])
			t2, _ := strconv.Atoi(rowng[1][2])
			tmpweight.Table.Init_Arr = append(tmpweight.Table.Init_Arr, t1, t2)
			tmpweight.Table.Acc_Arr = append(tmpweight.Table.Acc_Arr, t1, t1+t2)

		}

		get_bonus_scatter_weight := func() {
			var total int
			for i := 1; i < len(rowng[4]); i++ {
				t, _ := strconv.Atoi(rowng[4][i])
				total += t
				tmpweight.RespinScatter.Init_Arr = append(tmpweight.RespinScatter.Init_Arr, t)
				tmpweight.RespinScatter.Acc_Arr = append(tmpweight.RespinScatter.Acc_Arr, total)

			}
		}

		get_table_weight()
		get_bonus_scatter_weight()

		switch rtproutie[i] {
		case "95":
			temp := &Game.WeightRTP_Module.NGWeight95
			*temp = tmpweight
		case "965":
			temp := &Game.WeightRTP_Module.NGWeight965
			*temp = tmpweight
		case "99":
			temp := &Game.WeightRTP_Module.NGWeight99
			*temp = tmpweight
		}
	}

	//FG Weight
	for i := 0; i < len(rtproutie); i++ {
		rowng := xlsxfg.GetRows("Weight" + rtproutie[i])
		fmt.Println(rowng)
		var tmpweight Weight
		get_table_weight := func() {

			t1, _ := strconv.Atoi(rowng[1][1])
			t2, _ := strconv.Atoi(rowng[1][2])
			tmpweight.Table.Init_Arr = append(tmpweight.Table.Init_Arr, t1, t2)
			tmpweight.Table.Acc_Arr = append(tmpweight.Table.Acc_Arr, t1, t1+t2)

		}

		get_bonus_scatter_weight := func() {
			var total int
			for i := 1; i < len(rowng[4]); i++ {
				t, _ := strconv.Atoi(rowng[4][i])
				total += t
				tmpweight.RespinScatter.Init_Arr = append(tmpweight.RespinScatter.Init_Arr, t)
				tmpweight.RespinScatter.Acc_Arr = append(tmpweight.RespinScatter.Acc_Arr, total)

			}
		}

		get_table_weight()
		get_bonus_scatter_weight()

		switch rtproutie[i] {
		case "95":
			temp := &Game.WeightRTP_Module.FGWeight95
			*temp = tmpweight
		case "965":
			temp := &Game.WeightRTP_Module.FGWeight965
			*temp = tmpweight
		case "99":
			temp := &Game.WeightRTP_Module.FGWeight99
			*temp = tmpweight
		}
	}

}
