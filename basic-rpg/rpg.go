package main

type Character struct {
	Name   string
	Class  string
	Health int
	Power  int
}

type NPC_Shop struct {
	Character
}

type NPC_Quest struct {
	Character
}

type Enemy struct {
	Character
}

type Player struct {
	Character
}

func main() {

}
