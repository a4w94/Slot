package scoretools

import "package/command"

type Score_and_Combo interface {
	ComboJudge(panel command.GamePanel)
}

func (c *Cascading_Combo) ComboJudge() {
	c.ComboJudgeCascading()

}

func (c *Way_Game_Combo) ComboJudge(panel command.GamePanel) {

}
func (c *Line_Game_Combo) ComboJudge() {

}
