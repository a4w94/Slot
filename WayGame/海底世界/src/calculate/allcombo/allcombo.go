package allcombo

import (
	"fmt"
	"package/src/info"
	"package/src/public"

	"github.com/schollz/progressbar/v3"
)

type Panel struct {
	P [][info.Reelamount]int
}

func ProductAllPanel() []Panel {
	//grow := 0
	alllen := 1
	fmt.Println(public.Ngstritable)
	for _, m := range public.Ngstritable {
		alllen *= len(m)
	}
	bar := progressbar.Default(int64(alllen))
	var result []Panel
	table := public.Ngstritable
	//all:
	for a := 0; a < len(table[0]); a++ {
		for b := 0; b < len(table[1]); b++ {
			for c := 0; c < len(table[2]); c++ {
				for d := 0; d < len(table[3]); d++ {
					for e := 0; e < len(table[4]); e++ {
						var r Panel
						for m := 0; m < 3; m++ {
							r.P = append(r.P, [info.Reelamount]int{table[0][(a+m)%len(table[0])], table[1][(b+m)%len(table[1])], table[2][(c+m)%len(table[2])], table[3][(d+m)%len(table[3])], table[4][(e+m)%len(table[4])]})
						}
						result = append(result, r)

						bar.Add(1)

					}

				}
			}
		}

	}

	return result
}
