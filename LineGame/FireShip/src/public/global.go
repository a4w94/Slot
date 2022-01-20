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
	table.Init()
	tmp := &info.UseRTPModule
	*tmp = rtp

	switch info.UseRTPModule {
	case 95:
		ng, fg := &Ngstritable, &Fgstritable
		*ng = table.Game.NGStriTablertp95
		*fg = table.Game.FGStriTablertp95

	case 965:
		ng, fg := &Ngstritable, &Fgstritable
		*ng = table.Game.NGStriTablertp965
		*fg = table.Game.FGStriTablertp965
	case 99:
		ng, fg := &Ngstritable, &Fgstritable
		*ng = table.Game.NGStriTablertp99
		*fg = table.Game.FGStriTablertp99
	default:
		fmt.Println("wrong module")
	}
}
