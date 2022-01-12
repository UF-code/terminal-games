package main

import (
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

	Creator.CreateJson("Items.json")
}
