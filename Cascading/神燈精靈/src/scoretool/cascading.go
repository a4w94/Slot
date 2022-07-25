package scoretools

import (
	"fmt"
	"package/command"
	"package/graph"
	tools "package/src/rngtools"
	"package/src/table"

	"sync"
)

var g *graph.Graph

const (
	cascading_NeedCombo_Amount = 3
	unexist                    = 999
)

type Cascading_Combo struct {
	rng                    tools.Cascading_Rng
	Combo_and_Score_Result [command.Symbolamount][]Game_Result
	TotalScore             int
	CascadeTimes           int
	cascading_Combo_Result [command.Symbolamount]cascadingEachSymbol
}

type cascadingEachSymbol struct {
	symbol          int
	amount_in_Panel int
	result          []Game_Result
	visit           graph.Visit
	visit_Point_Arr [][]int
}

type Game_Result struct {
	Combo int
	Score int
}

func init() {
	g = graph.NewGraph()
	g.Init_Vertex()
	g.Init_Vertex_Graph()
}

func (c *Cascading_Combo) ComboJudgeCascading() {

	c.rng.ScanPanel()

judge:
	for {
		fmt.Println("初始盤面")
		c.rng.ScanPanel()
		c.init_Cascading_Combo_Result()
		c.panel_Scan_Symbol_Position()
		//c.DFS_mulitpleline()
		can_cascade := c.startDFS()
		if !can_cascade {
			break judge
		}
		c.CascadeTimes++
		c.rng.Generate_Rng_Cascading()
		c.Scan()
		c.rng.ScanPanel()
	}
}

//每次消除前初始化result struct
func (c *Cascading_Combo) init_Cascading_Combo_Result() {
	//初始化struct
	for i := 1; i < command.Symbolamount; i++ {
		c.cascading_Combo_Result[i].result = []Game_Result{}
		c.cascading_Combo_Result[i].amount_in_Panel = 0
		c.cascading_Combo_Result[i].visit.Visited = make(map[int]bool)
		c.cascading_Combo_Result[i].visit.Start = unexist
		c.cascading_Combo_Result[i].symbol = i

		for j := 0; j < command.Panel_len; j++ {
			c.cascading_Combo_Result[i].visit.Visited[j] = true

		}
	}

}

func turnPanel_a_Line(p command.GamePanel) command.CascadePanel {
	//將盤面改成一行
	var Panel_arr command.CascadePanel

	for _, r := range p {
		for _, m := range r {

			Panel_arr = append(Panel_arr, m)
		}
	}
	return Panel_arr
}

//尋訪盤面symbol 相對應位置
func (c *Cascading_Combo) panel_Scan_Symbol_Position() {

	//若盤面有該獎圖 則map 狀態改為未遍歷
	for index, r := range c.rng.Panel_Line {
		c.cascading_Combo_Result[r].visit.Visited[index] = false
		c.cascading_Combo_Result[r].amount_in_Panel++
		// if index < c.Cascading_Combo_Result[r].Visit.Start {
		// 	c.Cascading_Combo_Result[r].Visit.Start = index
		// }
		if r == command.Wild {
			for i := 1; i < command.Symbolamount; i++ {
				c.cascading_Combo_Result[i].visit.Visited[index] = false
				c.cascading_Combo_Result[i].amount_in_Panel++
			}
		}

	}

}

func (c *Cascading_Combo) DFS_mulitpleline() {

	var wg sync.WaitGroup
	var mux sync.Mutex
	wg.Add(command.Symbolamount - 1)
	for i := 1; i < command.Symbolamount; i++ {
		go func(i int) {
			defer wg.Done()
			mux.Lock()
			//盤面該symbol 總數量有超過才執行
			if c.cascading_Combo_Result[i].amount_in_Panel >= cascading_NeedCombo_Amount {
				c.cascading_Combo_Result[i].Loop_DFS(i)
				c.count_Combo_and_Score(c.cascading_Combo_Result[i])

				c.cascade_PanelSymbol(c.cascading_Combo_Result[i].visit_Point_Arr)

				// fmt.Println(c.Cascading_Combo_Result[i])
				// ScanPanel(c.Rng.Panel_Line)

			}
			mux.Unlock()
		}(i)
	}

	wg.Wait()

}

//開始探訪連線
func (c *Cascading_Combo) startDFS() bool {
	tmpscore := c.TotalScore
	for i := 1; i < command.Symbolamount; i++ {

		if i != command.Wild && i != command.SpecialWild && i != command.Scatter {

			//盤面該symbol 總數量有超過才執行
			if c.cascading_Combo_Result[i].amount_in_Panel >= cascading_NeedCombo_Amount {
				c.cascading_Combo_Result[i].Loop_DFS(i)
				c.count_Combo_and_Score(c.cascading_Combo_Result[i])

				c.cascade_PanelSymbol(c.cascading_Combo_Result[i].visit_Point_Arr)

				// fmt.Println(c.Cascading_Combo_Result[i])
				// ScanPanel(c.Rng.Panel_Line)

			}

		}
	}

	return c.TotalScore > tmpscore
}

//探訪連線
func (c *cascadingEachSymbol) Loop_DFS(sym int) {

	for i := 0; i < command.Panel_len; i++ {
		if !c.visit.Visited[i] {
			c.visit.Start = i
			c.visit.DFS(g)
			c.visit_Point_Arr = append(c.visit_Point_Arr, c.visit.Visited_Point_Arr)
			c.count_Combo_and_Score()
			c.Loop_DFS(sym)

		}

	}

}

//計算combo and score
func (c *cascadingEachSymbol) count_Combo_and_Score() {
	var tmp Game_Result
	tmp.Combo = len(c.visit.Visited_Point_Arr)
	tmp.Score = table.PayTable[c.symbol][tmp.Combo]
	c.result = append(c.result, tmp)

}

//計算combo and score
func (c *Cascading_Combo) count_Combo_and_Score(input cascadingEachSymbol) {

	for _, r := range input.result {
		if r.Combo >= cascading_NeedCombo_Amount {
			c.Combo_and_Score_Result[input.symbol] = append(c.Combo_and_Score_Result[input.symbol], r)
			c.TotalScore += r.Score
		}
	}
}

//將獎圖轉成消除代號
func (c *Cascading_Combo) cascade_PanelSymbol(cas [][]int) {

	for _, p := range cas {
		//有相鄰獎圖判斷相鄰數是否大於等於可消除數量
		if len(p) >= cascading_NeedCombo_Amount {
			for _, s := range p {
				c.rng.Panel_Line[s] = command.CascadeSymbol
			}
		}
	}

	c.rng.Turn_LinePanel_to_GamePanel()

}

func (c Cascading_Combo) Scan() {
	fmt.Println("Scan Cascade Result")

	fmt.Println("總分", c.TotalScore)
	fmt.Println("消除次數", c.CascadeTimes)

	fmt.Println("總結果")
	for i := 1; i < len(c.Combo_and_Score_Result); i++ {
		if len(c.Combo_and_Score_Result[i]) != 0 {

			fmt.Println("symbol", i)
			fmt.Println(c.Combo_and_Score_Result[i])
		}
	}
	fmt.Println()
	fmt.Println("細部結果")
	for i := 1; i < len(c.cascading_Combo_Result); i++ {
		if len(c.cascading_Combo_Result[i].result) != 0 {
			fmt.Println("symbol", i)
			fmt.Println(c.cascading_Combo_Result[i].result)
		}
	}
}
