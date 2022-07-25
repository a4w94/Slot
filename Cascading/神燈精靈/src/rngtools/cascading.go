package tools

import (
	"fmt"
	"math/rand"
	"package/command"
	"package/graph"
	"package/src/public"
	"package/weight"
)

type Cascading_Rng struct {
	command.GameStatusName                          //NG or FG
	Use_Panel_Index_Randresult weight.RandomResult  //使用盤面 rand 結果
	Use_Symbol_Weight          weight.Weight        //使用盤面對應symbol 權重
	Panel                      command.GamePanel    //產生盤面
	Panel_Line                 command.CascadePanel //盤面改為一個row
	ScattterAmount             int
}

//產生盤面
func (c *Cascading_Rng) Generate_Rng(status command.GameStatusName) {
	c.GameStatusName = status
	c.rand_Panel_Index()
	c.generate_Rng_Symbol()
	c.turnPanel_a_Line()
	c.insert_Scatter_Init()

	c.Scan()

}

//rand 使用盤面
func (c *Cascading_Rng) rand_Panel_Index() {
	switch c.GameStatusName {
	case command.MainGame:
		c.Use_Panel_Index_Randresult.RandResult(public.Ngstritable_Cascading.Panel_Weight)
		c.Use_Symbol_Weight = public.Ngstritable_Cascading.Symbol_Weight_for_EachPanel[c.Use_Panel_Index_Randresult.Result.(int)]
	case command.FreeGame:
		c.Use_Panel_Index_Randresult.RandResult(public.Fgstritable_Cascading.Panel_Weight)
		c.Use_Symbol_Weight = public.Fgstritable_Cascading.Symbol_Weight_for_EachPanel[c.Use_Panel_Index_Randresult.Result.(int)]

	}

}

//盤面產生symbol
func (c *Cascading_Rng) generate_Rng_Symbol() {

	for i := 0; i < command.Panel_Row; i++ {
		for j := 0; j < command.Panel_Col; j++ {
			sym := c.rand_Symbol().Result
			c.Panel[i][j] = sym.(int)
		}

	}

}

//rand symbol
func (c *Cascading_Rng) rand_Symbol() weight.RandomResult {
	var tmp weight.RandomResult
	tmp.RandResult(c.Use_Symbol_Weight)
	return tmp

}

//rand scatter 初始盤面數量
func rand_Scatter_Init() int {
	var tmp weight.RandomResult
	tmp.RandResult(public.Ngstritable_Cascading.Scatter.InitWeight)

	return tmp.Result.(int)

}

func (c *Cascading_Rng) insert_Scatter_Init() {
	c.ScattterAmount = rand_Scatter_Init()

	for i := 0; i < c.ScattterAmount; i++ {
		seed := rand.Intn(command.Panel_len)
		c.Panel_Line[seed] = command.Scatter
	}
	c.turnPanel_a_Line()

}

func (c *Cascading_Rng) turnPanel_a_Line() {
	//將盤面改成一行
	var Panel_arr command.CascadePanel

	for _, r := range c.Panel {
		for _, m := range r {

			Panel_arr = append(Panel_arr, m)
		}
	}
	c.Panel_Line = Panel_arr
}

func (c *Cascading_Rng) Generate_Rng_Cascading() {
	c.symbol_Fall()
	c.insert_Scatter_Fall()
	c.rand_Fall_Symbol()

}

//獎圖降落
func (c *Cascading_Rng) symbol_Fall() {
	fmt.Println("symbol_Fall")
	for _, p := range graph.Cascade_Position {
		for i, k := range p {
			//要置換sym
			sta_sym := c.Panel_Line[k]
			if sta_sym == command.CascadeSymbol {
				for index := i + 1; index < command.Panel_Row; index++ {
					compare_sym := c.Panel_Line[p[index]]
					if compare_sym != command.CascadeSymbol {
						c.Panel_Line[k] = compare_sym
						c.Panel_Line[p[index]] = sta_sym
						break
					}
				}
			}

		}
	}
	c.Turn_LinePanel_to_GamePanel()

}

//rand scatter 落下盤面數量
func rand_Scatter_Fall() int {
	var tmp weight.RandomResult
	tmp.RandResult(public.Ngstritable_Cascading.Scatter.FallWeight)

	return tmp.Result.(int)

}

func (c *Cascading_Rng) insert_Scatter_Fall() {
	c.ScattterAmount = rand_Scatter_Fall()
	var cas_sym_posarr []int
	for p, s := range c.Panel_Line {
		if s == command.CascadeSymbol {
			cas_sym_posarr = append(cas_sym_posarr, p)
		}
	}

	for i := 0; i < c.ScattterAmount; i++ {
		seed := rand.Intn(len(cas_sym_posarr))
		c.Panel_Line[seed] = command.Scatter
	}
	c.Turn_LinePanel_to_GamePanel()

}

//落下補牌
func (c *Cascading_Rng) rand_Fall_Symbol() {

	for i := 0; i < len(c.Panel_Line); i++ {
		if c.Panel_Line[i] == command.CascadeSymbol {
			c.Panel_Line[i] = c.rand_Symbol().Result.(int)
		}
	}

	c.Turn_LinePanel_to_GamePanel()

}

func (c *Cascading_Rng) Turn_LinePanel_to_GamePanel() {
	for i, s := range c.Panel_Line {
		row := i / command.Panel_Col
		col := i % command.Panel_Col
		c.Panel[row][col] = s
	}
}

func (c *Cascading_Rng) Scan() {
	fmt.Println("scan cascading rng")
	fmt.Println(c.GameStatusName)
	fmt.Println(c.Use_Panel_Index_Randresult.RandSeed, c.Use_Panel_Index_Randresult.Result)
	fmt.Println(c.Use_Symbol_Weight)

}

func (c *Cascading_Rng) ScanPanel() {
	fmt.Println("遊戲盤面")
	for _, r := range c.Panel {
		fmt.Println(r)
	}

	fmt.Println("直線盤面")
	fmt.Println(c.Panel_Line)
	fmt.Println()
}
