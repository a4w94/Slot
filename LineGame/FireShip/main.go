package main

import (
	"package/src/calculate"
)

type Name struct {
	user string
	id   int
}

func main() {

	calculate.Simulate(1, 965)

	// var tmp calculate.MainGameEachRoundResult
	// tmp.Panel = [info.Col][info.Reelamount]int{
	// 	{1, 2, 3, 1, 5},
	// 	{1, 2, 3, 12, 5},
	// 	{1, 2, 11, 4, 5},
	// 	{1, 2, 3, 4, 5},
	// }
	// tmp.BonusGame()
	//製圖
	//chart.WriteDataChart()
	//Test()

	// calculate.GenerateBonus(info.GameStatus.MainGame, 8)
	// var tmp public.RandWeight
	// tmp.Rand(table.Game.NGWeight965.RespinScatter)
	// fmt.Println(tmp)

}
