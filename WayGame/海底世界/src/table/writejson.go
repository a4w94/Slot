package table

import (
	"encoding/json"
	"fmt"
	"os"
	"package/src/info"
	"strconv"
	"time"
)

type Json struct {
	GameName    string              `json:"GameName"`
	NG          StriTable           `json:"NGInfo"`
	FG          StriTable           `json:"FGInfo"`
	Paytablemap map[int]map[int]int `json:"PayTable"`
	InitialBet  int                 `json:"InitialBet"`
}

type StriTable struct {
	RTP95  map[int]Reel `json:"RTP95"`
	RTP965 map[int]Reel `json:"RTP965"`
	RTP99  map[int]Reel `json:"RTP99"`
}

type Reel struct {
	ReelString map[int]string `json:"Reel"`
}

func MapJson() {
	Init()

	var result Json
	result.GameName = info.GameName
	result.InitialBet = info.Bet
	result.NG.mapStriTable("NG")
	result.FG.mapStriTable("FG")
	result.PayTableMap()

	resultjson, _ := json.Marshal(result)
	WriteJson(resultjson)

}

func (result *StriTable) mapStriTable(status string) {

	foo := func(index int, table [info.Reelamount][]int) map[int]Reel {
		tmp := make(map[int]Reel)
		var tmpreel Reel
		tmpreel.MapReel(table)
		tmp[index] = tmpreel
		fmt.Println(tmp)
		return tmp
	}
	if status == "NG" {
		result.RTP95 = foo(0, Game.NGStriTablertp95)
		result.RTP965 = foo(0, Game.NGStriTablertp965)
		result.RTP99 = foo(0, Game.NGStriTablertp99)
	} else if status == "FG" {
		result.RTP95 = foo(0, Game.FGStriTablertp95)
		result.RTP965 = foo(0, Game.FGStriTablertp965)
		result.RTP99 = foo(0, Game.FGStriTablertp99)
	} else {
		fmt.Println("Wrong Input Status")
	}

}

func (result *Reel) MapReel(input [info.Reelamount][]int) {
	tmp := make(map[int]string)
	for i := 0; i < info.Reelamount; i++ {
		var arr = "{"
		for j, k := range input[i] {
			ele := strconv.Itoa(k)

			if j != len(input[i])-1 {
				arr = arr + ele + ","
			} else {
				arr = arr + ele + "}"

			}
		}
		tmp[i] = arr

	}
	result.ReelString = tmp

}

func (result *Json) PayTableMap() {
	tmp := make(map[int]map[int]int)
	fmt.Println(Game.PayTable)

	for i := 0; i < len(Game.PayTable); i++ {
		tmpmap := make(map[int]int)
		for k := 0; k < len(Game.PayTable[i]); k++ {
			index := len(Game.PayTable[i]) - k - 1

			tmpmap[index] = Game.PayTable[i][index]
		}
		tmp[i] = tmpmap
	}
	result.Paytablemap = tmp

}

func WriteJson(data []byte) {
	str_time := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("%s%s%s.json", info.GameName, str_time[5:7], str_time[8:10])

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("is not exist")
	} else {
		err := os.Remove(filename)
		if err != nil {
			fmt.Println("delete failed")
		} else {
			fmt.Println("delete", filename)
		}
	}

	fp, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Creat file", filename)
	}

	defer fp.Close()
	_, err1 := fp.Write(data)
	if err1 != nil {
		panic(err1)
	}
}
