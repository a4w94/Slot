package chart

import (
	"fmt"
	scoretools "package/src/scoretool"
	"strconv"

	"github.com/go-echarts/go-echarts/charts"
)

func ScoreMultipleRangeRate_Chart(game_status string, input Data) *charts.Bar {

	nameItems := []string{}

	for i := 0; i < len(scoretools.Multiple); i++ {
		if i == len(scoretools.Multiple)-1 {
			tmp := strconv.Itoa(scoretools.Multiple[i]) + "以上"
			nameItems = append(nameItems, tmp)

		} else {
			tmp := fmt.Sprint(strconv.Itoa(scoretools.Multiple[i]), "~", strconv.Itoa(scoretools.Multiple[i+1]))
			nameItems = append(nameItems, tmp)
		}
	}
	bar := charts.NewBar()

	bar.SetGlobalOptions(charts.TitleOpts{Title: game_status + "倍率區間"},
		charts.YAxisOpts{SplitLine: charts.SplitLineOpts{Show: true}},
		charts.XAxisOpts{Name: "倍率區間"},
	)

	bar.SetSeriesOptions(
		charts.LabelTextOpts{Show: true},
	)

	bar.AddXAxis(nameItems)
	if game_status == "MainGame" {
		bar.AddYAxis("95", input.RTP95.MainGame_ScoreRange_Rate).
			AddYAxis("965", input.RTP965.MainGame_ScoreRange_Rate).
			AddYAxis("99", input.RTP99.MainGame_ScoreRange_Rate)
	} else if game_status == "FreeGame" {
		bar.AddYAxis("95", input.RTP95.FreeGame_ScoreRange_Rate).
			AddYAxis("965", input.RTP965.FreeGame_ScoreRange_Rate).
			AddYAxis("99", input.RTP99.FreeGame_ScoreRange_Rate)
	} else {
		fmt.Println("wrong status")
	}

	// filename := fmt.Sprint(game_status, "倍率區間", ".html")
	// f, err := os.Create("./chart/" + filename)
	// if err != nil {
	// 	fmt.Println("failed build")
	// } else {
	// 	fmt.Println("create file:", filename)
	// }
	// bar.Render(f)
	return bar
}

func ScoreMultipleRangeRTP_Chart(game_status string, input Data) *charts.Bar {

	nameItems := []string{}

	for i := 0; i < len(scoretools.Multiple); i++ {
		if i == len(scoretools.Multiple)-1 {
			tmp := strconv.Itoa(scoretools.Multiple[i]) + "以上"
			nameItems = append(nameItems, tmp)

		} else {
			tmp := fmt.Sprint(strconv.Itoa(scoretools.Multiple[i]), "~", strconv.Itoa(scoretools.Multiple[i+1]))
			nameItems = append(nameItems, tmp)
		}
	}
	bar := charts.NewBar()

	bar.SetGlobalOptions(charts.TitleOpts{Title: game_status + "倍率區間RTP"},
		charts.YAxisOpts{SplitLine: charts.SplitLineOpts{Show: true}},
		charts.XAxisOpts{Name: "倍率區間"},
	)

	bar.SetSeriesOptions(
		charts.LabelTextOpts{Show: true},
	)

	bar.AddXAxis(nameItems)
	if game_status == "MainGame" {
		bar.AddYAxis("95", input.RTP95.MainGame_ScoreRange_RTP).
			AddYAxis("965", input.RTP965.MainGame_ScoreRange_RTP).
			AddYAxis("99", input.RTP99.MainGame_ScoreRange_RTP)
	} else if game_status == "FreeGame" {
		bar.AddYAxis("95", input.RTP95.FreeGame_ScoreRange_RTP).
			AddYAxis("965", input.RTP965.FreeGame_ScoreRange_RTP).
			AddYAxis("99", input.RTP99.FreeGame_ScoreRange_RTP)
	} else {
		fmt.Println("wrong status")
	}

	// filename := fmt.Sprint(game_status, "倍率區間RTP", ".html")
	// f, err := os.Create("./chart/" + filename)
	// if err != nil {
	// 	fmt.Println("failed build")
	// } else {
	// 	fmt.Println("create file:", filename)
	// }
	// bar.Render(f)
	return bar
}

func PlayerScoreWinLoseRate_Chart(input SimData) *charts.Bar {
	nameItems := []string{"RTP95", "RTP965", "RTP99"}
	scorelose := []float64{input.RTP95.Number_of_People_ScoreZero_Rate, input.RTP965.Number_of_People_ScoreZero_Rate, input.RTP99.Number_of_People_ScoreZero_Rate}
	scorewin := []float64{input.RTP95.Number_of_People_ScoreDouble_Rate, input.RTP965.Number_of_People_ScoreDouble_Rate, input.RTP99.Number_of_People_ScoreDouble_Rate}

	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "輸錢贏錢比例"})
	bar.AddXAxis(nameItems).
		AddYAxis("輸到沒錢", scorelose, charts.BarOpts{Stack: "stack"}).
		AddYAxis("贏兩倍", scorewin, charts.BarOpts{Stack: "stack"})
	bar.XYReversal()
	bar.SetSeriesOptions(charts.LabelTextOpts{Show: true})

	return bar
}
