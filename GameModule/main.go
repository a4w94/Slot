package main

import "package/src/calculate"

type Name struct {
	user string
	id   int
}

func main() {
	calculate.Simulate(10, 965)
	//製圖
	//chart.WriteDataChart()

}
