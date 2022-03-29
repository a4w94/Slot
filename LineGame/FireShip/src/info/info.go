package info

import "fmt"

const (

	///遊戲名稱///
	GameName string = "火燒連環船"

	///遊戲模式///
	GameMode string = "LineGame"

	Retrigger = false

	///Reel 轉輪數///
	Reelamount int = 5

	Col int = 4

	///基本投注額Bet///
	Bet int = 40

	PlayBetLevel = 1
	//玩家下注額
	PlayerBet int = Bet * PlayBetLevel

	///線數///
	Linenum int = 40

	///獎圖總數+1 沒有零號///
	Symbolamount int = 14

	///combo數目0~5combo///
	Comboresultnum int = 6

	///WILD 代號///
	Wild int = 11
	///Scatter 代號///
	Scatter int = 12

	Space int = 13
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
