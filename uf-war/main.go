package main

import (
	"fmt"
	"uf-war/Character"
	"uf-war/Character/PlayableCharacter"
	"uf-war/Creator"
	"uf-war/Inventory"
	"uf-war/Inventory/Tired"
	"uf-war/Items"
	"uf-war/Shop"
)

func main() {
	Character.Character()
	PlayableCharacter.PC()
	Inventory.Test()
	Tired.Hey()

	Shop.Market_test()

	if !Creator.CheckFileExists("./Items/Items.json") {
		Creator.CreateJson("./Items/Items.json")
		fmt.Println("File is created")
	} else {
		fmt.Println("Sucker hehehe")
	}

	// Creator.WriteToJson("./Items/Items.json", []byte("hey there"))
	// Creator.ReadJson("./Items/Items.json")

	////////////////////////////////
	//
	// i1 := Creator.Item{
	// 	CreateItem.Item{
	// 		Name: "Hey Dude",
	// 	},
	// }

	// Creator.WriteToJson("./Items/Items.json", i1)
	// i2 := Creator.Quest{
	// 	CreateQuest.Quest{
	// 		Name: "Hey Dude",
	// 	},
	// }

	// Creator.WriteToJson("./Items/Items.json", i2)
	Creator.Create()
	//
	/////////////////////

	axe := Items.Axe{
		Level: 10,
		Power: 2000,
	}
	fmt.Println(axe)
}
