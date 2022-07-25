package public

import (
	"fmt"
	"package/command"
	"package/src/table"
)

var (

	///NG轉輪表///
	Ngstritable [command.Panel_Col][]int

	///FG轉輪表///
	Fgstritable [command.Panel_Col][]int

	///BG轉輪表///

	//MegaWay
	//NG Table
	Ngstritable_MegaWay table.MegaWay_Table
	//FG Table
	Fgstritable_MegaWay table.MegaWay_Table

	//Cascading
	//NG Table
	Ngstritable_Cascading table.Cascading_Weight
	//FG Table
	Fgstritable_Cascading table.Cascading_Weight
)

func Init(rtp int) {
	switch command.GameMode {
	case command.Cascading:
		ChangeGlobalRtpModule_Cascading(rtp)
	case command.MegaWay:
		ChangeGlobalRtpModule_MegaWay(rtp)
	case command.WayGame:
		ChangeGlobalRtpModule_WayGame(rtp)
	case command.LineGame:
		ChangeGlobalRtpModule_LineGame(rtp)
	}
}

func ChangeGlobalRtpModule_LineGame(rtp int) {

	tmp := &command.UseRTPModule
	*tmp = rtp
	ng, fg := &Ngstritable, &Fgstritable

	switch command.UseRTPModule {

	case 95:
		*ng = table.LineGameStriTable.NGStriTablertp95
		*fg = table.LineGameStriTable.FGStriTablertp95

	case 965:
		*ng = table.LineGameStriTable.NGStriTablertp965
		*fg = table.LineGameStriTable.FGStriTablertp965
	case 99:
		*ng = table.LineGameStriTable.NGStriTablertp99
		*fg = table.LineGameStriTable.FGStriTablertp99
	case 92:
		*ng = table.LineGameStriTable.NGStriTablertp92
		*fg = table.LineGameStriTable.FGStriTablertp92
	case 90:
		*ng = table.LineGameStriTable.NGStriTablertp90
		*fg = table.LineGameStriTable.FGStriTablertp90
	default:
		fmt.Println("wrong module")
	}
}

func ChangeGlobalRtpModule_WayGame(rtp int) {

	tmp := &command.UseRTPModule
	*tmp = rtp
	ng, fg := &Ngstritable, &Fgstritable

	switch command.UseRTPModule {

	case 95:
		*ng = table.WayGameStriTable.NGStriTablertp95
		*fg = table.WayGameStriTable.FGStriTablertp95

	case 965:
		*ng = table.WayGameStriTable.NGStriTablertp965
		*fg = table.WayGameStriTable.FGStriTablertp965
	case 99:
		*ng = table.WayGameStriTable.NGStriTablertp99
		*fg = table.WayGameStriTable.FGStriTablertp99
	case 92:
		*ng = table.WayGameStriTable.NGStriTablertp92
		*fg = table.WayGameStriTable.FGStriTablertp92
	case 90:
		*ng = table.WayGameStriTable.NGStriTablertp90
		*fg = table.WayGameStriTable.FGStriTablertp90
	default:
		fmt.Println("wrong module")
	}
}

func ChangeGlobalRtpModule_MegaWay(rtp int) {
	tmp := &command.UseRTPModule
	*tmp = rtp
	ng, fg := &Ngstritable_MegaWay, &Fgstritable_MegaWay

	switch command.UseRTPModule {

	case 95:
		*ng = table.MegaWayTable.NGStriTablertp95
		*fg = table.MegaWayTable.FGStriTablertp95

	case 965:
		*ng = table.MegaWayTable.NGStriTablertp965
		*fg = table.MegaWayTable.FGStriTablertp965
	case 99:
		*ng = table.MegaWayTable.NGStriTablertp99
		*fg = table.MegaWayTable.FGStriTablertp99
	case 92:
		*ng = table.MegaWayTable.NGStriTablertp92
		*fg = table.MegaWayTable.FGStriTablertp92
	case 90:
		*ng = table.MegaWayTable.NGStriTablertp90
		*fg = table.MegaWayTable.FGStriTablertp90
	default:
		fmt.Println("wrong module")
	}

}

func ChangeGlobalRtpModule_Cascading(rtp int) {
	tmp := &command.UseRTPModule
	*tmp = rtp
	ng, fg := &Ngstritable_Cascading, &Fgstritable_Cascading

	switch command.UseRTPModule {

	case 95:
		*ng = table.CascadingTable.NGStriTablertp95
		*fg = table.CascadingTable.FGStriTablertp95

	case 965:
		*ng = table.CascadingTable.NGStriTablertp965
		*fg = table.CascadingTable.FGStriTablertp965
	case 99:
		*ng = table.CascadingTable.NGStriTablertp99
		*fg = table.CascadingTable.FGStriTablertp99
	case 92:
		*ng = table.CascadingTable.NGStriTablertp92
		*fg = table.CascadingTable.FGStriTablertp92
	case 90:
		*ng = table.CascadingTable.NGStriTablertp90
		*fg = table.CascadingTable.FGStriTablertp90
	default:
		fmt.Println("wrong module")
	}

}
