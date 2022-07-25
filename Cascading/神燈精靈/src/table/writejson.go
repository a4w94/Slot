package table

import (
	"encoding/json"
	"fmt"
	"os"
	"package/command"
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
	InitTable()

	var result Json
	result.GameName = command.GameName
	result.InitialBet = command.Bet
	// result.NG.mapStriTable("NG")
	// result.FG.mapStriTable("FG")
	result.PayTableMap()

	resultjson, _ := json.Marshal(result)
	WriteJson(resultjson)

}

// func (result *StriTable) mapStriTable(status string) {

// 	foo := func(index int, table [command.Panel_Col][]int) map[int]Reel {
// 		tmp := make(map[int]Reel)
// 		var tmpreel Reel
// 		tmpreel.MapReel(table)
// 		tmp[index] = tmpreel
// 		fmt.Println(tmp)
// 		return tmp
// 	}
// 	if status == "NG" {
// 		// result.RTP95 = foo(0, NGStriTablertp95)
// 		// result.RTP965 = foo(0, NGStriTablertp965)
// 		// result.RTP99 = foo(0, NGStriTablertp99)
// 	} else if status == "FG" {
// 		// result.RTP95 = foo(0, FGStriTablertp95)
// 		// result.RTP965 = foo(0, FGStriTablertp965)
// 		// result.RTP99 = foo(0, FGStriTablertp99)
// 	} else {
// 		fmt.Println("Wrong Input Status")
// 	}

// }

func (result *Reel) MapReel(input [command.Panel_Col][]int) {
	tmp := make(map[int]string)
	for i := 0; i < command.Panel_Col; i++ {
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
	fmt.Println(PayTable)

	for i := 0; i < len(PayTable); i++ {
		tmpmap := make(map[int]int)
		for k := 0; k < len(PayTable[i]); k++ {
			index := len(PayTable[i]) - k - 1

			tmpmap[index] = PayTable[i][index]
		}
		tmp[i] = tmpmap
	}
	result.Paytablemap = tmp

}

func WriteJson(data []byte) {
	str_time := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("%s%s%s.json", command.GameName, str_time[5:7], str_time[8:10])

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
