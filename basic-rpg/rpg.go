package main

type Character struct {
	Name       string
	Class      string
	Level      int
	Experience int
	Health     int
	Power      int
	Money      int
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

type Item struct {
	Name  string
	Price int
}

type Potion struct {
	Item
	Health int
}

type Weapon struct {
	Item
	Power int
}

func main() {

}
