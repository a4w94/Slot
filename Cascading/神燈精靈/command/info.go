package command

import (
	"fmt"
)

const (

	///遊戲名稱///
	GameName string = "神燈精靈"

	///遊戲模式///
	GameMode GameStatusName = Cascading

	Retrigger = false

	///基本投注額Bet///
	Bet int = 50

	PlayBetLevel = 1
	//玩家下注額
	PlayerBet int = Bet * PlayBetLevel

	///理論值ＲＴＰ///
	ThmRTP float32 = 0.965

	///ＲＴＰ95///
	RTP95 float32 = 0.95

	///ＲＴＰ99///
	RTP99 float32 = 0.99
)

///倍數表///

func Infomation() {
	fmt.Println("遊戲名稱 : ", GameName)
	fmt.Println("PayLines : ", GameMode)
	fmt.Println("轉輪數 : ", Panel_Col)

	fmt.Println("基本投注額 : ", Bet)
	fmt.Println("獎圖總數 : ", Symbolamount)
	fmt.Println("combo數目0~5combo : ", Comboresultnum)

	fmt.Println("理論值ＲＴＰ ：", ThmRTP)

}
