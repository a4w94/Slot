package tools

import "package/command"

///統計盤面scatter個數///
func CountPanelScatterAmount(panel command.GamePanel) int {

	var scatteramount int

	for i := 0; i < len(panel); i++ {

		for k := 0; k < command.Panel_Col; k++ {
			if panel[i][k] == command.Scatter {
				scatteramount++

			}
		}

	}

	return scatteramount

}

///統計盤面wild個數///
func CountPanelWildAmount(panel [command.Panel_Row][command.Panel_Col]int) int {

	var amount int

	for i := 0; i < command.Panel_Row; i++ {

		for k := 0; k < command.Panel_Col; k++ {
			if panel[i][k] == command.Wild {
				amount++

			}
		}

	}

	return amount

}
