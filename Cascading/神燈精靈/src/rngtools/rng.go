package tools

///盤面工具///
import (
	"package/command"
)

type Rng interface {
	Generate_Rng(status command.GameStatusName)
}

///生成盤面

//DoubleWay盤面//
func DoubleWayRng(panel [command.Panel_Row][command.Panel_Col]int) [command.Panel_Row][command.Panel_Col]int {

	var reversepanel [command.Panel_Row][command.Panel_Col]int
	for k := 0; k < command.Panel_Row; k++ {
		for i := 0; i < command.Panel_Col; i++ {
			reversepanel[k][command.Panel_Col-1-i] = panel[k][i]
		}
	}

	return reversepanel
}
