package tools

import (
	info "package/src/info"
)

///生成盤面
func GameRng_Special(gameStatus string) info.GamePanel {
	var result info.GamePanel
	//	var randNumArr []int

	if gameStatus == info.GameStatus.MainGame {

	} else if gameStatus == info.GameStatus.FreeGame {

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
