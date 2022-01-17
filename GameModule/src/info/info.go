package info

import "fmt"

const (

	///遊戲名稱///
	GameName string = "舞動峇里島"

	///遊戲模式///
	GameMode string = "LineGame"

	Retrigger = false

	///Reel 轉輪數///
	Reelamount int = 5

	Col int = 3

	///基本投注額Bet///
	Bet int = 2

	PlayBetLevel = 1
	//玩家下注額
	PlayerBet int = Bet * PlayBetLevel

	///線數///
	Linenum int = 1

	///獎圖總數+1 沒有零號///
	Symbolamount int = 13

	///combo數目0~5combo///
	Comboresultnum int = 6

	///WILD 代號///
	Wild int = 12
	///Scatter 代號///
	Scatter int = 11

	// Bonus int = 11

	///理論值ＲＴＰ///
	ThmRTP float32 = 0.965

	///ＲＴＰ95///
	RTP95 float32 = 0.95

	///ＲＴＰ99///
	RTP99 float32 = 0.99
)

type GameStatusName struct {
	MainGame string
	FreeGame string
	WayGame  string
	LineGame string
}

var (
	GameStatus = GameStatusName{
		MainGame: "MainGame",
		FreeGame: "FreeGame",
		WayGame:  "WayGame",
		LineGame: "LineGame",
	}
	UseRTPModule int
)

///倍數表///

func Infomation() {
	fmt.Println("遊戲名稱 : ", GameName)
	fmt.Println("PayLines : ", GameMode)
	fmt.Println("轉輪數 : ", Reelamount)

	fmt.Println("基本投注額 : ", Bet)
	fmt.Println("獎圖總數 : ", Symbolamount)
	fmt.Println("combo數目0~5combo : ", Comboresultnum)

	fmt.Println("理論值ＲＴＰ ：", ThmRTP)

}
