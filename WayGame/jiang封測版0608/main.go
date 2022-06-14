package main

import (
	"package/src/calculate"
)

func main() {

	//calculate.Simulate(10000000, 965)
	calculate.PlayerSim(100000, 965)

	//製圖
	//chart.WriteDataChart()
	// for a := 1; a < 6; a++ {
	// 	for b := 0; b < 6; b++ {
	// 		for c := 0; c < 6; c++ {
	// 			fmt.Println(a, ",", b, ",", c)

	// 		}
	// 	}
	// }

}
