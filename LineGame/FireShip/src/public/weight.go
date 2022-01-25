package public

import (
	"math/rand"
	"package/src/table"
)

type RandWeight struct {
	Randnumber int
	RandIndex  int
	table.WeightArr
}

func (r *RandWeight) Rand(input table.WeightArr) {
	r.WeightArr = input
	r.Randnumber = rand.Intn(input.Acc_Arr[len(input.Acc_Arr)-1])
r:
	for i := 0; i < len(input.Acc_Arr); i++ {
		if i == 0 {
			if r.Randnumber < input.Acc_Arr[i] {
				r.RandIndex = i
				break r
			}
		} else {
			if r.Randnumber >= input.Acc_Arr[i-1] && r.Randnumber < input.Acc_Arr[i] {
				r.RandIndex = i
				break r
			}

		}
	}

}
