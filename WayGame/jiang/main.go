package main

import (
	"fmt"
	"math/rand"
	"package/src/table"
	"time"
)

type Name struct {
	user string
	id   int
}

func main() {

	rand.Seed(time.Now().Unix())
	table.Init()
	tmp := table.RandomResult{}
	tmp.RandResult(table.Game.RTP965.MainGame_Panel_Grow)
	fmt.Println(table.Game.RTP965.MainGame_Panel_Grow)
	fmt.Println(tmp)

	//calculate.Simulate(1000000, 965)
	//製圖
	//chart.WriteDataChart()

}
