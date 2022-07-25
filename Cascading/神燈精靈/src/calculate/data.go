package calculate

import (
	"math"
	"package/command"
)

type Math_Data struct {
	ALl_Score []int
	Avg       float64
	Std       float64
	Var       float64
}

func (c *Math_Data) Statistic_All_Score(m EachRoundResult) {

	c.ALl_Score = append(c.ALl_Score, m.MainGame.TotalScore+m.FreeGame.TotalScore)

}

func (c *Math_Data) Calculate_Data() {
	c.calculate_Avg()
	c.calculate_Std()
	c.calculate_Var()

	//fmt.Println(c)

}

func (c *Math_Data) calculate_Avg() {
	var total int
	for i := 0; i < len(c.ALl_Score); i++ {
		total += c.ALl_Score[i]
	}

	c.Avg = float64(total) / float64(len(c.ALl_Score))
}

func (c *Math_Data) calculate_Std() {
	var t float64
	for i := 0; i < len(c.ALl_Score); i++ {
		t += math.Pow(float64(c.ALl_Score[i])-c.Avg, 2)
	}
	c.Std = t / float64(len(c.ALl_Score))

}
func (c *Math_Data) calculate_Var() {
	c.Var = math.Sqrt(c.Std) / float64(command.PlayerBet)
}
