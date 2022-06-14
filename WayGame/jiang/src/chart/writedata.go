package chart

import (
	"fmt"
	"os"
	"package/src/calculate"

	"github.com/go-echarts/go-echarts/charts"
)

type Data struct {
	RTP95  calculate.TotalRoundResultRate
	RTP965 calculate.TotalRoundResultRate
	RTP99  calculate.TotalRoundResultRate
}
type SimData struct {
	RTP95  calculate.PlayerSimResultRate
	RTP965 calculate.PlayerSimResultRate
	RTP99  calculate.PlayerSimResultRate
}

func WriteDataChart() {
	p := charts.NewPage()

	p = RTPModuleData(p)
	p = PlaySimData(p)

	filename := fmt.Sprint("數據", ".html")
	f, err := os.Create("./chart/" + filename)
	if err != nil {
		fmt.Println("failed build")
	} else {
		fmt.Println("create file:", filename)
	}
	p.Render(f)
}

func RTPModuleData(p *charts.Page) *charts.Page {
	testsession := 1000000

	var result Data
	result.RTP965 = calculate.Simulate(testsession, 965)
	result.RTP95 = calculate.Simulate(testsession, 965)
	result.RTP99 = calculate.Simulate(testsession, 965)

	p.Add(
		ScoreMultipleRangeRate_Chart("MainGame", result),
		ScoreMultipleRangeRTP_Chart("MainGame", result),

		ScoreMultipleRangeRate_Chart("FreeGame", result),
		ScoreMultipleRangeRTP_Chart("FreeGame", result),
	)

	return p
}

func PlaySimData(p *charts.Page) *charts.Page {
	var playerAmount int
	playerAmount = 1000
	var result SimData
	result.RTP95 = calculate.PlayerSim(playerAmount, 965)
	result.RTP965 = calculate.PlayerSim(playerAmount, 965)
	result.RTP99 = calculate.PlayerSim(playerAmount, 965)
	p.Add(
		PlayerScoreWinLoseRate_Chart(result),
	)

	return p

}
