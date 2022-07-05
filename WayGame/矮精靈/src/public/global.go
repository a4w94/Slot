package public

import (
	"fmt"
	"package/src/info"
)

var (

	///NG轉輪表///
	Ngstritable [info.Reelamount][]int

	///FG轉輪表///
	Fgstritable [info.Reelamount][]int

	///BG轉輪表///

)

func ChangeGlobalRtpModule(rtp int) {

	tmp := &info.UseRTPModule
	*tmp = rtp
	//ng, fg := &Ngstritable, &Fgstritable

	switch info.UseRTPModule {

	// case 95:
	// 	*ng = table.Game.NGStriTablertp95
	// 	*fg = table.Game.FGStriTablertp95

	// case 965:
	// 	*ng = table.Game.NGStriTablertp965
	// 	*fg = table.Game.FGStriTablertp965
	// case 99:
	// 	*ng = table.Game.NGStriTablertp99
	// 	*fg = table.Game.FGStriTablertp99
	// case 92:
	// 	*ng = table.Game.NGStriTablertp92
	// 	*fg = table.Game.FGStriTablertp92
	// case 90:
	// 	*ng = table.Game.NGStriTablertp90
	// 	*fg = table.Game.FGStriTablertp90
	default:
		fmt.Println("wrong module")
	}
}
