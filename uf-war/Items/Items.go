package Items

type Items struct {
	Potion
	Weapons
	Armors
}

type Potion struct {
	ManaPotion
	HealthPotion
}

type ManaPotion struct {
	Mana int
}

type HealthPotion struct {
	Health int
}

type Weapons struct {
	WarriorWeapons
	AssassinWeapons
	MageWeapons
}
type WarriorWeapons struct {
	Class string
	Sword
	Axe
}
type Sword struct {
	Level int
	Power int
}

type Axe struct {
	Level int
	Power int
}

type AssassinWeapons struct {
	Class string
	Dagger
	Bow
}

type Dagger struct {
	Level int
	Power int
}

type Bow struct {
	Level int
	Power int
}

type MageWeapons struct {
	Class string
	Wand
	Staff
}

type Wand struct {
	Level int
	Power int
}

type Staff struct {
	Level int
	Power int
}

type Armors struct {
	WarriorClothes
	AssassinClothes
	MageClothes
}

type WarriorClothes struct {
	Class string
	Plate
	IronPlate
}
type Plate struct {
	Level  int
	Health int
}
type IronPlate struct {
	Level  int
	Health int
}

type AssassinClothes struct {
	Class string
	Leather
	Chitin
}
type Leather struct {
	Level  int
	Health int
}
type Chitin struct {
	Level  int
	Health int
}

type MageClothes struct {
	Class string
	Robe
	SilkRobe
}
type Robe struct {
	Level  int
	Health int
}
type SilkRobe struct {
	Level  int
	Health int
}
