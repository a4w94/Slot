package public

import (
	"fmt"
	"package/src/info"
	"package/src/table"
)

var (

	///NG轉輪表///
	Ngstritable [info.Reelamount][]int

	///FG轉輪表///
	Fgstritable [info.Reelamount][]int

	///BG轉輪表///

	NGBonusTable [info.Reelamount][]int
	FGBonusTable [info.Reelamount][]int
)

func ChangeGlobalRtpModule(rtp int) {
	//table.Init()
	tmp := &info.UseRTPModule
	*tmp = rtp

	switch info.UseRTPModule {
	case 95:
		ng, fg := &Ngstritable, &Fgstritable
		*ng = table.Game.NGStriTablertp95
		*fg = table.Game.FGStriTablertp95
		ngbg, fgbg := &NGBonusTable, &FGBonusTable
		*ngbg = table.Game.NGBonusTable95
		*fgbg = table.Game.FGBonusTable95

	case 965:
		ng, fg := &Ngstritable, &Fgstritable
		*ng = table.Game.NGStriTablertp965
		*fg = table.Game.FGStriTablertp965
		ngbg, fgbg := &NGBonusTable, &FGBonusTable
		*ngbg = table.Game.NGBonusTable965
		*fgbg = table.Game.FGBonusTable965
	case 99:
		ng, fg := &Ngstritable, &Fgstritable
		*ng = table.Game.NGStriTablertp99
		*fg = table.Game.FGStriTablertp99
		ngbg, fgbg := &NGBonusTable, &FGBonusTable
		*ngbg = table.Game.NGBonusTable99
		*fgbg = table.Game.FGBonusTable99
	default:
		fmt.Println("wrong module")
	}
}

func ChangeTable(game_status string, rtp int) {
	var tmp RandWeight
	switch game_status {
	case info.GameStatus.MainGame:
		switch rtp {
		case 95:
			tmp.Rand(table.Game.NGWeight95.Table)
			if tmp.RandIndex == 0 {
				Ngstritable = table.Game.NGStriTablertp95
			} else {
				Ngstritable = table.Game.NGStriTable2
			}
		case 965:
			tmp.Rand(table.Game.NGWeight965.Table)
			if tmp.RandIndex == 0 {
				Ngstritable = table.Game.NGStriTablertp965
			} else {
				Ngstritable = table.Game.NGStriTable2
			}
		case 99:
			tmp.Rand(table.Game.NGWeight99.Table)
			if tmp.RandIndex == 0 {
				Ngstritable = table.Game.NGStriTablertp99
			} else {
				Ngstritable = table.Game.NGStriTable2
			}

		}

	case info.GameStatus.FreeGame:
		switch rtp {
		case 95:
			tmp.Rand(table.Game.FGWeight95.Table)
			if tmp.RandIndex == 0 {
				Fgstritable = table.Game.FGStriTablertp95
			} else {
				Fgstritable = table.Game.FGStriTable2
			}
		case 965:
			tmp.Rand(table.Game.FGWeight965.Table)
			if tmp.RandIndex == 0 {
				Fgstritable = table.Game.FGStriTablertp965
			} else {
				Fgstritable = table.Game.FGStriTable2
			}
		case 99:
			tmp.Rand(table.Game.FGWeight99.Table)
			if tmp.RandIndex == 0 {
				Fgstritable = table.Game.FGStriTablertp99
			} else {
				Fgstritable = table.Game.FGStriTable2
			}

		}

	}

}
