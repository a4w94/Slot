package command

type GameStatusName string

const (
	MainGame  GameStatusName = "MainGame"
	FreeGame  GameStatusName = "FreeGame"
	WayGame   GameStatusName = "FreeGame"
	LineGame  GameStatusName = "LineGame"
	Cascading GameStatusName = "Cascading"
	MegaWay   GameStatusName = "MegaWay"
)

var (
	UseRTPModule int
)
