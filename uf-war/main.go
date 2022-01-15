package main

import (
	"fmt"
	"uf-war/Character"
	"uf-war/Character/PlayableCharacter"
	"uf-war/Creator"
	"uf-war/Inventory"
	"uf-war/Inventory/Tired"
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

	Creator.Create()
}
