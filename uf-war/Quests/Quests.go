package Quests

type Quests struct {
	Level0
	Level1
	Level2
	Level3
	Level4
	Level5
	Level6
	Level7
	Level8
	Level9
}

type Reward struct {
	Exp  int
	Gold int
}
type Level0 struct {
	Level     int
	Completed bool
	Reward
}
type Level1 struct {
	Level     int
	Completed bool
	Reward
}
type Level2 struct {
	Level     int
	Completed bool
	Reward
}
type Level3 struct {
	Level     int
	Completed bool
	Reward
}
type Level4 struct {
	Level     int
	Completed bool
	Reward
}
type Level5 struct {
	Level     int
	Completed bool
	Reward
}
type Level6 struct {
	Level     int
	Completed bool
	Reward
}
type Level7 struct {
	Level     int
	Completed bool
	Reward
}
type Level8 struct {
	Level     int
	Completed bool
	Reward
}
type Level9 struct {
	Level     int
	Completed bool
	Reward
}
