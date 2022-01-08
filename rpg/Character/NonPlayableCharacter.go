package Character

import (
	"fmt"
	"rpg/Inventory"
	"rpg/Inventory/tired"
)

type Shop struct {
	Character_Template
	Inventory.Bag
}

func NPC() {
	fmt.Println("Hey NPC")

	s := Shop{
		Character_Template{
			Health: 100,
			Power:  200,
		},
		Inventory.Bag{
			Items: []Inventory.Item{{Name: "Sword", Power: 100}, {Name: "Dagger", Power: 200}},
		},
	}

	fmt.Println(s)
	tired.Hey()
}
