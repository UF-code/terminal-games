package Character

import (
	"fmt"
	"rpg/Inventory"
)

type Shop struct {
	Character_Template
	Inventory.Items
}

func NPC() {
	fmt.Println("Hey NPC")

	s := Shop{
		Character_Template{
			Health: 100,
			Power:  200,
		},
		Inventory.Items{
			Item: []Inventory.Item{
				Item
			}
		},
	}
}
