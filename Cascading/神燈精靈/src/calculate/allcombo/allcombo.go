package allcombo

import (
	"fmt"
	"package/command"
	"package/src/public"
)

type Panel struct {
	P []command.GamePanel
}

func ProductAllPanel() []command.GamePanel {

	alllen := 1
	fmt.Println(public.Ngstritable)
	for _, m := range public.Ngstritable {
		alllen *= len(m)
	}
	//bar := progressbar.Default(int64(alllen))
	var result []command.GamePanel
	table := public.Ngstritable

	for a := 0; a < len(table[0]); a++ {
		for b := 0; b < len(table[1]); b++ {
			for c := 0; c < len(table[2]); c++ {
				// for d := 0; d < len(table[3]); d++ {
				// 	for e := 0; e < len(table[4]); e++ {
				// 		fmt.Println(a, b, c, d, e)

				// 		var r [][]int
				// 		for m := 0; m < command.Col; m++ {

				// 			r = append(r, []int{table[0][(a+m)%len(table[0])], table[1][(b+m)%len(table[1])], table[2][(c+m)%len(table[2])], table[3][(d+m)%len(table[3])], table[4][(e+m)%len(table[4])]})

				// 		}
				// 		var tmp command.GamePanel

				// 		for i, row := range r {
				// 			for j, col := range row {
				// 				tmp[i][j] = col
				// 			}
				// 		}

				// 		result = append(result, tmp)

				// 		bar.Add(1)

				// 	}

				// }
			}
		}

	}

	return result
}

func Test(input int) {
	table := [5][]int{
		[]int{11, 12, 13},
		[]int{21, 22, 23, 24},
		[]int{31, 32, 33, 34, 35},
		[]int{41, 42, 43, 44, 45, 46},
		[]int{51, 52, 53},
	}
	for i := 0; i < len(table[i]); i++ {
		fmt.Println(input, i)
		if input < 4 {
			input++
			Test(input)
		}
	}

}
