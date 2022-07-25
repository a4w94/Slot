package command

const (
	///Reel 轉輪數///
	Panel_Col int = 5

	Panel_Row int = 3

	///線數///
	Linenum int = 1

	MegaWay_EachReel_Max int = 5

	Panel_len = Panel_Row * Panel_Col
)

type GamePanel [Panel_Row][Panel_Col]int

//單行盤面
type CascadePanel []int
