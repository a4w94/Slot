package tools

import (
	"package/command"
)

///生成盤面
func GameRng_MegaWay(gameStatus command.GameStatusName) command.GamePanel {
	var result command.GamePanel
	//	var randNumArr []int

	if gameStatus == command.MainGame {

	} else if gameStatus == command.FreeGame {

	}

	// fmt.Println(randNumArr)
	// fmt.Println("rngresult")
	// for _, m := range rngresult {
	// 	fmt.Println(m)
	// }

	// fmt.Println("grow")
	// for _, m := range result.AfterGrowPanel {
	// 	fmt.Println(m)
	// }
	return result

}
