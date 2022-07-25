package tools

import (
	"math/rand"
	"package/command"
)

//中列拓展wild//
func MidWildExpand(panel [command.Panel_Row][command.Panel_Col]int) [command.Panel_Row][command.Panel_Col]int {

	for k := 0; k < command.Panel_Col; k++ {
		if panel[1][k] == command.Wild {
			panel[0][k] = command.Wild
			panel[2][k] = command.Wild
		}
	}

	return panel

}

//隨機wild//
func RandomWild(panel [command.Panel_Row][command.Panel_Col]int) [command.Panel_Row][command.Panel_Col]int {
	//fmt.Println("input", panel)

	randomwildweight := [command.Panel_Col][4]int{
		{0, 9, 11, 20},
		{0, 9, 11, 20},
		{0, 9, 11, 20},
		//	{0, 9, 11, 20},
	}

	var wildposition [command.Panel_Col]int

	for i := 0; i < command.Panel_Col; i++ {
		seed := rand.Intn(randomwildweight[i][3])
		//fmt.Println("seed", seed)

		for k := 1; k < 4; k++ {
			if seed >= randomwildweight[i][k-1] && seed < randomwildweight[i][k] {

				wildposition[i] = k - 1
			}
		}
	}
	//fmt.Println(wildposition)

	for i := 0; i < command.Panel_Col; i++ {
		if wildposition[i] == 1 {
			panel[0][i] = command.Wild
			panel[1][i] = command.Wild
			panel[2][i] = command.Wild
		} else {
			panel[wildposition[i]][i] = command.Wild
		}

	}
	//fmt.Println("output", panel)

	return panel

}
