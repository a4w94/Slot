package weight

import (
	"math/rand"
)

type RandomResult struct {
	RandSeed int

	Result interface{}
}

type Weight struct {
	Index      []interface{}
	InitWeight []int
	AccWeight  []int
}

func (result *RandomResult) RandResult(input Weight) {
	result.RandSeed = rand.Intn(input.AccWeight[len(input.AccWeight)-1])
	if result.RandSeed < input.AccWeight[0] {

		if len(input.Index) != 0 {
			result.Result = input.Index[0]
		}
	} else {
		for i := 0; i < len(input.AccWeight)-1; i++ {

			if input.AccWeight[i] <= result.RandSeed && result.RandSeed < input.AccWeight[i+1] {

				if len(input.Index) != 0 {
					result.Result = input.Index[i+1]
				}
				break
			}
		}

	}
}
