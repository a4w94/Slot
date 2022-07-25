package calculate

import (
	"fmt"
	"os"
	scoretools "package/src/scoretool"
	"package/src/table"
	"reflect"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func (result *TotalRoundResultRate) PrintResult() {
	toString := func(input float64) string {
		//var result string
		result := strconv.FormatFloat(input*100, 'f', 3, 64) + "%"
		return result
	}

	RTPPrint := func() {

		data := [][]string{
			[]string{"Total RTP", toString(result.TotalRTP)},
			[]string{"Main Game RTP", toString(result.MainGameRTP_with_scatter)},
			[]string{"Free Game RTP", toString(result.FreeGameRTP_with_scatter)},
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Rate%"})
		table.SetRowLine(true)
		table.AppendBulk(data)

		table.Render() // Send output
	}

	MainGame := func() {

		func() {
			data := [][]string{
				[]string{"Main Game Without Scatter RTP", toString(result.MainGameRTP_with_scatter - result.MainGame_ScatterRTP)},
				[]string{"Free Game Hit Rate", toString(result.MainGame_TriggeFree_Rate)},
				[]string{"Scatter RTP", toString(result.MainGame_ScatterRTP)},
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Name", "Rate%"})

			table.SetRowLine(true)
			table.AppendBulk(data)

			table.Render() // Send output
		}()

		//倍率區間
		func() {
			data := [][]string{}

			for t := 0; t < scoretools.MutipelRange; t++ {
				if t == scoretools.MutipelRange-1 {
					tmp := []string{strconv.Itoa(scoretools.Multiple[t]) + "以上", toString(result.MainGame_ScoreRange_Rate[t]), toString(result.MainGame_ScoreRange_RTP[t])}
					data = append(data, tmp)
				} else {
					tmp := []string{strconv.Itoa(scoretools.Multiple[t]) + "~" + strconv.Itoa(scoretools.Multiple[t+1]), toString(result.MainGame_ScoreRange_Rate[t]), toString(result.MainGame_ScoreRange_RTP[t])}
					data = append(data, tmp)

				}

			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"倍率區間", "Rate %", "RTP %"})

			table.SetRowLine(true)
			table.AppendBulk(data)

			table.Render() // Send output

		}()

		//各個symbol hit rate
		func() {
			data := [][]string{}

			for i := 1; i < len(table.PayTableSymbol); i++ {
				tmp := []string{table.PayTableSymbol[i]}
				for j := 1; j < len(result.SymbolComboTotalHit.NGHitRate[i]); j++ {
					tmp = append(tmp, toString(result.SymbolComboTotalHit.NGHitRate[i][j]))
				}
				data = append(data, tmp)
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Symbol Combo Hit Rate", "1", "2", "3", "4", "5"})

			table.SetRowLine(true)
			table.AppendBulk(data)

			table.Render() // Send output

		}()

		func() {
			var totalrtp float64

			data := [][]string{}

			for i := 1; i < len(table.PayTableSymbol); i++ {
				tmp := []string{table.PayTableSymbol[i]}
				var symboltotalrtp float64
				for j := 1; j < len(result.SymbolComboTotalHit.NGRTP[i]); j++ {
					tmp = append(tmp, toString(result.SymbolComboTotalHit.NGRTP[i][j]))
					symboltotalrtp += result.SymbolComboTotalHit.NGRTP[i][j]
					totalrtp += result.SymbolComboTotalHit.NGRTP[i][j]
				}
				tmp = append(tmp, toString(symboltotalrtp))
				data = append(data, tmp)
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Symbol Combo RTP ", "1", "2", "3", "4", "5", "Total"})
			table.SetFooter([]string{"", "", "", "", "", "Total", toString(totalrtp)}) // Add Footer

			table.SetRowLine(true)
			table.AppendBulk(data)

			table.Render() // Send output
		}()

	}

	FreeGame := func() {

		func() {
			data := [][]string{
				[]string{"Retrigger Hit Rate", toString(result.FreeGame_Retrigger_Rate)},
				[]string{"Scatter RTP", toString(result.FreeGame_ScatterRTP)},
				[]string{"中獎率", toString(result.FreeGame_Score_Hit_Rate)},
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Name", "Rate%"})

			table.SetRowLine(true)
			table.AppendBulk(data)

			table.Render() // Send output
		}()

		//倍率區間
		func() {
			data := [][]string{}

			for t := 0; t < scoretools.MutipelRange; t++ {
				if t == scoretools.MutipelRange-1 {
					tmp := []string{strconv.Itoa(scoretools.Multiple[t]) + "以上", toString(result.FreeGame_ScoreRange_Rate[t]), toString(result.FreeGame_ScoreRange_RTP[t])}
					data = append(data, tmp)
				} else {
					tmp := []string{strconv.Itoa(scoretools.Multiple[t]) + "~" + strconv.Itoa(scoretools.Multiple[t+1]), toString(result.FreeGame_ScoreRange_Rate[t]), toString(result.FreeGame_ScoreRange_RTP[t])}
					data = append(data, tmp)
				}

			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"倍率區間", "Rate %", "RTP %"})

			table.SetRowLine(true)
			table.AppendBulk(data)

			table.Render() // Send output

		}()

		func() {

		}()

	}
	fmt.Println("===========> Total <===================================")
	RTPPrint()
	fmt.Println()

	fmt.Println("===========> Main Game <===============================")
	MainGame()
	fmt.Println()

	fmt.Println("===========> Free Game <===============================")
	FreeGame()
	fmt.Println()

	fmt.Println("===========> Main Game 補充<===============================")

	fmt.Println()

	fmt.Println("===========> Free Game 補充<===============================")

	fmt.Println()
	fmt.Println("===========> 數據<===============================")
	// fmt.Println("得分平均:", result.Math_Data.Avg)
	// fmt.Println("得分變異數:", result.Math_Data.Std)
	fmt.Println("得分標準差:", result.Math_Data.Var)

	fmt.Println()

}

func (result MainGameEachRoundResult) PrintEachRoudResult() {
	t := reflect.TypeOf(result)
	v := reflect.ValueOf(result)
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name, ":", v.Field(i))
	}
	fmt.Println()
}

func (result FreeGameEachRoundResult) PrintEachRoudResult() {
	t := reflect.TypeOf(result)
	v := reflect.ValueOf(result)
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name, ":", v.Field(i))
	}
	fmt.Println()
}

func (result FreeGameTotalResult) PrintEachRoudResult() {
	t := reflect.TypeOf(result)
	v := reflect.ValueOf(result)
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name, ":", v.Field(i))
	}
	fmt.Println()
}
