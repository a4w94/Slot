package tools

import (
	"fmt"
	"math/rand"
	"package/command"
	"package/src/public"
	"package/src/table"
	"testing"
	"time"
)

func BenchmarkTestGenerate_Rng_test(b *testing.B) {
	rand.Seed(int64(time.Now().UnixNano()))
	fmt.Println("test generate rng")

	var tmp Cascading_Rng
	table.InitTable()
	public.Init(965)
	tmp.Generate_Rng(command.MainGame)

}
