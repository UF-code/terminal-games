package Items

type Items struct {
	Potion
	Weapons
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
