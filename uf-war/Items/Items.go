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
	Sword
	Axe
	Dagger
	Bow
	Wand
	Staff
}

type Sword struct {
	Class string
	Power int
}

type Axe struct {
	Class string
	Power int
}

type Dagger struct {
	Class string
	Power int
}

type Bow struct {
	Class string
	Power int
}

type Wand struct {
	Class string
	Power int
}

type Staff struct {
	Class string
	Power int
}

type Armors struct {
	Plate
	IronPlate
	Leather
	Chitin
	Robe
	SilkRobe
}

type Plate struct {
	Class  string
	Level  int
	Health int
}
type IronPlate struct {
	Class  string
	Level  int
	Health int
}
type Leather struct {
	Class  string
	Level  int
	Health int
}
type Chitin struct {
	Class  string
	Level  int
	Health int
}
type Robe struct {
	Class  string
	Level  int
	Health int
}
type SilkRobe struct {
	Class  string
	Level  int
	Health int
}
