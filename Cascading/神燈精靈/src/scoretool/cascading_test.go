package scoretools

import (
	"fmt"
	"math/rand"
	"package/command"
	"package/src/public"
	"package/src/table"
	"testing"
	"time"
)

func BenchmarkTest(b *testing.B) {
	rand.Seed(int64(time.Now().UnixNano()))
	fmt.Println("test generate rng")

	//for i := 0; i < b.N; i++ {
	table.InitTable()
	public.Init(965)

	var tmp = command.GamePanel{
		{2, 1, 5, 5, 3},
		{3, 3, 1, 9, 3},
		{1, 1, 1, 5, 3},
	}
	var t Cascading_Combo
	t.rng.Generate_Rng(command.MainGame)
	t.rng.Panel = tmp
	t.rng.Panel_Line = turnPanel_a_Line(t.rng.Panel)

	t.ComboJudgeCascading()
	var tmp2 Score_and_Combo
	tmp2 = t
	tmp2.ComboJudge(t.rng.Panel)
	//}
}
